package handler

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/handler/adaptor"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/oapi"
)

// - ハンドラー: API処理の制御
// - ユースケース: ビジネスロジック
// - アダプター: データ形式の変換

func (h Core) CreatePublisher(ctx context.Context, request oapi.CreatePublisherRequestObject) (oapi.CreatePublisherResponseObject, error) {
	var output oapi.Publisher
	if err := WithTx(ctx, h.Logger, func(ctx context.Context) error {
		// APIリクエストのデータを、ビジネスロジック用のデータ形式に変換
		input := input.NewCreatePublisher(
			request.Body.Name,
		)
		// ユースケース層でクライアント作成の実際の処理を実行
		// データベース保存、バリデーションなどを行う
		response, err := h.publisherUsecase.Create(ctx, input)
		if err != nil {
			return err
		}
		// 出力データの変換
		// ビジネスロジックの結果を、API用のデータ形式に変換
		output = adaptor.NewPublisher(response.Publisher)

		return nil

	}); err != nil {
		// エラー処理
		// エラーが発生した場合、適切なHTTPステータスコードとエラーメッセージを返す
		return oapi.CreatePublisherdefaultJSONResponse{
			Body:       adaptor.NewError(err),
			StatusCode: adaptor.ErrorToStatusCode(err),
		}, nil
	}
	// PublisherResponse型をPublisher型に変換して返却
	return oapi.CreatePublisher201JSONResponse{Publisher: output}, nil

}
