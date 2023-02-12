package model

import (
	"time"

	"github.com/go-pg/pg"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"` // TODO fix to `json:"-"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u *User) SetBcryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) VerifyPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Insert(db *pg.DB) error {
	err := db.Insert(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetAll(db *pg.DB) ([]User, error) {
	var users []User
	err := db.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) GetByEmail(db *pg.DB) error {

	err := db.Model(u).Where("email = ?", u.Email).Select()
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetByID(db *pg.DB) error {

	// Select user by primary key.
	err := db.Select(u)
	if err != nil {
		return err
	}

	return nil
}
