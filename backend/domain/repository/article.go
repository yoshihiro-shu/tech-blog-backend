package repository

import "github.com/yoshihiro-shu/draft-backend/domain/model"

type ArticleRepository interface {
	Create(article *model.Article) (*model.Article, error)
	FindByID(article *model.Article) error
	GetArticles(articles *[]model.Article, limit, offset int) error
	GetPager(article *model.Article) (int, error)
	Update(article *model.Article) (*model.Article, error)
	Delete(article *model.Article) error
}
