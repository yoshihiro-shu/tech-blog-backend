package model

import (
	"time"
)

type Article struct {
	Id           int       `gorm:"primaryKey;" json:"id,omitempty"`
	UserId       int       `json:"userId,omitempty"`
	ThumbnailUrl string    `json:"thumbnailUrl,omitempty"`
	Title        string    `json:"title,omitempty"`
	Content      string    `json:"content,omitempty"`
	Status       int       `json:"status,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
	CategoryId   int       `json:"categoryId,omitempty"`
	User         *User     `gorm:"foreignKey:UserId;" json:"user,omitempty"`
	Category     *Category `gorm:"foreignKey:CategoryId;" json:"category,omitempty"`
	Tags         []Tag     `gorm:"many2many:article_tags;" json:"tags,omitempty"`
}

func NewArticle(Id int) *Article {
	return &Article{
		Id: Id,
	}
}
