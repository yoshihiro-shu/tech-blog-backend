package persistence

import (
	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/domain/repository"
)

type articlePersistence struct {
	Conn *pg.DB
}

func NewArticlePersistence(conn *pg.DB) repository.ArticleRepository {
	return &articlePersistence{Conn: conn}
}

func (ap *articlePersistence) Create(article *model.Article) (*model.Article, error) {
	return &model.Article{}, nil
}

func (ap *articlePersistence) FindByID(id int) (*model.Article, error) {
	article := &model.Article{Id: id}
	query := ap.Conn.Model(article).WherePK()

	err := query.Select()
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (ap *articlePersistence) Update(article *model.Article) (*model.Article, error) {
	return &model.Article{}, nil
}

func (ap *articlePersistence) Delete(article *model.Article) error {
	return nil
}
