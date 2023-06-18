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
	query := ap.Reprica().Model(articles).
		Relation("Tags").
		Relation("Category").
		Relation("User").
		Order("created_at ASC").
		Limit(limit).
		Offset(offset)

	err := query.Select()
	if err != nil {
		return err
	}

	if len(*articles) == 0 {
		return pg.ErrNoRows
	}

	return nil
}

func (ap *articlePersistence) GetPager(article *model.Article) (int, error) {
	query := ap.Reprica().Model(article)

	count, err := query.Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (ap *articlePersistence) Update(article *model.Article) (*model.Article, error) {
	return &model.Article{}, nil
}

func (ap *articlePersistence) Delete(article *model.Article) error {
	return nil
}
