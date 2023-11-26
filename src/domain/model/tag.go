package model

import "time"

type Tag struct {
	Id          int       `gorm:"primaryKey;" json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Slug        string    `json:"slug,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}

type ArticleTags struct {
	ArticleId int `gorm:"primaryKey;" json:"article_id"`
	TagId     int `gorm:"primaryKey;" json:"tag_id"`
}
