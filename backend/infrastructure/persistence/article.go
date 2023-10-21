package persistence

import (
	"github.com/yoshihiro-shu/draft-backend/backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/backend/domain/repository"
	"gorm.io/gorm"
)

type articlePersistence struct {
	Master  func() *gorm.DB
	Reprica func() *gorm.DB
}

func NewArticlePersistence(master, reprica func() *gorm.DB) repository.ArticleRepository {
	return &articlePersistence{
		Master:  master,
		Reprica: reprica,
	}
}

func (ap *articlePersistence) Create(article *model.Article) (*model.Article, error) {
	return &model.Article{}, nil
}

func (ap *articlePersistence) FindByID(article *model.Article) error {
	return ap.Reprica().
		Joins("User").
		Joins("Category").
		Preload("Tags").
		First(article).
		Error
}

func (ap *articlePersistence) GetArticles(articles *[]model.Article, limit, offset int) error {
	return ap.Reprica().
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
	err := ap.Reprica().Model(article).Count(&count).Error
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
