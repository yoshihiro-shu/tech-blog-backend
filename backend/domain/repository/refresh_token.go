package repository

import "time"

type RefreshTokenRepository interface {
	Create(userId int, jwtId string, expires time.Time) error
}
