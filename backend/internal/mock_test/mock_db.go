package mock_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/model"
)

// TODO 将来的にsqlmockを使ってDBのテストを自動化したい
func MockDB(t *testing.T) (*model.DBContext, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.NewWithDSN("postgres://user:pass@localhost:5432/mydb?sslmode=disable")
	if err != nil {
		t.Fatalf("can't create sqlmock: %s", err)
	}
	defer sqlDB.Close()
	// pgDB := pg.Connect(&pg.Options{
	// 	User:     "user",
	// 	Password: "pass",
	// 	Addr:     "localhost:5432",
	// 	Database: "mydb",
	// })
	// // Pingを使ってデータベース接続を確認
	// _, err = pgDB.Exec("SELECT 1")
	// if err != nil {
	// 	t.Fatalf("Failed to connect to the database.")
	// }
	// t.Log("Successfully connected to the database.")
	// return model.NewTest(pgDB), mock
	return &model.DBContext{}, mock
}
