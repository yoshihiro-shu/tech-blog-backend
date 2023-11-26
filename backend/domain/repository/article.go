package repository

import "github.com/yoshihiro-shu/tech-blog-backend/backend/domain/model"

type ArticleRepository interface {
	Create(article *model.Article) (*model.Article, error)
	FindByID(article *model.Article, id int) error
	GetArticles(articles *[]model.Article, limit, offset int) error
	GetArticlesByCategory(articles *[]model.Article, slug string) error
	GetArticlesByTag(articles *[]model.Article, slug string) error
	GetPager() (int, error)
	Update(article *model.Article) (*model.Article, error)
	Delete(article *model.Article) error
}

type ArticleCacheRepository interface {
	GetArticleDetailById(article *model.Article, id int) error
	SetArticleDetailById(article model.Article, id int) error
}

type ArticlesCacheRepository interface {
	GetByCategory(articles *[]model.Article, slug string) error
	SetByCategory(articles *[]model.Article, slug string) error
	GetByTag(articles *[]model.Article, slug string) error
	SetByTag(articles *[]model.Article, slug string) error
	GetLastest(articles *[]model.Article, pageNumber int) error
	SetLastest(articles *[]model.Article, pageNumber int) error
	GetTotalPager(dest *int) error
	SetTotalPagerr(totalPager int) error
}
