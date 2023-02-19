package persistence

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/domain/repository"
)

type refreshTokenPersistence struct {
	Master  func() *pg.DB
	Reprica func() *pg.DB
}

func NewRefreshTokenPersistence(master, reprica func() *pg.DB) repository.RefreshTokenRepository {
	return &refreshTokenPersistence{
		Master:  master,
		Reprica: reprica,
	}
}

func (rp *refreshTokenPersistence) Create(userId int, jwtId string) error {
	now := time.Now()
	token := &model.RefreshToken{
		UserId:    userId,
		JwtId:     jwtId,
		CreatedAt: now,
		UpdatedAt: now,
	}
	_, err := rp.Master().Model(token).Insert()
	if err != nil {
		return err
	}
	return nil
}
