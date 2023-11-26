package usecase

import (
	"github.com/yoshihiro-shu/tech-blog-backend/src/domain/model"
	"github.com/yoshihiro-shu/tech-blog-backend/src/domain/repository"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/config"
)

type TwitterUseCase interface {
	GetTimelines(conf config.Configs) ([]model.Tweet, error)
}

type twitterUseCase struct {
	twitterRepo repository.TwitterRepository
}

func NewTwitterUseCase(twitterRepo repository.TwitterRepository) TwitterUseCase {
	return &twitterUseCase{twitterRepo: twitterRepo}
}

func (tu *twitterUseCase) GetTimelines(conf config.Configs) ([]model.Tweet, error) {
	res, err := tu.twitterRepo.GetTimelines(conf)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
