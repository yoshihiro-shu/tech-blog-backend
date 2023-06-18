package model

import (
	"time"
)

type Article struct {
	Id           int       `gorm:"primaryKey;" json:"id"`
	UserId       int       `json:"userId"`
	ThumbnailUrl string    `json:"thumbnailUrl"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	CategoryId   int       `json:"categoryId"`
	User         *User     `gorm:"foreignKey:UserId;" json:"user"`
	Category     *Category `gorm:"foreignKey:CategoryId;" json:"category"`
	Tags         []Tag     `gorm:"many2many:article_tags;" json:"tags"`
}

func NewArticle(Id int) *Article {
	return &Article{
		Id: Id,
	}
}
