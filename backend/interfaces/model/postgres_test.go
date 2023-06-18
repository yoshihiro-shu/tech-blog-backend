package model_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/mock_test"
)

const SelectOne = "SELECT 1;"

func TestDBClient(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer sqlDB.Close()
	db, err := mock_test.MockDB(sqlDB)
	if err != nil {
		t.Error(err.Error())
	}

	// SELECT 1 クエリのモックを設定
	rows := sqlmock.NewRows([]string{"number"}).AddRow(1)
	mock.ExpectQuery("^SELECT 1$").WillReturnRows(rows)

	// クエリの実行と結果の検証
	type Result struct {
		Number int
	}
	var result Result
	if err := db.Master().Raw("SELECT 1").Scan(&result).Error; err != nil {
		t.Fatalf("an error '%s' was not expected while querying", err)
	}

	if result.Number != 1 {
		t.Errorf("expected 1, but got %d", result.Number)
	}

	// すべてのモックが完了したことを確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
