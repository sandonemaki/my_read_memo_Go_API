package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/query"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
)

// mockPublisherQuery はquery.Publisherインターフェースのモック実装
// SQLの詳細を隠蔽し、振る舞いのみを定義
type mockPublisherQuery struct {
	// GetByIDの振る舞いを定義する関数
	GetByIDFunc func(ctx context.Context, query query.PublisherGetQuery, orFail bool) (*model.Publisher, error)
	// Listの振る舞いを定義する関数
	ListFunc func(ctx context.Context) ([]*model.Publisher, error)
}

// GetByID はモックのGetByID実装
func (m *mockPublisherQuery) GetByID(ctx context.Context, q query.PublisherGetQuery, orFail bool) (*model.Publisher, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, q, orFail)
	}
	return nil, nil
}

// List はモックのList実装
func (m *mockPublisherQuery) List(ctx context.Context) ([]*model.Publisher, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx)
	}
	return nil, nil
}

// mockPublisherRepository はrepository.Publisherインターフェースのモック実装
type mockPublisherRepository struct {
	// Createの振る舞いを定義する関数
	CreateFunc func(ctx context.Context, publisher *model.Publisher) (int64, error)
	// UpdateByIDの振る舞いを定義する関数
	UpdateByIDFunc func(ctx context.Context, publisher *model.Publisher) error
	// DeleteByIDの振る舞いを定義する関数
	DeleteByIDFunc func(ctx context.Context, id int64, hardDelete bool) error
}

// Create はモックのCreate実装
func (m *mockPublisherRepository) Create(ctx context.Context, publisher *model.Publisher) (int64, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, publisher)
	}
	return 0, nil
}

// UpdateByID はモックのUpdateByID実装
func (m *mockPublisherRepository) UpdateByID(ctx context.Context, publisher *model.Publisher) error {
	if m.UpdateByIDFunc != nil {
		return m.UpdateByIDFunc(ctx, publisher)
	}
	return nil
}

// DeleteByID はモックのDeleteByID実装
func (m *mockPublisherRepository) DeleteByID(ctx context.Context, id int64, hardDelete bool) error {
	if m.DeleteByIDFunc != nil {
		return m.DeleteByIDFunc(ctx, id, hardDelete)
	}
	return nil
}

