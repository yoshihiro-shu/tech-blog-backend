package model

import "time"

type Article struct {
	Id         int       `json:"id"`
	UserId     int       `json:"userId"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	CategoryId int       `json:"categoryId"`
	User       *User     `pg:"fk:user_id"`
	Category   *Category `pg:"fk:category_id"`
	Tags       []Tag     `pg:"many2many:article_tags"`
}

func NewArticle(Id int) *Article {
	return &Article{
		Id: Id,
	}
}
