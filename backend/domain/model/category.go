package model

import "time"

type Category struct {
	Id          int       `gorm:"primaryKey;" json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}