// TestMockCreatePublisher_WithInterfaceMock はインターフェースモックを使用したテスト
// SQLを一切書かず、ビジネスロジックのみをテスト
func TestMockCreatePublisher_WithInterfaceMock(t *testing.T) {
	// 固定値を使用（動的値は禁止）
	const (
		TestName = "テスト出版社"
		TestID   = int64(1)
	)

	vectors := map[string]struct {
		params   input.CreatePublisher
		expected *output.CreatePublisher
		wantErr  bool
		// モックの振る舞いを設定
		setupMock func() (*mockPublisherQuery, *mockPublisherRepository)
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
			wantErr: false,
			setupMock: func() (*mockPublisherQuery, *mockPublisherRepository) {
				mockQuery := &mockPublisherQuery{}
				mockRepo := &mockPublisherRepository{
					// Create時の振る舞い：引数のPublisherにIDを設定して成功
					CreateFunc: func(ctx context.Context, publisher *model.Publisher) (int64, error) {
						// 入力値の検証（SQLではなくビジネスロジックの検証）
						if publisher.Name != TestName {
							t.Errorf("unexpected name: got %s, want %s", publisher.Name, TestName)
						}
						// DBが自動採番するIDを返す
						return TestID, nil
					},
				}
				return mockQuery, mockRepo
			},
		},
		"EmptyName": {
			params: input.CreatePublisher{
				Name: "", // 空の名前（validation error）
			},
			expected: nil,
			wantErr:  true,
			setupMock: func() (*mockPublisherQuery, *mockPublisherRepository) {
				// バリデーションエラーの場合、repository.Createは呼ばれない
				mockQuery := &mockPublisherQuery{}
				mockRepo := &mockPublisherRepository{
					CreateFunc: func(ctx context.Context, publisher *model.Publisher) (int64, error) {
						t.Error("CreateFunc should not be called for validation error")
						return 0, nil
					},
				}
				return mockQuery, mockRepo
			},
		},
		"RepositoryError": {
			params: input.CreatePublisher{
				Name: TestName,
			},
			expected: nil,
			wantErr:  true,
			setupMock: func() (*mockPublisherQuery, *mockPublisherRepository) {
				mockQuery := &mockPublisherQuery{}
				mockRepo := &mockPublisherRepository{
					// リポジトリ層でエラーを返す
					CreateFunc: func(ctx context.Context, publisher *model.Publisher) (int64, error) {
						return 0, fmt.Errorf("repository error")
					},
				}
				return mockQuery, mockRepo
			},
		},
	}

	for k, v := range vectors {
		t.Run(k, func(t *testing.T) {
			// モックのセットアップ（SQLの詳細は一切含まない）
			mockQuery, mockRepo := v.setupMock()

			// UseCase層のインスタンス作成
			usecase := NewPublisher(mockQuery, mockRepo)

			// ビジネスロジックの実行
			actual, err := usecase.Create(context.Background(), v.params)

			// エラーの検証
			if v.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// 結果の検証（ビジネスロジックの出力のみ検証）
			if diff := cmp.Diff(v.expected, actual); diff != "" {
				t.Errorf("output mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// TestMockListPublisher_WithInterfaceMock は一覧取得のモックテスト
func TestMockListPublisher_WithInterfaceMock(t *testing.T) {
	// テスト用の固定データ
	testPublishers := []*model.Publisher{
		{ID: 1, Name: "講談社"},
		{ID: 2, Name: "集英社"},
		{ID: 3, Name: "小学館"},
	}

	vectors := map[string]struct {
		params    input.ListPublisher
		expected  *output.ListPublishers
		wantErr   bool
		setupMock func() (*mockPublisherQuery, *mockPublisherRepository)
	}{
		"OK": {
			params: input.ListPublisher{},
			expected: &output.ListPublishers{
				Publishers: testPublishers,
			},
			wantErr: false,
			setupMock: func() (*mockPublisherQuery, *mockPublisherRepository) {
				mockQuery := &mockPublisherQuery{
					// List時の振る舞い：固定のリストを返す
					// sqlmockではSELECT文を書くが、ここでは単純にデータを返す
					ListFunc: func(ctx context.Context) ([]*model.Publisher, error) {
						// フィルター条件の検証などの
						// ビジネスロジックをここでテスト可能
						return testPublishers, nil
					},
				}
				mockRepo := &mockPublisherRepository{}
				return mockQuery, mockRepo
			},
		},
		"Empty": {
			params: input.ListPublisher{},
			expected: &output.ListPublishers{
				Publishers: []*model.Publisher{},
			},
			wantErr: false,
			setupMock: func() (*mockPublisherQuery, *mockPublisherRepository) {
				mockQuery := &mockPublisherQuery{
					ListFunc: func(ctx context.Context) ([]*model.Publisher, error) {
						// 空のリストを返す
						return []*model.Publisher{}, nil
					},
				}
				mockRepo := &mockPublisherRepository{}
				return mockQuery, mockRepo
			},
		},
		"QueryError": {
			params:   input.ListPublisher{},
			expected: nil,
			wantErr:  true,
			setupMock: func() (*mockPublisherQuery, *mockPublisherRepository) {
				mockQuery := &mockPublisherQuery{
					ListFunc: func(ctx context.Context) ([]*model.Publisher, error) {
						return nil, fmt.Errorf("query error")
					},
				}
				mockRepo := &mockPublisherRepository{}
				return mockQuery, mockRepo
			},
		},
	}

	// ===== 各テストケースの実行 =====
	for k, v := range vectors {
		t.Run(k, func(t *testing.T) {
			// モックのセットアップ
			// シンプルな構造体の作成（sqlmock.New()のような複雑さはない）
			mockQuery, mockRepo := v.setupMock()

			// UseCase層のインスタンス作成
			// モックを依存性として注入
			usecase := NewPublisher(mockQuery, mockRepo)

			// ビジネスロジックの実行
			// SQLは実行されず、モックの振る舞いが呼ばれる
			actual, err := usecase.List(context.Background())

			// エラーの検証
			if v.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// 結果の検証
			// ビジネスロジックの出力のみを検証（SQLの検証は不要）
			if diff := cmp.Diff(v.expected, actual); diff != "" {
				t.Errorf("output mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// ================================================================================
// まとめ：sqlmock版とインターフェースモック版の比較
//
// 【sqlmock版の特徴】
// - SQL文を正確に記述する必要がある
// - Bob ORMの動作を含めて検証
// - エスケープや引数の型に注意が必要
// - DB層まで含めた統合的なテスト
//
// 【インターフェースモック版の特徴】
// - SQLを一切書かない
// - ビジネスロジックに集中
// - シンプルで理解しやすい
// - 高速で保守性が高い
//
// 【使い分け】
// - UseCase層テスト: インターフェースモック（推奨）
// - Infra層テスト: sqlmock
// - 統合テスト: 実DB
// ================================================================================
