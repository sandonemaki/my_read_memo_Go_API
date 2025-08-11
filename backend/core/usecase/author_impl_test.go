package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
)

func TestMockCreateAuthor(t *testing.T) {

	const (
		TestName = "テスト著者"
		TestID   = int64(1)
	)
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	vectors := map[string]struct {
		params   input.CreateAuthor
		expected *output.CreateAuthor
		wantErr  error
		options  cmp.Options
		prepare  func(mock sqlmock.Sqlmock)
	}{
		"OK": {
			params: input.CreateAuthor{
				Name: TestName,
			},
			expected: &output.CreateAuthor{
				Author: &model.Author{
					ID:   TestID,
					Name: TestName,
				},
			},
			wantErr: nil,
			options: cmp.Options{
				cmpopts.IgnoreFields(output.CreateAuthor{}, "CreatedAt", "UpdatedAt"),
			},
			prepare: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).AddRow(TestID, TestName, fixedTime, fixedTime)
				insertQuery := `INSERT INTO authors \(name, created_at, updated_at\) VALUES \(\?, \?, \?\)`
				mock.ExpectQuery(insertQuery).
					WithArgs(TestName, fixedTime, fixedTime).
					WillReturnRows(rows)
			},
		},
		"ValidationError": {
			params: input.CreateAuthor{
				Name: "", // 空の名前は無効
			},
			expected: nil,
			wantErr:  errors.New("validation error"), // 入力検証エラーを期待
			options:  cmp.Options{},
			prepare: func(mock sqlmock.Sqlmock) {
				// SQLの期待値は設定しない（入力検証で早期リターンするため）
			},
		},
		"DatabaseError": {
			params: input.CreateAuthor{
				Name: TestName,
			},
			expected: nil,
			wantErr:  errors.New("database error"), // データベースエラーを期待
			options:  cmp.Options{},
			prepare: func(mock sqlmock.Sqlmock) {
				insertQuery := `INSERT INTO authors \(name, created_at, updated_at\) VALUES \(\?, \?, \?\)`
				mock.ExpectQuery(insertQuery).
					WithArgs(TestName, fixedTime, fixedTime).
					WillReturnError(errors.New("database error")) // データベースエラーを返す
			},
		},
	}

	// テストケースの実行
	for k, v := range vectors {
		t.Run(k, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("sqlmock error: %v", err)
			}
			defer db.Close()
			v.prepare(mock)

		})
	}
}
