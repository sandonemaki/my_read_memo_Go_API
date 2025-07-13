package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	query_mock "github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/query/mock"
	repository_mock "github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/repository/mock"
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
		},
	}

	for k, v := range vectors {
		t.Run(k, func(t *testing.T) {
			// gomockコントローラー作成
			ctrl := gomock.NewController(t)

			// モックインターフェース作成
			userQuery := query_mock.NewMockUser(ctrl)
			userRepo := repository_mock.NewMockUser(ctrl)

			// モックの期待値設定
			if v.wantErr != nil {
				// エラーケース: Createでエラーを返し、GetByUIDは呼ばれない
				userRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(v.wantErr)
			} else {
				// 正常ケース: Createは成功、GetByUIDで結果を返す
				userRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
				userQuery.EXPECT().GetByUID(gomock.Any(), gomock.Any()).Return(v.expected.User, nil)
			}

			// usecaseを作成
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
		})
	}
}
