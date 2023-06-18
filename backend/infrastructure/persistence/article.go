package persistence

import (
	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/backend/domain/repository"
	"gorm.io/gorm"
)

type articlePersistence struct {
	Master  func() *pg.DB
	Reprica func() *pg.DB
	Primary func() *gorm.DB
}

func NewArticlePersistence(master, reprica func() *pg.DB, primary func() *gorm.DB) repository.ArticleRepository {
	return &articlePersistence{
		Master:  master,
		Reprica: reprica,
		Primary: primary,
	}
}

func (ap *articlePersistence) Create(article *model.Article) (*model.Article, error) {
	return &model.Article{}, nil
}

func (ap *articlePersistence) FindByID(article *model.Article) error {
	return ap.Primary().
		Joins("User").
		Joins("Category").
		Preload("Tags").
		Find(article).
		Error
}

func (ap *articlePersistence) GetArticles(articles *[]model.Article, limit, offset int) error {
	return ap.Primary().
		Joins("User").
		Joins("Category").
		Preload("Tags").
		Order("created_at ASC").
		Limit(limit).
		Offset(offset).
		Find(articles).
		Error
}

func (ap *articlePersistence) GetPager(article *model.Article) (int, error) {
	var count int64
	err := ap.Primary().Model(article).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (ap *articlePersistence) Update(article *model.Article) (*model.Article, error) {
	return &model.Article{}, nil
}

func (ap *articlePersistence) Delete(article *model.Article) error {
	return nil
}
