package repository

import "github.com/yoshihiro-shu/draft-backend/domain/model"

type ArticleRepository interface {
	Create(article *model.Article) (*model.Article, error)
	FindByID(id int) (*model.Article, error)
	Update(article *model.Article) (*model.Article, error)
	Delete(article *model.Article) error
}
