use reqwest::Error;
use tokio; // tokioは非同期ランタイムです

mod qiita_response;

#[tokio::main]
async fn main() -> Result<(), Error> {
    // 例として、httpbin.orgのGETエンドポイントを使います
    let request_url = "https://qiita.com/api/v2/items?page=1&per_page=100&query=user:yoshihiro-shu";

    // クライアントのインスタンスを作成します
    let client = reqwest::Client::new();

    // リクエストを送り、レスポンスを待ちます
    let response = client.get(request_url).send().await?;

    // 成功した場合、レスポンスのテキストを表示します
    if response.status().is_success() {
        let response_text = response.text().await?;
        let res: qiita_response::QiitaResponse = serde_json::from_str(&response_text).unwrap();
        for r in res {
            println!("title = {:?}", r.title as String);
            for t in r.tags {
                println!("tag = {:?}", t.name as String);
            }
            // println!("content = {:?}", r.rendered_body as String);
        }
    } else {
        println!("Failed to get a successful response. Status: {}", response.status());
    }

    Ok(())
}
