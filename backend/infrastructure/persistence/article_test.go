package persistence_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/yoshihiro-shu/draft-backend/backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/backend/infrastructure/persistence"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/mock_test"
	"gorm.io/gorm"
)

func TestArticleCreate(t *testing.T) {
	orm, mock, err := mock_test.MockGorm()
	if err != nil {
		t.Fatalf("failed at mock db. err is %s", err.Error())
	}

	article_persistence := persistence.NewArticlePersistence(DummyHandler(orm), DummyHandler(orm))
	sampleNow := time.Now()

	type args struct {
		Article *model.Article
	}
	type want struct{}
	tests := []struct {
		Key  string
		Args args
		Want want
	}{
		{
			Key: "test1",
			Args: args{
				Article: &model.Article{
					Id:           1,
					UserId:       1,
					ThumbnailUrl: "hogehoge.com",
					Title:        "test1",
					Content:      "test content",
					Status:       1,
					CreatedAt:    sampleNow,
					UpdatedAt:    sampleNow,
					CategoryId:   1,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Key, func(t *testing.T) {

			// Set Mock SQL Qsuery
			mock.ExpectBegin()
			mock.ExpectQuery(`INSERT INTO "articles" (.+) RETURNING`).
				WithArgs(
					test.Args.Article.UserId,
					test.Args.Article.ThumbnailUrl,
					test.Args.Article.Title,
					test.Args.Article.Content,
					test.Args.Article.Status,
					test.Args.Article.CreatedAt,
					test.Args.Article.UpdatedAt,
					test.Args.Article.CategoryId,
					test.Args.Article.Id,
				).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			mock.ExpectCommit()

			// Execute Target function
			err := article_persistence.Create(test.Args.Article)
			if err != nil {
				t.Fatal(err)
			}

			// Check it out that Every Mock SQL is Executed.
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func DummyHandler(conn *gorm.DB) func() *gorm.DB {
	return func() *gorm.DB {
		return conn
	}
}
