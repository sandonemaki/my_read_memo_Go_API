package usecase

import (
	"context"
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

func TestMockCreatePublisher(t *testing.T) {
	// 固定値を使用（動的値は禁止）
	const (
		TestName = "テスト出版社"
		TestID   = int64(1)
	)
	
	// 固定時刻
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	vectors := map[string]struct {
		params   input.CreatePublisher
		expected *output.CreatePublisher
		wantErr  error
		options  cmp.Options
		prepare  func(mock sqlmock.Sqlmock)
	}{
		"OK": {
			params: input.CreatePublisher{
				Name: TestName,
			},
			expected: &output.CreatePublisher{
				Publisher: &model.Publisher{
					ID:   TestID,
					Name: TestName,
				},
			},
			wantErr: nil,
			options: cmp.Options{
				// 時刻フィールドは無視（DB側で自動設定されるため）
				cmpopts.IgnoreFields(output.CreatePublisher{}, "Publisher.CreatedAt", "Publisher.UpdatedAt"),
			},
			prepare: func(mock sqlmock.Sqlmock) {
				// INSERT期待値（Bob ORMのPublisher用SQL）
				insertQuery := `INSERT INTO "publishers" AS "publishers"\("name", "created_at", "updated_at"\) VALUES \(\$1, DEFAULT, DEFAULT\) RETURNING`
				mock.ExpectExec(insertQuery).
					WithArgs(TestName).
					WillReturnResult(sqlmock.NewResult(TestID, 1))
			},
		},
		"EmptyName": {
			params: input.CreatePublisher{
				Name: "", // 空の名前（validation error）
			},
			expected: nil,
			wantErr:  errors.New("validation error"), // バリデーションエラーを期待
			options:  cmp.Options{},
			prepare: func(mock sqlmock.Sqlmock) {
				// バリデーションエラーの場合はSQLは実行されない
			},
		},
		"DatabaseError": {
			params: input.CreatePublisher{
				Name: TestName,
			},
			expected: nil,
			wantErr:  errors.New("database error"),
			options:  cmp.Options{},
			prepare: func(mock sqlmock.Sqlmock) {
				// INSERT でエラーを返す
				insertQuery := `INSERT INTO "publishers" AS "publishers"\("name", "created_at", "updated_at"\) VALUES \(\$1, DEFAULT, DEFAULT\) RETURNING`
				mock.ExpectExec(insertQuery).
					WithArgs(TestName).
					WillReturnError(errors.New("database error"))
			},
		},
	}

	for k, v := range vectors {
		t.Run(k, func(t *testing.T) {
			// sqlmockセットアップ
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("sqlmock error: %v", err)
			}
			defer db.Close()

			// SQL期待値を設定
			v.prepare(mock)

			// TODO: 実際のテスト実行部分はPublisherDIが実装されてから追加
			// usecase := NewPublisherDI(db)
			// actual, err := usecase.Create(context.Background(), v.params)
			
			// 現在はコンパイルエラーを防ぐためコメントアウト
			t.Skip("TODO: PublisherDI実装後に有効化")

			// 期待値通りのSQLが実行されたかチェック
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("unfulfilled expectations: %v", err)
			}
		})
	}
}