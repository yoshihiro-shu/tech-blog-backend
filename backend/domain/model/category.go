package model

import "time"

type Category struct {
	Id          int       `gorm:"primaryKey;" json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	ParentId    int       `json:"parentId"`
	CreatedAt   time.Time `json:"createdAt"`
}
