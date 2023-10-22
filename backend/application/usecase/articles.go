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
	articleRepo repository.ArticleRepository
}

func NewArticlesUseCase(articleRepo repository.ArticleRepository) ArticlesUseCase {
	return &articlesUseCase{articleRepo: articleRepo}
}

func (au *articlesUseCase) GetArticlesByCategory(articles *[]model.Article, slug string) error {
	err := au.articleRepo.GetArticlesByCategory(articles, slug)
	if err != nil {
		return err
	}

	return nil
}

func (au *articlesUseCase) GetArticlesByTag(articles *[]model.Article, slug string) error {
	return au.articleRepo.GetArticlesByTag(articles, slug)
}
