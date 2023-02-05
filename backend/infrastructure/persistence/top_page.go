package persistence

import (
	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/domain/repository"
)

type topPagePersistence struct {
	WriteDB *pg.DB
	ReadDB  *pg.DB
}

func NewTopPagePersistence(primary, reprica *pg.DB) repository.TopPageRepository {
	return &topPagePersistence{
		WriteDB: primary,
		ReadDB:  reprica,
	}
}

func (tp topPagePersistence) GetArticles(articles *[]model.Article, limit, offset int) error {
	query := tp.ReadDB.Model(articles).
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

	return nil
}

func (tp topPagePersistence) GetPager(a *model.Article) (int, error) {
	query := tp.ReadDB.Model(a)

	count, err := query.Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}
