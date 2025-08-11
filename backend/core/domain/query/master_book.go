package query

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/volatiletech/null"
)

// MasterBook はマスターブックのクエリを定義するインターフェース
// Public（大文字始まり）: 他のパッケージから使用するための契約として公開
// - infra層で実装される
// - usecase層から利用される
// - domain層の一部として、ビジネスロジックのクエリ要件を定義
type MasterBook interface {
	// List はフィルター条件に基づいてマスターブックの一覧を取得
	// - filter: 検索条件（MasterBookListFilter型、後で定義予定）
	// - output: 結果のマスターブック配列
	// - err: エラー情報
	List(ctx context.Context, filter MasterBookListFilter) (output []*model.MasterBook, err error)
	
	// GetByID は指定されたIDでマスターブックを取得
	// - query: 検索クエリ（ID指定）
	// - orFail: trueの場合、見つからない時にエラーを返す
	// - output: 見つかったマスターブック（orFail=falseで見つからない場合はnil）
	// - err: エラー情報
	GetByID(ctx context.Context, query MasterBookGetQuery, orFail bool) (output *model.MasterBook, err error)
	
	// Search はタイトルでマスターブックを検索
	// - query: 検索クエリ（タイトル検索）
	// - output: 検索結果のマスターブック配列
	// - err: エラー情報
	Search(ctx context.Context, query MasterBookSearchQuery) (output []*model.MasterBook, err error)
}

// MasterBookGetQuery はGetByIDメソッドで使用する検索クエリ
// Public（大文字始まり）: usecase層から使用されるため公開
// - ID指定での取得に使用
// - null.Int64でnull許容（IDが指定されない場合を考慮）
type MasterBookGetQuery struct {
	ID null.Int64 // Public: 外部から設定可能なフィールド
}

// MasterBookSearchQuery はSearchメソッドで使用する検索クエリ
// Public（大文字始まり）: usecase層から使用されるため公開
// - タイトル検索に使用
// - null.Stringでnull許容（検索条件が指定されない場合を考慮）
type MasterBookSearchQuery struct {
	Title null.String // Public: 外部から設定可能なフィールド
}
