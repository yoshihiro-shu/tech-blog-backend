use serde_json::Value;
use serde::{Serialize, Deserialize};


pub type QiitaResponse = Vec<Response>;

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Response {
    #[serde(rename = "rendered_body")]
    pub rendered_body: String, // article content?
    pub body: String,
    pub coediting: bool,
    #[serde(rename = "comments_count")]
    pub comments_count: i64,
    #[serde(rename = "created_at")]
    pub created_at: String,
    pub group: Value,
    pub id: String,
    #[serde(rename = "likes_count")]
    pub likes_count: i64,
    pub private: bool,
    #[serde(rename = "reactions_count")]
    pub reactions_count: i64,
    #[serde(rename = "stocks_count")]
    pub stocks_count: i64,
    pub tags: Vec<Tag>,
    pub title: String,
    #[serde(rename = "updated_at")]
    pub updated_at: String,
    pub url: String,
    pub user: User,
    #[serde(rename = "page_views_count")]
    pub page_views_count: Value,
    #[serde(rename = "team_membership")]
    pub team_membership: Value,
    #[serde(rename = "organization_url_name")]
    pub organization_url_name: Value,
    pub slide: bool,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Tag {
    pub name: String,
    pub versions: Vec<Value>,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct User {
    pub description: String,
    #[serde(rename = "facebook_id")]
    pub facebook_id: String,
    #[serde(rename = "followees_count")]
    pub followees_count: i64,
    #[serde(rename = "followers_count")]
    pub followers_count: i64,
    #[serde(rename = "github_login_name")]
    pub github_login_name: String,
    pub id: String,
    #[serde(rename = "items_count")]
    pub items_count: i64,
    #[serde(rename = "linkedin_id")]
    pub linkedin_id: String,
    pub location: String,
    pub name: String,
    pub organization: String,
    #[serde(rename = "permanent_id")]
    pub permanent_id: i64,
    #[serde(rename = "profile_image_url")]
    pub profile_image_url: String,
    #[serde(rename = "team_only")]
    pub team_only: bool,
    #[serde(rename = "twitter_screen_name")]
    pub twitter_screen_name: String,
    #[serde(rename = "website_url")]
    pub website_url: String,
}