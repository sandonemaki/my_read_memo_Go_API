package usecase

import (
	"context"
	"errors"
	"regexp"
	"testing"

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
				// INSERT期待値
				insertQuery := `INSERT INTO "users" ("ulid","display_name","deleted_at","uid") VALUES ($1,$2,$3,$4)`
				mock.ExpectExec(regexp.QuoteMeta(insertQuery)).
					WithArgs(TestUlid, TestDisplayName, nil, TestUID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				
				// SELECT期待値（GetByUID）
				selectQuery := `SELECT "users".* FROM "users" WHERE ("users"."uid" = $1) LIMIT 1`
				rows := sqlmock.NewRows([]string{"ulid", "display_name", "deleted_at", "uid"}).
					AddRow(TestUlid, TestDisplayName, nil, TestUID)
				mock.ExpectQuery(regexp.QuoteMeta(selectQuery)).
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
				// INSERT でエラーを返す
				insertQuery := `INSERT INTO "users" ("ulid","display_name","deleted_at","uid") VALUES ($1,$2,$3,$4)`
				mock.ExpectExec(regexp.QuoteMeta(insertQuery)).
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
		},
		"UserNotFound": {
			params: input.GetCurrentUserDetail{
				UID: "not_exists_uid",
			},
			expected: nil,
			wantErr:  errors.New("user not found"),
			options:  cmp.Options{},
		},
	}

	for k, v := range vectors {
		t.Run(k, func(t *testing.T) {
			// gomockコントローラー作成
			ctrl := gomock.NewController(t)

			// モックインターフェース作成
			userQuery := query_mock.NewMockUser(ctrl)
			userRepo := repository_mock.NewMockUser(ctrl) // 使用しないが、NewUserに必要

			// モックの期待値設定
			// input→query変換が正しく行われるかを検証
			expectedQuery := query.UserGetQuery{
				UID: null.StringFrom(v.params.UID),
			}
			
			if v.wantErr != nil {
				// エラーケース: GetByUIDでエラーを返す
				userQuery.EXPECT().GetByUID(gomock.Any(), expectedQuery).Return(nil, v.wantErr)
			} else {
				// 正常ケース: GetByUIDでユーザーを返す
				userQuery.EXPECT().GetByUID(gomock.Any(), expectedQuery).Return(v.expected.User, nil)
			}

			// usecaseを作成
			userUsecase := NewUser(userQuery, userRepo)

			// テスト実行
			actual, err := userUsecase.GetCurrentUser(context.Background(), v.params)

			// テスト結果の検証
			if err != nil {
				if v.wantErr == nil {
					t.Fatalf("unexpected error: %v", err)
				} else if !errors.Is(err, v.wantErr) {
					t.Fatalf("expected error: %v, got: %v", v.wantErr, err)
				}
			} else if v.wantErr != nil {
				t.Fatalf("expected an error, got none")
			}

			if !cmp.Equal(actual, v.expected, v.options...) {
				t.Errorf("unexpected result: %s", cmp.Diff(v.expected, actual, v.options...))
			}
		})
	}
}
