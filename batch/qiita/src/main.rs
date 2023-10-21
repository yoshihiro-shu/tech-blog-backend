use chrono::{NaiveDateTime};
use std::error::Error;
use tokio; // tokioは非同期ランタイムです
use tokio_postgres::{NoTls};

mod qiita_response;
mod entity;

#[derive(Debug)]
struct Tag {
    id: i32,
    name: String,
}


#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    // 例として、httpbin.orgのGETエンドポイントを使います
    let request_url = "https://qiita.com/api/v2/items?page=1&per_page=100&query=user:yoshihiro-shu";

    // クライアントのインスタンスを作成します
    let api_client = reqwest::Client::new();

    // リクエストを送り、レスポンスを待ちます
    let response = api_client.get(request_url).send().await?;

    let mut response_text: String = Default::default();
    // 成功した場合、レスポンスのテキストを表示します
    if response.status().is_success() {
        response_text = response.text().await?;
    } else {
        println!("Failed to get a successful response. Status: {}", response.status());
    }

    let res: qiita_response::QiitaResponse = serde_json::from_str(&response_text).unwrap();

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

    let mut tags: Vec<Tag> = Vec::new();
    let rows = db_client.query("SELECT id, name FROM tags", &[]).await?;
    for row in rows {
        let id: i32 = row.get("id");
        let name: String = row.get("name");
        tags.push(Tag{
            id: id,
            name: name.clone(),
        });
    };

    for r in res {
        let check = db_client.query("SELECT * FROM articles WHERE title = $1", &[&r.title]).await?;
        if check.len() != 0 {
            println!("already exists!");
            continue;
        }

        let inserted_data = db_client.query("INSERT INTO articles (user_id, thumbnail_url, title, content, status) VALUES ($1, $2, $3, $4, $5) RETURNING id", &[&1, &"",&r.title, &r.body, &2]).await?;
        let inserted_id: i32 = inserted_data[0].get("id");
        for t in r.tags {
            for tt in &tags {
                if t.name == tt.name {
                    println!("tag = {:?}", tt.name);
                    db_client.execute("INSERT INTO article_tags (article_id, tag_id) VALUES ($1, $2)", &[&a_id, &tt.id]).await?;
                }
            }
        }
    }

    Ok(())
}
