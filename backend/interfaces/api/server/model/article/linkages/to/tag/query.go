package article_linkages_to_tag

import (
	"time"

	"github.com/go-pg/pg"

	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/model/article/table"
)

type (
	Article struct {
		table.Article
		Tags []Tag `pg:"many2many:article_tags"`
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

func (a *Article) GetArticle(db *pg.DB) error {

	query := db.Model(a).
		Relation("Tags").
		Order("created_at ASC").
		Where("id = ?", a.Id)

	err := query.Select()
	if err != nil {
		return err
	}

	return nil
}
