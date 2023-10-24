package usecase

import (
	"github.com/yoshihiro-shu/draft-backend/backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/backend/domain/repository"
)

type ArticlesUseCase interface {
	GetArticlesByCategory(articles *[]model.Article, slug string) error
	GetArticlesByTag(articles *[]model.Article, slug string) error
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
