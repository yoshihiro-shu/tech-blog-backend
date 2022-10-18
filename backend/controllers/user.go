package controllers

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/model"
)

type User struct {
	model.User
	Password string `json:"-"`
}

func NewUser(email, password string) *User {
	user := &User{Password: password}
	user.Email = email
	return user
}

func (u *User) Login(db *pg.DB) error {
	// get password by user.email
	err := u.GetByEmail(db)
	if err != nil {
		return fmt.Errorf("Email is mistaken")
	}

	// compare password and crypt password
	err = u.VerifyPassword(u.Password)
	if err != nil {
		return fmt.Errorf("password is mistaken")
	}
	return nil
}
