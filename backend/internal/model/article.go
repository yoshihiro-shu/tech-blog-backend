package model

import (
	"time"

	"github.com/go-pg/pg"
)

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

func (a *Article) Insert(db *pg.DB) error {
	err := db.Insert(a)
	if err != nil {
		return err
	}
	return nil
}

func (a Article) GetAll(db *pg.DB) ([]Article, error) {
	var articles []Article
	err := db.Model(&articles).Select()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *Article) GetByID(db *pg.DB) error {

	// Select user by primary key.
	err := db.Select(a)
	if err != nil {
		return err
	}

	return nil
}
