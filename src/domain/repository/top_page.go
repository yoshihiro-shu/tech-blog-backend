package repository

import "github.com/yoshihiro-shu/tech-blog-backend/src/domain/model"

type TopPageRepository interface {
	GetArticles(*[]model.Article, int, int) error
	GetPager(*model.Article) (int, error)
}
