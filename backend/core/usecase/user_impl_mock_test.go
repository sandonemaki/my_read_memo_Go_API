package usecase

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	queryImpl "github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/query"
	repositoryImpl "github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
)

func TestMockCreateUser(t *testing.T) {
	// 固定値を使用（動的値は禁止）
	const (
		TestUID         = "test_uid_001"
		TestDisplayName = "テストユーザー001"
		TestUlid        = "TEST123456789ABCDEF"
	)
	
	vectors := map[string]struct {
		params   input.CreateUser
		expected *output.CreateUser
		wantErr  error
		prepare  func(mock sqlmock.Sqlmock)
		options  cmp.Options
	}{
		"OK": {
			params: input.CreateUser{
				UID:         TestUID,
				DisplayName: TestDisplayName,
			},
			expected: &output.CreateUser{
				User: &model.User{
					Ulid:        TestUlid,
					UID:         TestUID,
					DisplayName: TestDisplayName,
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				// INSERT - 混在パターンをテスト（固定値 + AnyArg）
				insertQuery := `INSERT INTO "users" AS "users"("ulid", "display_name", "deleted_at", "created_at", "updated_at", "uid") VALUES ($1, $2, $3, DEFAULT, DEFAULT, $4) RETURNING "users"."ulid" AS "ulid", "users"."display_name" AS "display_name", "users"."deleted_at" AS "deleted_at", "users"."created_at" AS "created_at", "users"."updated_at" AS "updated_at", "users"."uid" AS "uid"`
				mock.ExpectExec(regexp.QuoteMeta(insertQuery)).
					WithArgs(sqlmock.AnyArg(), TestDisplayName, nil, TestUID).
					WillReturnResult(sqlmock.NewResult(1, 1))

				// GET - 全列を返す（時刻系は固定値）
				rows := sqlmock.NewRows([]string{"ulid", "display_name", "deleted_at", "created_at", "updated_at", "uid"}).AddRow(
					TestUlid,
					TestDisplayName,
					nil,
					time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), // 固定時刻
					time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), // 固定時刻
					TestUID)
				selectQuery := `SELECT "users"."ulid" AS "ulid", "users"."display_name" AS "display_name", "users"."deleted_at" AS "deleted_at", "users"."created_at" AS "created_at", "users"."updated_at" AS "updated_at", "users"."uid" AS "uid" FROM "users" AS "users" WHERE ("users"."uid" = $1)`
				mock.ExpectQuery(regexp.QuoteMeta(selectQuery)).
					WithArgs(TestUID).
					WillReturnRows(rows)
			},
			options: cmp.Options{
				cmpopts.IgnoreFields(output.CreateUser{}, "User.Ulid", "User.CreatedAt", "User.UpdatedAt"),
			},
		},
	}

	for k, v := range vectors {
		t.Run(k, func(t *testing.T) {
			sqlDB, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer sqlDB.Close()

			v.prepare(mock)

			// usecaseの依存関係を作成
			dbClient := db.NewClient(sqlDB)
			userQuery := queryImpl.NewUser(&dbClient)
			userRepo := repositoryImpl.NewUser(&dbClient)
			
			// 実際のusecaseを作成
			userUsecase := NewUser(userQuery, userRepo)
			
			// テスト実行
			actual, err := userUsecase.Create(context.Background(), v.params)

			// テスト結果の検証
			if err != nil {
				if v.wantErr == nil {
					t.Fatalf("unexpected error: %v", err)
				} else if v.wantErr.Error() != err.Error() {
					t.Fatalf("expected error: %v, got: %v", v.wantErr, err)
				}
			} else if v.wantErr != nil {
				t.Fatalf("expected an error, got none")
			}

			if !cmp.Equal(actual, v.expected, v.options...) {
				t.Errorf("unexpected result: %s", cmp.Diff(v.expected, actual, v.options...))
			}

			// モックの期待が満たされたかチェック
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("mock expectations were not met: %v", err)
			}
		})
	}
}
