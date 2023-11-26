package postgres

import (
	"time"

	"github.com/yoshihiro-shu/tech-blog-backend/backend/domain/model"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/domain/repository"
	"gorm.io/gorm"
)

type refreshTokenPersistence struct {
	Master  func() *gorm.DB
	Reprica func() *gorm.DB
}

func NewRefreshTokenPersistence(master, reprica func() *gorm.DB) repository.RefreshTokenRepository {
	return &refreshTokenPersistence{
		Master:  master,
		Reprica: reprica,
	}
}

func (rp *refreshTokenPersistence) Create(userId int, jwtId string, expires time.Time) error {
	now := time.Now()
	token := &model.RefreshToken{
		UserId:    userId,
		JwtId:     jwtId,
		ExpiredAt: expires,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return rp.Master().Model(&model.RefreshToken{}).Create(token).Error
}

func (rp *refreshTokenPersistence) GetByJwtId(jwtId string) (*model.RefreshToken, error) {
	var rt model.RefreshToken
	err := rp.Reprica().Model(&model.RefreshToken{}).Where("jwt_id = ?", jwtId).First(&rt).Error
	if err != nil {
		return nil, err
	}
	return &rt, nil
}

func (rp *refreshTokenPersistence) Update(id int, jwtId string, expires time.Time) error {
	return rp.Master().Model(&model.RefreshToken{}).Where("id = ?", id).Updates(&model.RefreshToken{
		Id:        id,
		JwtId:     jwtId,
		ExpiredAt: expires,
		UpdatedAt: time.Now(),
	}).Error
}
