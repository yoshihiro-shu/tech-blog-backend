package table

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

const (
	StatusUnpublished = iota
	StatusPublished
	StatusClosed
)
