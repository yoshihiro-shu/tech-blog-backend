use std::error::Error;
use tokio; // tokioは非同期ランタイムです
use tokio_postgres::{NoTls};

mod qiita_response;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    // 例として、httpbin.orgのGETエンドポイントを使います
    let request_url = "https://qiita.com/api/v2/items?page=1&per_page=100&query=user:yoshihiro-shu";

    // クライアントのインスタンスを作成します
    let api_client = reqwest::Client::new();

    // リクエストを送り、レスポンスを待ちます
    let response = api_client.get(request_url).send().await?;

    // 成功した場合、レスポンスのテキストを表示します
    if response.status().is_success() {
        let response_text = response.text().await?;
        let res: qiita_response::QiitaResponse = serde_json::from_str(&response_text).unwrap();
        for r in res {
            println!("title = {:?}", r.title);
            for t in r.tags {
                println!("tag = {:?}", t.name);
            }
            // println!("content = {:?}", r.rendered_body as String);
        }
    } else {
        println!("Failed to get a successful response. Status: {}", response.status());
    }

    let db_endpoint = "postgresql://postgres:password@127.0.0.1:5432/postgres";

    // 非同期接続
    let (db_client, db_connection) = tokio_postgres::connect(db_endpoint, NoTls).await?;

    // The connection object performs the actual communication with the database,
    // so spawn it off to run on its own.
    tokio::spawn(async move {
        if let Err(e) = db_connection.await {
            eprintln!("connection error: {}", e)
        }
    });

    // Now we can execute a simple statement that just returns its parameter.
    let rows = db_client
        .query("SELECT $1::TEXT", &[&"hello world"])
        .await?;

    // And then check that we got back the same string we sent over.
    let value: &str = rows[0].get(0);
    assert_eq!(value, "hello world");
    println!("value = {:?}", value);

    Ok(())
}
