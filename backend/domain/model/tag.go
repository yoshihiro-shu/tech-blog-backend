package model

import "time"

type Tag struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ArticleTags struct {
	ArticleId int `json:"article_id"`
	TagId     int `json:"tag_id"`
}
