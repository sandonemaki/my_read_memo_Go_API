package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/query"
	query_mock "github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/query/mock"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/repository"
	repository_mock "github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/repository/mock"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
	// testutil は削除：GoMockではtestutil.TxDBを使用しない
	// 理由：実DBを使用せず、インターフェースモックのみでテストするため
)

// TestMockCreatePublisher_WithGoMock は別アプリのパターンに従ったGoMockテスト
func TestMockCreatePublisher_WithGoMock(t *testing.T) {
	// 固定値を使用（動的値は禁止）
	const (
		TestName = "テスト出版社GoMock"
		TestID   = int64(1)
	)

	vectors := map[string]struct {
		params   input.CreatePublisher
		expected *output.CreatePublisher
		wantErr  error
		options  cmp.Options
		// ========== GoMockパターンの特徴 ==========
		// 【変更点1】setupMock関数：GoMockコントローラーを受け取り、モックを返す
		// 手動モックの場合：CreateFuncフィールドに関数を直接設定
		// GoMockの場合：gomock.Controllerでモックを作成し、EXPECT()で期待値設定
		setupMock func(ctrl *gomock.Controller) (query.Publisher, repository.Publisher)
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
			setupMock: func(ctrl *gomock.Controller) (query.Publisher, repository.Publisher) {
				// ========== GoMockの特徴 ==========
				// 【変更点2】自動生成されたモックを使用
				// 手動モック：type mockPublisherRepository struct { CreateFunc func(...) }
				// GoMock：mockgenコマンドで自動生成されたNewMockPublisher()を使用
				mockQuery := query_mock.NewMockPublisher(ctrl)
				mockRepo := repository_mock.NewMockPublisher(ctrl)

				// 【変更点3】EXPECT()チェーンで期待値設定
				// 手動モック：CreateFunc: func(...) { /* 直接実装 */ }
				// GoMock：EXPECT().Create().DoAndReturn()で期待値と振る舞いを設定
				mockRepo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, publisher *model.Publisher) (int64, error) {
						// 【共通点】引数検証の方法は同じ
						// 手動モック・GoMock共に関数内で引数をチェック
						if publisher.Name != TestName {
							t.Errorf("unexpected name: got %s, want %s", publisher.Name, TestName)
						}
						// DBが自動採番するIDを返す
						return TestID, nil
					}).
					Times(1) // 【GoMock独自】呼び出し回数の厳密な管理

				// 【GoMock独自】期待値を設定していないメソッドが呼ばれるとエラー
				// 手動モック：関数が設定されていなければnilを返すだけ
				// GoMock：EXPECTしていないメソッドが呼ばれると自動でテスト失敗

				return mockQuery, mockRepo
			},
		},
		"ValidationError": {
			params: input.CreatePublisher{
				Name: "", // 空の名前（validation error）
			},
			expected: nil,
			wantErr:  errors.New("Key: 'CreatePublisher.Name' Error:Field validation for 'Name' failed on the 'required' tag: 不正な入力エラー"), // バリデーションエラー
			setupMock: func(ctrl *gomock.Controller) (query.Publisher, repository.Publisher) {
				mockQuery := query_mock.NewMockPublisher(ctrl)
				mockRepo := repository_mock.NewMockPublisher(ctrl)

				// 【GoMock独自】期待値を設定しないことで「呼ばれないこと」を検証
				// 手動モック：関数を設定せず、呼ばれたら明示的にt.Error()
				// GoMock：EXPECT()を設定しないことで自動検証
				// バリデーションエラーの場合、repositoryは呼ばれない想定

				return mockQuery, mockRepo
			},
		},
	}

	// ========== 【重要な変更点4】テスト実行方法 ==========
	// 手動モック：通常のt.Run()でテスト実行
	// GoMock：通常のt.Run()でテスト実行（testutil.TxDBは不要）
	// 理由：実DBを使わないため、トランザクション管理が不要
	for k, v := range vectors {
		t.Run(k, func(t *testing.T) {
			// 【変更点5】GoMockコントローラーの作成
			// 手動モック：不要
			// GoMock：gomock.NewController(t)でコントローラー作成が必須
			ctrl := gomock.NewController(t)
			defer ctrl.Finish() // テスト終了時に期待値の検証を実行

			// 【変更点6】モックのセットアップ
			// 手動モック：構造体とフィールドに関数を直接設定
			// GoMock：setupMock関数でコントローラーを渡してモックを取得
			mockQuery, mockRepo := v.setupMock(ctrl)

			// 【共通点】UseCase層のインスタンス作成方法は同じ
			// 手動モック・GoMock共にコンストラクター注入
			usecase := NewPublisher(mockQuery, mockRepo)

			// 【共通点】ビジネスロジックの実行方法も同じ
			actual, err := usecase.Create(context.Background(), v.params)

			// 【共通点】エラー検証方法も同じ
			if v.wantErr != nil {
				if err == nil {
					t.Errorf("expected error but got nil")
				} else if err.Error() != v.wantErr.Error() {
					t.Errorf("error mismatch: got %v, want %v", err, v.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// 【共通点】結果の検証方法も同じ
			if diff := cmp.Diff(v.expected, actual); diff != "" {
				t.Errorf("output mismatch (-want +got):\n%s", diff)
			}

			// 【GoMock独自】ctrl.Finish()で期待値の検証が自動実行される
			// 手動モック：手動で検証を書く必要がある
			// GoMock：defer ctrl.Finish()で自動検証
		})
	}
}

// TestMockListPublisher_WithGoMock は一覧取得のGoMockテスト
func TestMockListPublisher_WithGoMock(t *testing.T) {
	// テスト用の固定データ
	testPublishers := []*model.Publisher{
		{ID: 1, Name: "講談社"},
		{ID: 2, Name: "集英社"},
		{ID: 3, Name: "小学館"},
	}

	vectors := map[string]struct {
		params    input.ListPublisher
		expected  *output.ListPublishers
		wantErr   error
		setupMock func(ctrl *gomock.Controller) (query.Publisher, repository.Publisher)
	}{
		"OK": {
			params: input.ListPublisher{},
			expected: &output.ListPublishers{
				Publishers: testPublishers,
			},
			wantErr: nil,
			setupMock: func(ctrl *gomock.Controller) (query.Publisher, repository.Publisher) {
				mockQuery := query_mock.NewMockPublisher(ctrl)
				mockRepo := repository_mock.NewMockPublisher(ctrl)

				// List操作はQueryインターフェースを使用
				mockQuery.EXPECT().
					List(gomock.Any()).
					Return(testPublishers, nil).
					Times(1)

				return mockQuery, mockRepo
			},
		},
		"Empty": {
			params: input.ListPublisher{},
			expected: &output.ListPublishers{
				Publishers: []*model.Publisher{},
			},
			wantErr: nil,
			setupMock: func(ctrl *gomock.Controller) (query.Publisher, repository.Publisher) {
				mockQuery := query_mock.NewMockPublisher(ctrl)
				mockRepo := repository_mock.NewMockPublisher(ctrl)

				// 空のリストを返すパターン
				mockQuery.EXPECT().
					List(gomock.Any()).
					Return([]*model.Publisher{}, nil).
					Times(1)

				return mockQuery, mockRepo
			},
		},
	}

	// ========== GoMockパターンのList実装 ==========
	for k, v := range vectors {
		t.Run(k, func(t *testing.T) {
			// GoMockコントローラーを作成
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// モックをセットアップ
			mockQuery, mockRepo := v.setupMock(ctrl)

			// UseCase層のインスタンス作成
			usecase := NewPublisher(mockQuery, mockRepo)

			// ビジネスロジックの実行
			actual, err := usecase.List(context.Background())

			// エラー検証
			if v.wantErr != nil {
				if err == nil {
					t.Errorf("expected error but got nil")
				} else if err.Error() != v.wantErr.Error() {
					t.Errorf("error mismatch: got %v, want %v", err, v.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// 結果の検証
			if diff := cmp.Diff(v.expected, actual); diff != "" {
				t.Errorf("output mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// ================================================================================
// 【完全対比】GoMock版 vs 手動モック版の具体的な違い
// ================================================================================
//
// ========== 【1. インポート文の違い】 ==========
// 手動モック：
// import (
//     "context"
//     "testing"
//     // testutilなし、gomockなし
// )
//
// GoMock：
// import (
//     "context"
//     "testing"
//     "github.com/golang/mock/gomock"                              // ← 追加
//     query_mock "github.com/.../core/domain/query/mock"          // ← 自動生成
//     repository_mock "github.com/.../core/domain/repository/mock"// ← 自動生成
// )
//
// ========== 【2. モック定義の違い】 ==========
// 手動モック：
// type mockPublisherRepository struct {
//     CreateFunc func(ctx context.Context, publisher *model.Publisher) error
// }
// func (m *mockPublisherRepository) Create(ctx context.Context, publisher *model.Publisher) error {
//     if m.CreateFunc != nil { return m.CreateFunc(ctx, publisher) }
//     return nil
// }
//
// GoMock：
// // 自動生成されたファイル（mock/mock_publisher.go）を使用
// // 手動定義は一切不要
//
// ========== 【3. テスト構造体の違い】 ==========
// 手動モック：
// vectors := map[string]struct {
//     setupMock func() (*mockPublisherQuery, *mockPublisherRepository)
// }
//
// GoMock：
// vectors := map[string]struct {
//     setupMock func(ctrl *gomock.Controller) (query.Publisher, repository.Publisher)
//                    // ↑ コントローラーを受け取る        ↑ インターフェース型を返す
// }
//
// ========== 【4. モック作成の違い】 ==========
// 手動モック：
// setupMock: func() (*mockPublisherQuery, *mockPublisherRepository) {
//     mockRepo := &mockPublisherRepository{
//         CreateFunc: func(ctx context.Context, publisher *model.Publisher) error {
//             if publisher.Name != TestName {
//                 t.Errorf("unexpected name")
//             }
//             publisher.ID = TestID
//             return nil
//         },
//     }
//     return mockQuery, mockRepo
// }
//
// GoMock：
// setupMock: func(ctrl *gomock.Controller) (query.Publisher, repository.Publisher) {
//     mockRepo := repository_mock.NewMockPublisher(ctrl)
//     mockRepo.EXPECT().
//         Create(gomock.Any(), gomock.Any()).
//         DoAndReturn(func(ctx context.Context, publisher *model.Publisher) error {
//             if publisher.Name != TestName {
//                 t.Errorf("unexpected name")
//             }
//             publisher.ID = TestID
//             return nil
//         }).
//         Times(1)
//     return mockQuery, mockRepo
// }
//
// ========== 【5. テスト実行の違い】 ==========
// 手動モック：
// for k, v := range vectors {
//     t.Run(k, func(t *testing.T) {
//         mockQuery, mockRepo := v.setupMock()
//         usecase := NewPublisher(mockQuery, mockRepo)
//         actual, err := usecase.Create(context.Background(), v.params)
//         // 手動で結果検証
//     })
// }
//
// GoMock：
// for k, v := range vectors {
//     t.Run(k, func(t *testing.T) {
//         ctrl := gomock.NewController(t)      // ← コントローラー作成
//         defer ctrl.Finish()                  // ← 自動検証
//         mockQuery, mockRepo := v.setupMock(ctrl)
//         usecase := NewPublisher(mockQuery, mockRepo)
//         actual, err := usecase.Create(context.Background(), v.params)
//         // ctrl.Finish()で期待値が自動検証される
//     })
// }
//
// ========== 【6. 期待値設定の違い】 ==========
// 手動モック：直接的な関数実装
// - CreateFunc: func(...) { /* 処理を直接書く */ }
// - 呼び出し回数チェックは手動実装
// - 引数チェックは関数内でif文
//
// GoMock：宣言的な期待値設定
// - EXPECT().Create(引数マッチャー).DoAndReturn(実装).Times(回数)
// - 呼び出し回数は自動チェック
// - 引数チェックはgomock.Eq()またはDoAndReturn内
//
// ========== 【7. 成功するために必要な変更点】 ==========
// 手動モック → GoMock への移行：
// 1. gomock import追加
// 2. 手動モック構造体削除
// 3. setupMock関数をctrl引数ありに変更
// 4. NewMockXXX(ctrl)でモック作成
// 5. EXPECT()チェーンで期待値設定
// 6. gomock.NewController(t)とdefer ctrl.Finish()追加
// 7. testutil.TxDBを通常のt.Run()に変更
//
// ========== 【まとめ】 ==========
// 手動モック：シンプル、直感的、学習コストが低い
// GoMock：自動生成、型安全、呼び出し回数の厳密管理、業界標準
// ================================================================================
