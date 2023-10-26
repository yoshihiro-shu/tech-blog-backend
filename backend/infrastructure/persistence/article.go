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

func (ap *articlePersistence) FindByID(article *model.Article, id int) error {
	return ap.Reprica().
		// Joins("User").
		Joins("Category").
		Preload("Tags").
		Where("articles.id = ?", id).
		First(article).
		Error
}

func (ap *articlePersistence) GetArticles(articles *[]model.Article, limit, offset int) error {
	return ap.Reprica().
		// Joins("User").
		Joins("Category").
		Preload("Tags").
		Order("created_at ASC").
		Limit(limit).
		Offset(offset).
		Find(articles).
		Error
}

func (ap *articlePersistence) GetPager() (int, error) {
	var count int64
	err := ap.Reprica().Model(&model.Article{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (ap *articlePersistence) GetArticlesByCategory(articles *[]model.Article, slug string) error {
	return ap.Reprica().
		// Preload("User").
		Preload("Category").
		Preload("Tags").
		Joins("LEFT JOIN categories AS category ON articles.category_id = category.id").
		Where("category.slug = ?", slug).
		Find(&articles).Error
}

// GetArticlesByTag retrieves articles based on a given tag slug.
//
// articles: a pointer to a slice of model.Article to store the retrieved articles.
// slug: the slug of the tag to filter the articles by.
// error: an error indicating if there was any issue retrieving the articles.
func (ap *articlePersistence) GetArticlesByTag(articles *[]model.Article, slug string) error {
	return ap.Reprica().
		// Preload("User").
		Preload("Category").
		Preload("Tags").
		Joins("JOIN article_tags ON articles.id = article_tags.article_id").
		Joins("JOIN tags ON tags.id = article_tags.tag_id").
		Where("tags.slug = ?", slug).
		Find(&articles).Error
}

func (ap *articlePersistence) Update(article *model.Article) (*model.Article, error) {
	return &model.Article{}, nil
}

func (ap *articlePersistence) Delete(article *model.Article) error {
	return nil
}
