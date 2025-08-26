package usecase

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/query"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
	"github.com/volatiletech/null"
)

// publisher は実装構造体（private）
// なぜprivate？：外部から直接newできないようにして、必ずNewPublisher経由で作成させるため
type publisher struct {
	// フィールドもprivate：内部の依存性を隠蔽し、外部から直接触らせない
	publisherQuery query.Publisher      // データ取得用（SELECT）
	publisherRepo  repository.Publisher // データ変更用（INSERT/UPDATE）
}

// NewPublisher はコンストラクタ（Public）
// なぜPublic？：外部（main.goやwire）からインスタンスを作成するため
// 引数：必要な依存性を全て受け取る（依存性の注入）
// 戻り値：インターフェース型で返す（実装を隠蔽）
func NewPublisher(
	publisherQuery query.Publisher,
	publisherRepo repository.Publisher,
) Publisher { // Publisher インターフェースを返す（実装詳細を隠す）
	return &publisher{
		publisherQuery: publisherQuery,
		publisherRepo:  publisherRepo,
	}
}

// Create は新しい出版社を作成
func (u *publisher) Create(ctx context.Context, in input.CreatePublisher) (*output.CreatePublisher, error) {
	// Step 1: 入力値の検証
	// なぜ？：不正なデータがDBに入らないようにするため
	if err := in.Validate(); err != nil {
		return nil, err
	}

	// Step 2: ドメインモデルを作成
	// なぜmodel.Publisher？：ビジネスロジックを表現する型
	publisher := &model.Publisher{
		Name: in.Name,
		// ID, CreatedAt, UpdatedAtはDB側で自動設定される
	}

	// Step 3: repositoryのCreateを呼び出し（INSERT実行）
	// なぜrepository？：データの永続化はrepositoryの責任
	publisherID, err := u.publisherRepo.Create(ctx, publisher)
	if err != nil {
		return nil, err
	}
	publisher.ID = publisherID

	// Step 4: 出力を作成して返す
	// なぜoutput型？：APIレスポンスの構造を定義
	return output.NewCreatePublisher(publisher), nil
}

// GetByID は指定されたIDの出版社を取得
func (u *publisher) GetByID(ctx context.Context, in input.GetPublisherByID) (*output.GetPublisher, error) {
	// 入力値のバリデーション
	if err := in.Validate(); err != nil {
		return nil, err
	}

	// queryのGetByIDを呼び出し（SELECT実行）
	// なぜquery？：データ取得はqueryの責任
	// null.Int64From：null許容型に変換（DBのNULL対応）
	publisher, err := u.publisherQuery.GetByID(ctx, query.PublisherGetQuery{
		ID: null.Int64From(in.ID),
	}, true) // orFail: trueは必須データ（見つからない場合エラー）
	if err != nil {
		return nil, err
	}

	return output.NewGetPublisher(publisher), nil
}

// List は全ての出版社を取得
func (u *publisher) List(ctx context.Context) (*output.ListPublishers, error) {
	// queryのListを呼び出し
	publishers, err := u.publisherQuery.List(ctx)
	if err != nil {
		return nil, err
	}

	return output.NewListPublishers(publishers), nil
}

// SearchByName は名前で出版社を検索
func (u *publisher) SearchByName(ctx context.Context, name string) (*output.ListPublishers, error) {
	// TODO: query層にSearchByNameメソッドを追加後に実装
	// 現時点では未実装
	return output.NewListPublishers([]*model.Publisher{}), nil
}