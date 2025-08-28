package handler

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/handler/adaptor"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/oapi"
)

func (h Core) CreateAuthor(ctx context.Context, request oapi.CreateAuthorRequestObject) (oapi.CreateAuthorResponseObject, error) {

	var output oapi.Author
	if err := WithTx(ctx, h.Logger, func(ctx context.Context) error {
		input := input.NewCreateAuthor(
			request.Body.Name,
		)
		response, err := h.authorUsecase.Create(ctx, input)
		if err != nil {
			return err
		}
		output = adaptor.NewAuthor(response.Author)
		return nil
	}); err != nil {
		return oapi.CreateAuthordefaultJSONResponse{
			Body:       adaptor.NewError(err),
			StatusCode: adaptor.ErrorToStatusCode(err),
		}, nil
	}
	return oapi.CreateAuthor201JSONResponse{Author: output}, nil
}

func (h Core) ListAuthors(ctx context.Context, request oapi.ListAuthorsRequestObject) (oapi.ListAuthorsResponseObject, error) {
	var output []oapi.Author
	if err := WithTx(ctx, h.Logger, func(ctx context.Context) error {
		response, err := h.authorUsecase.List(ctx)
		if err != nil {
			return err
		}

		output = adaptor.NewAuthors(response.Authors)

		return nil
	}); err != nil {
		return oapi.ListAuthorsdefaultJSONResponse{
			Body:       adaptor.NewError(err),
			StatusCode: adaptor.ErrorToStatusCode(err),
		}, nil
	}
	return oapi.ListAuthors200JSONResponse{Authors: output}, nil
}

func (h Core) GetAuthorById(ctx context.Context, request oapi.GetAuthorByIdRequestObject) (oapi.GetAuthorByIdResponseObject, error) {
	var output oapi.Author
	if err := WithTx(ctx, h.Logger, func(ctx context.Context) error {
		// APIリクエストのデータを、ビジネスロジック用のデータ形式に変換
		input := input.NewGetAuthorByID(request.Id)
		// ユースケース層で著者取得の実際の処理を実行
		response, err := h.authorUsecase.GetByID(ctx, input)
		if err != nil {
			return err
		}
		// 出力データの変換
		output = adaptor.NewAuthor(response.Author)

		return nil

	}); err != nil {
		return oapi.GetAuthorByIddefaultJSONResponse{
			Body:       adaptor.NewError(err),
			StatusCode: adaptor.ErrorToStatusCode(err),
		}, nil
	}
	return oapi.GetAuthorById200JSONResponse{Author: output}, nil
}

func (h Core) SearchAuthors(ctx context.Context, request oapi.SearchAuthorsRequestObject) (oapi.SearchAuthorsResponseObject, error) {
	// TODO: 実装
	return oapi.SearchAuthors200JSONResponse{Authors: []oapi.Author{}}, nil
}