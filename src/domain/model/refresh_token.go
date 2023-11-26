package model

import (
	"time"
)

type RefreshToken struct {
	Id        int       `json:"id,omitempty"`
	UserId    int       `json:"user_id,omitempty"`
	JwtId     string    `json:"jwt_id,omitempty"`
	ExpiredAt time.Time `json:"expired_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	User      *User     `gorm:"foreignKey:user_id;" json:"user,omitempty"`
}
