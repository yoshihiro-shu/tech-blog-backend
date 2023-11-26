package model

import "time"

type User struct {
	Id        int       `gorm:"primaryKey;" json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func NewUser(name, password, email string) *User {
	return &User{
		Name:     name,
		Password: password,
		Email:    email,
	}
}
