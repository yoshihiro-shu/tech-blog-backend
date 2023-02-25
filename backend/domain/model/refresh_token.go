package model

import (
	"time"
)

type RefreshToken struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	JwtId     string    `json:"jwt_id"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User     `pg:"fk:user_id" json:"user"`
}
