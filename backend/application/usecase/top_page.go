package usecase

import (
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/domain/repository"
)

type TopPageUseCase interface {
	GetArticles(*[]model.Article) error
}

type topPageUseCase struct {
	topPageRepo repository.TopPageRepository
}

func NewTopPageUseCase(topPagerepo repository.TopPageRepository) TopPageUseCase {
	return &topPageUseCase{topPageRepo: topPagerepo}
}

func (tp topPageUseCase) GetArticles(articles *[]model.Article) error {

	err := tp.topPageRepo.GetArticles(articles)
	if err != nil {
		return err
	}

	return nil
}
