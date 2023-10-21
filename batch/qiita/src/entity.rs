use chrono::NaiveDateTime;

#[derive(Debug)]
pub struct Tag {
    pub id: i32,
    pub name: String,
    pub slug: String,
    pub description: Option<String>,
    pub created_at: Option<NaiveDateTime>,
}

pub struct Category {
    pub id: i64,
    pub name: String,
    pub slug: String,
    pub description: Option<String>,
    pub created_at: Option<NaiveDateTime>,
}

pub struct Article {
    pub id : i64,
    pub user_id: i64,
    pub thumbnail_url: String,
    pub title: String,
    pub content: String,
    pub status: i64,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
    pub caategory_id: i64,
    pub category: Category,
    pub tag_ids: Vec<i64>,
    pub tag: Vec<Tag>,
}
