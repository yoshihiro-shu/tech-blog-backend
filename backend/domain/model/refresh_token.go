package model

import (
	"time"
)

type RefreshToken struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	JwtId     string    `json:"jwt_id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      *User     `pg:"fk:user_id" json:"user"`
}
