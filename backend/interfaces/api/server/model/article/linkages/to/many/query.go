package article_linkages_to_many

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/model/article/table"
)

type (
	Article struct {
		table.Article
		Category *Category `pg:"fk:category_id"`
		Tags     []Tag     `pg:"many2many:article_tags"`
	}

	Category struct {
		Id          int       `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		ParentId    int       `json:"parentId"`
		CreatedAt   time.Time `json:"createdAt"`
	}

	Tag struct {
		Id          int       `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"createdAt"`
	}

	ArticleTags struct {
		ArticleId int `json:"article_id"`
		TagId     int `json:"tag_id"`
	}
)

func New(id int) *Article {
	return &Article{
		Article: *table.NewArticle(id),
	}
}

func (a *Article) GetArticle(db *pg.DB) error {

	query := db.Model(a).
		Relation("Tags").
		Relation("Category").
		Order("created_at ASC").
		Where("article.id = ?", a.Id)

	err := query.Select()
	if err != nil {
		return err
	}

	return nil
}

func GetArticleList(db *pg.DB, articles *[]Article) error {

	query := db.Model(articles).
		Relation("Tags").
		Relation("Category").
		Order("created_at ASC")

	err := query.Select()
	if err != nil {
		return err
	}

	return nil
}
