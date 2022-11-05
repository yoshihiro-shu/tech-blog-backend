package persistence

import (
	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/domain/repository"
)

type topPagePersistence struct {
	Conn *pg.DB
}

func NewTopPagePersistence(conn *pg.DB) repository.TopPageRepository {
	return &topPagePersistence{Conn: conn}
}

func (tp topPagePersistence) GetArticles(articles *[]model.Article) error {
	query := tp.Conn.Model(articles).
		Relation("Tags").
		Relation("Category").
		Relation("User").
		Order("created_at ASC")

	err := query.Select()
	if err != nil {
		return err
	}

	return nil
}

func (tp topPagePersistence) GetPager(a *model.Article) (int, error) {
	query := tp.Conn.Model(a)

	count, err := query.Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}
