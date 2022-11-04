package model

import "time"

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUser(name, password, email string) *User {
	return &User{
		Name:     name,
		Password: password,
		Email:    email,
	}
}
