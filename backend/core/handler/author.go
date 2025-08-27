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
