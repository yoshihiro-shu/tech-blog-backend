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

func (ap *articlePersistence) GetArticlesByCategory(articles *[]model.Article, slug string) error {
	// return ap.Reprica().
	// 	Joins("User").
	// 	Joins("Category").
	// 	Preload("Tags").
	// 	Order("created_at ASC").
	// 	Where("category.slug = ?", slug).
	// 	Find(articles).
	// 	Error
	return nil
}

// GetArticlesByTag retrieves articles based on a given tag slug.
//
// articles: a pointer to a slice of model.Article to store the retrieved articles.
// slug: the slug of the tag to filter the articles by.
// error: an error indicating if there was any issue retrieving the articles.
func (ap *articlePersistence) GetArticlesByTag(articles *[]model.Article, slug string) error {
	query := `
		SELECT
			article.id,
			article.title,
			article.thumbnail_url,
			article.created_at,
			article.updated_at,
			category.id AS category_id,
			category.name AS category_name,
			category.slug AS category_slug,
			tag.id AS tag_id,
			-- tag.name AS tag_name,
			tag.slug AS tag_slug,
			article_tag.article_id AS article_tag_article_id,
			article_tag.tag_id AS article_tag_tag_id
		FROM
			articles AS article
		LEFT JOIN
			categories AS category
		ON
			category.id = article.category_id
		LEFT JOIN
			article_tags AS article_tag
		ON
			article_tag.article_id = article.id
		LEFT JOIN
			tags AS tag
		ON
			tag.id = article_tag.tag_id
		WHERE
			tag.slug = ?
		AND
			article.status = 2;
	`
	return ap.Reprica().Raw(query, slug).Scan(articles).Error
}

func (ap *articlePersistence) Update(article *model.Article) (*model.Article, error) {
	return &model.Article{}, nil
}

func (ap *articlePersistence) Delete(article *model.Article) error {
	return nil
}
