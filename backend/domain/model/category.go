package model

import "time"

type Category struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ParentId    int       `json:"parentId"`
	CreatedAt   time.Time `json:"createdAt"`
}
