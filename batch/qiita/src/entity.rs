
pub struct Tag {
    Id: i64,
    Name: String,
    Slug: String,
    Description: String,
    CreatedAt: DateTime<Utc>,
}

pub struct Category {
    Id: i64,
    Name: String,
    Slug: String,
    Description: String,
    CreatedAt: DateTime<Utc>,
}

pub struct Article {
    Id : i64,
    UserId: i64,
    ThumbnailUrl: String,
    Title: String,
    Content: String,
    Status: i64,
    CreatedAt: DateTime<Utc>,
    UpdatedAt: DateTime<Utc>,
    CategoryId: i64,
    Category: Category,
    TagIds: Vec<i64>,
    Tag: Vec<Tag>,
}
