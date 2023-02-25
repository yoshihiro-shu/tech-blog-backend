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

func (rp *refreshTokenPersistence) Create(userId int, jwtId string, expires time.Time) error {
	now := time.Now()
	token := &model.RefreshToken{
		UserId:    userId,
		JwtId:     jwtId,
		ExpiredAt: expires,
		CreatedAt: now,
		UpdatedAt: now,
	}
	_, err := rp.Master().Model(token).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (rp *refreshTokenPersistence) GetByJwtId(jwtId string) (*model.RefreshToken, error) {
	var rt model.RefreshToken
	query := rp.Reprica().Model(&rt).
		Where("jwt_id = ?", jwtId)

	err := query.Select()
	if err != nil {
		return nil, err
	}

	return &rt, nil
}

func (rp *refreshTokenPersistence) Update(id int, jwtId string, expires time.Time) error {
	now := time.Now()
	rf := &model.RefreshToken{
		Id:        id,
		JwtId:     jwtId,
		ExpiredAt: expires,
		UpdatedAt: now,
	}
	_, err := rp.Master().
		Model(rf).
		Column("id", "jwt_id", "expired_at", "updated_at").
		WherePK().
		Update()

	if err != nil {
		return err
	}

	return nil
}
