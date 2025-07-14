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

func TestMockCreateUser(t *testing.T) {
	// 固定値を使用（動的値は禁止）
	const (
		TestUID         = "test_uid_001"
		TestDisplayName = "テストユーザー001"
		TestUlid        = "TEST123456789ABCDEF"
	)
	
	// 固定時刻
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	vectors := map[string]struct {
		params   input.CreateUser
		expected *output.CreateUser
		wantErr  error
		options  cmp.Options
		prepare  func(mock sqlmock.Sqlmock)
	}{
		"OK": {
			params: input.CreateUser{
				Ulid:        TestUlid, // 固定ULIDを追加
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
			wantErr: nil,
			options: cmp.Options{
				cmpopts.IgnoreFields(output.CreateUser{}, "User.CreatedAt", "User.UpdatedAt"),
			},
			prepare: func(mock sqlmock.Sqlmock) {
				// INSERT期待値（Bob ORMの実際のSQL形式に合わせる）
				insertQuery := `INSERT INTO "users" AS "users"\("ulid", "display_name", "deleted_at", "created_at", "updated_at", "uid"\) VALUES \(\$1, \$2, \$3, DEFAULT, DEFAULT, \$4\) RETURNING`
				mock.ExpectExec(insertQuery).
					WithArgs(TestUlid, TestDisplayName, nil, TestUID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				
				// SELECT期待値（GetByUID）- Bob ORMの実際のSQL形式に合わせる
				selectQuery := `SELECT "users"\."ulid" AS "ulid", "users"\."display_name" AS "display_name", "users"\."deleted_at" AS "deleted_at", "users"\."created_at" AS "created_at", "users"\."updated_at" AS "updated_at", "users"\."uid" AS "uid" FROM "users" AS "users" WHERE \("users"\."uid" = \$1\)`
				rows := sqlmock.NewRows([]string{"ulid", "display_name", "deleted_at", "created_at", "updated_at", "uid"}).
					AddRow(TestUlid, TestDisplayName, nil, fixedTime, fixedTime, TestUID)
				mock.ExpectQuery(selectQuery).
					WithArgs(TestUID).
					WillReturnRows(rows)
			},
		},
		"DuplicateULID": {
			params: input.CreateUser{
				Ulid:        TestUlid,       // 同じULIDを使用
				UID:         "test_uid_002", // 異なるUID
				DisplayName: "テストユーザー002",
			},
			expected: nil,                                                // エラーの場合は期待値なし
			wantErr:  errors.New("UNIQUE constraint failed: users.ulid"), // 重複エラー
			options:  cmp.Options{},
			prepare: func(mock sqlmock.Sqlmock) {
				// INSERT でエラーを返す（Bob ORMの実際のSQL形式に合わせる）
				insertQuery := `INSERT INTO "users" AS "users"\("ulid", "display_name", "deleted_at", "created_at", "updated_at", "uid"\) VALUES \(\$1, \$2, \$3, DEFAULT, DEFAULT, \$4\) RETURNING`
				mock.ExpectExec(insertQuery).
					WithArgs(TestUlid, "テストユーザー002", nil, "test_uid_002").
					WillReturnError(errors.New("UNIQUE constraint failed: users.ulid"))
				
				// エラーケースではGetByUIDは呼ばれない
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

			// Wire DIでユースケース生成
			userUsecase := NewUserDI(db)

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

			if diff := cmp.Diff(actual, v.expected, v.options...); diff != "" {
				t.Errorf("Mismatch (-want +got):\n%s", diff)
			}

			// sqlmockの期待値が全て満たされたかチェック
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestMockGetCurrentUser(t *testing.T) {
	// 固定値を使用（動的値は禁止）
	const (
		TestUID         = "test_uid_001"
		TestDisplayName = "テストユーザー001"
		TestUlid        = "TEST123456789ABCDEF"
	)
	
	// 固定時刻
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	vectors := map[string]struct {
		params   input.GetCurrentUserDetail
		expected *output.GetUser
		wantErr  error
		options  cmp.Options
		prepare  func(mock sqlmock.Sqlmock)
	}{
		"OK": {
			params: input.GetCurrentUserDetail{
				UID: TestUID,
			},
			expected: &output.GetUser{
				User: &model.User{
					Ulid:        TestUlid,
					UID:         TestUID,
					DisplayName: TestDisplayName,
				},
			},
			wantErr: nil,
			options: cmp.Options{
				cmpopts.IgnoreFields(output.GetUser{}, "User.CreatedAt", "User.UpdatedAt"),
			},
			prepare: func(mock sqlmock.Sqlmock) {
				// SELECT期待値（GetByUID）- Bob ORMの実際のSQL形式に合わせる
				selectQuery := `SELECT "users"\."ulid" AS "ulid", "users"\."display_name" AS "display_name", "users"\."deleted_at" AS "deleted_at", "users"\."created_at" AS "created_at", "users"\."updated_at" AS "updated_at", "users"\."uid" AS "uid" FROM "users" AS "users" WHERE \("users"\."uid" = \$1\)`
				rows := sqlmock.NewRows([]string{"ulid", "display_name", "deleted_at", "created_at", "updated_at", "uid"}).
					AddRow(TestUlid, TestDisplayName, nil, fixedTime, fixedTime, TestUID)
				mock.ExpectQuery(selectQuery).
					WithArgs(TestUID).
					WillReturnRows(rows)
			},
		},
		"UserNotFound": {
			params: input.GetCurrentUserDetail{
				UID: "not_exists_uid",
			},
			expected: nil,
			wantErr:  errors.New("user not found"),
			options:  cmp.Options{},
			prepare: func(mock sqlmock.Sqlmock) {
				// SELECT期待値（ユーザーが見つからない場合）
				selectQuery := `SELECT "users"\."ulid" AS "ulid", "users"\."display_name" AS "display_name", "users"\."deleted_at" AS "deleted_at", "users"\."created_at" AS "created_at", "users"\."updated_at" AS "updated_at", "users"\."uid" AS "uid" FROM "users" AS "users" WHERE \("users"\."uid" = \$1\)`
				mock.ExpectQuery(selectQuery).
					WithArgs("not_exists_uid").
					WillReturnError(errors.New("user not found"))
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

			// Wire DIでユースケース生成
			userUsecase := NewUserDI(db)

			// テスト実行
			actual, err := userUsecase.GetCurrentUser(context.Background(), v.params)

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

			// sqlmockの期待値が全て満たされたかチェック
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("unfulfilled expectations: %s", err)
			}
		})
	}
}
