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
}

func NewArticle(Id int) *Article {
	return &Article{
		Id: Id,
	}
}
