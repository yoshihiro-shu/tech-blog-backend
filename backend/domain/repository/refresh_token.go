package repository

import (
	"time"

	"github.com/yoshihiro-shu/tech-blog-backend/backend/domain/model"
)

type RefreshTokenRepository interface {
	Create(userId int, jwtId string, expires time.Time) error
	GetByJwtId(jwtId string) (*model.RefreshToken, error)
	Update(id int, jwtId string, expires time.Time) error
}
