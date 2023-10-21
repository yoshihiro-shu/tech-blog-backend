package model

import "time"

type Tag struct {
	Id          int       `gorm:"primaryKey;" json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ArticleTags struct {
	ArticleId int `gorm:"primaryKey;" json:"article_id"`
	TagId     int `gorm:"primaryKey;" json:"tag_id"`
}
