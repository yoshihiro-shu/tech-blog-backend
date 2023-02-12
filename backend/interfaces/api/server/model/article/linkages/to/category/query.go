package article_linkages_to_category

import (
	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/model/article/table"
	"github.com/yoshihiro-shu/draft-backend/internal/config"
)

type Article struct {
	table.Article
	Category *Category `pg:"fk:category_id"`
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (a *Article) GetList(db *pg.DB) ([]Article, error) {
	articles := make([]Article, 10)

	query := db.Model(&articles).
		Relation("Category").
		Where("status = ?", config.StatusPublished).
		Order("created_at ASC").
		Limit(10)

	err := query.Select()
	if err != nil {
		return nil, err
	}

	return articles, nil
}
