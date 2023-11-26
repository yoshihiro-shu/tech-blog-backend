package usecase

import (
	"github.com/redis/go-redis/v9"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/domain/model"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/domain/repository"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/internal/pager"
)

type ArticlesUseCase interface {
	GetArticlesByCategory(articles *[]model.Article, slug string) error
	GetArticlesByTag(articles *[]model.Article, slug string) error
	GetPager(currentPage, offset int) (*pager.Pager, error)
}

type articlesUseCase struct {
	articleRepo       repository.ArticleRepository
	articlesCacheRepo repository.ArticlesCacheRepository
}

func NewArticlesUseCase(articleRepo repository.ArticleRepository, articlesCahceRepo repository.ArticlesCacheRepository) ArticlesUseCase {
	return &articlesUseCase{
		articleRepo:       articleRepo,
		articlesCacheRepo: articlesCahceRepo,
	}
}

func (au *articlesUseCase) GetArticlesByCategory(articles *[]model.Article, slug string) error {
	if err := au.articlesCacheRepo.GetByCategory(articles, slug); err == nil {
		return nil
	}
	err := au.articleRepo.GetArticlesByCategory(articles, slug)
	if err != nil {
		return err
	}

	if err := au.articlesCacheRepo.SetByCategory(articles, slug); err != nil {
		return err
	}
	return nil
}

func (au *articlesUseCase) GetArticlesByTag(articles *[]model.Article, slug string) error {
	if err := au.articlesCacheRepo.GetByTag(articles, slug); err == nil {
		return nil
	}
	err := au.articleRepo.GetArticlesByTag(articles, slug)
	if err != nil {
		return err
	}

	if err := au.articlesCacheRepo.SetByTag(articles, slug); err != nil {
		return err
	}
	return nil
}

func (au *articlesUseCase) GetPager(currentPage, offset int) (*pager.Pager, error) {
	var totalPager int

	err := au.articlesCacheRepo.GetTotalPager(&totalPager)
	if err != nil && err != redis.Nil {
		return nil, err
	} else {
		totalPager, err = au.articleRepo.GetPager()
		if err != nil {
			return nil, err
		}
		err = au.articlesCacheRepo.SetTotalPagerr(totalPager)
		if err != nil {
			// logを出力するエラーハンドリングに変えたい。
			return nil, err
		}
	}

	pager := pager.New(currentPage)
	pager.SetLastPage(offset, totalPager)

	return pager, nil
}
