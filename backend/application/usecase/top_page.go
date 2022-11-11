package usecase

import (
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/domain/repository"
	"github.com/yoshihiro-shu/draft-backend/internal/pkg/pager"
)

type TopPageUseCase interface {
	GetArticles(*[]model.Article, int, int) error
	GetPager(currentPage, offset int) (*pager.Pager, error)
}

type topPageUseCase struct {
	topPageRepo repository.TopPageRepository
}

func NewTopPageUseCase(topPagerepo repository.TopPageRepository) TopPageUseCase {
	return &topPageUseCase{topPageRepo: topPagerepo}
}

func (tp topPageUseCase) GetArticles(articles *[]model.Article, limit, offset int) error {

	err := tp.topPageRepo.GetArticles(articles, limit, offset)
	if err != nil {
		return err
	}

	return nil
}

func (tp topPageUseCase) GetPager(currentPage, offset int) (*pager.Pager, error) {
	var a model.Article

	numOfArticles, err := tp.topPageRepo.GetPager(&a)
	if err != nil {
		return nil, err
	}

	pager := pager.New(currentPage)
	pager.SetLastPage(offset, numOfArticles)

	return pager, nil
}
