package controller

import (
	"database/sql"

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

func (u *User) Login(db *sql.DB) error {
	// get password by user.email
	err := u.GetByEmail(db)
	if err != nil {
		return err
	}

	// compare password and crypt password
	err = u.VerifyPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
