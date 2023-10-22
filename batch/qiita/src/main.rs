use std::error::Error;
use std::env;
use tokio; // tokioは非同期ランタイムです
use tokio_postgres::{NoTls};

mod qiita_response;
mod tag_map;

#[derive(Debug)]
struct Tag {
    id: i32,
    name: String,
}

fn construct_db_url() -> String {
    // 環境変数から値を取得します。環境変数が設定されていない場合はデフォルト値を使用します。
    let db_host = env::var("DB_HOST").unwrap_or_else(|_| "localhost".to_string());
    let db_port = env::var("DB_PORT").unwrap_or_else(|_| "5432".to_string());
    let db_user = env::var("DB_USER").unwrap_or_else(|_| "user".to_string());
    let db_password = env::var("DB_PASSWORD").unwrap_or_else(|_| "password".to_string());
    let db_name = env::var("DB_NAME").unwrap_or_else(|_| "database".to_string());
    let db_ssl = env::var("DB_SSL").unwrap_or_else(|_| "disable".to_string());

    // SSLの設定に基づいて、SSLモードを指定します。
    let ssl_mode = match db_ssl.as_str() {
        "disable" => "?sslmode=disable",
        "require" => "?sslmode=require",
        "prefer" => "?sslmode=prefer",
        _ => "", // デフォルトのSSLモードを使用します（例：SSLモードが指定されていない場合や不明な値の場合）
    };

    // すべての情報を結合して、接続URLを形成します。
    format!(
        "postgresql://{}:{}@{}:{}/{}{}",
        db_user, db_password, db_host, db_port, db_name, ssl_mode
    )
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

    let db_endpoint = construct_db_url();

    // 非同期接続
    let (db_client, db_connection) = tokio_postgres::connect(&db_endpoint, NoTls).await?;

    // The connection object performs the actual communication with the database,
    // so spawn it off to run on its own.
    tokio::spawn(async move {
        if let Err(e) = db_connection.await {
            eprintln!("connection error: {}", e)
        }
    });

    // Get Tags From DB
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

    // Insert New Tags
    let tag_map = tag_map::create_map();
    for r in &res {
        for t in &r.tags {
            // insert tag if not exists
            let check = db_client.query("SELECT * FROM tags WHERE name = $1", &[&t.name]).await?;
            if check.len() == 0 {
                let slug = tag_map.get(&t.name.as_str());
                let mut tag_id: i32 = 0;
                if slug.is_none() {
                    let inserted_tag = db_client.query("INSERT INTO tags (name, slug) VALUES ($1, $2) RETURNING id", &[&t.name, &t.name]).await?;
                    tag_id = inserted_tag[0].get("id");
                } else {
                    let inserted_tag = db_client.query("INSERT INTO tags (name, slug) VALUES ($1, $2) RETURNING id", &[&t.name, &slug]).await?;
                    tag_id = inserted_tag[0].get("id");
                }
                tags.push(Tag{
                    id: tag_id,
                    name: t.name.clone(),
                });
                println!("inserted tag: {}", t.name);
            }
        }
    }

    // Insert articles from Qiita
    for r in res {
        let check = db_client.query("SELECT * FROM articles WHERE title = $1", &[&r.title]).await?;
        if check.len() != 0 {
            println!("already exists!");
            continue;
        }

        let inserted_data = db_client.query("INSERT INTO articles (user_id, thumbnail_url, title, content, status) VALUES ($1, $2, $3, $4, $5) RETURNING id", &[&1, &"",&r.title, &r.body, &2]).await?;
        let inserted_id: i32 = inserted_data[0].get("id");
        for t in r.tags {
            // insert tag if not exists
            let check = db_client.query("SELECT * FROM tags WHERE name = $1", &[&t.name]).await?;
            if check.len() == 0 {
               let inserted_tag = db_client.query("INSERT INTO tags (name, slug) VALUES ($1, $2) RETURNING id", &[&t.name, &t.name]).await?;
               let tag_id: i32 = inserted_tag[0].get("id");
               tags.push(Tag{
                   id: tag_id,
                   name: t.name.clone(),
               })
            }
            // insert article_tags
            for tt in &tags {
                if t.name == tt.name {
                    db_client.execute("INSERT INTO article_tags (article_id, tag_id) VALUES ($1, $2)", &[&inserted_id, &tt.id]).await?;
                }
            }
        }
        println!("inserted article: {}", r.title);
    }

    Ok(())
}
