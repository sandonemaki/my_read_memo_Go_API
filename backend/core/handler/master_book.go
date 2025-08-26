package handler

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/oapi"
)

func (h Core) ListMasterBooks(ctx context.Context, request oapi.ListMasterBooksRequestObject) (oapi.ListMasterBooksResponseObject, error) {
	// TODO: 実装
	return oapi.ListMasterBooks200JSONResponse{}, nil
}

func (h Core) CreateMasterBook(ctx context.Context, request oapi.CreateMasterBookRequestObject) (oapi.CreateMasterBookResponseObject, error) {
	// TODO: 実装
	return oapi.CreateMasterBook201JSONResponse{}, nil
}

func (h Core) SearchMasterBooks(ctx context.Context, request oapi.SearchMasterBooksRequestObject) (oapi.SearchMasterBooksResponseObject, error) {
	// TODO: 実装
	return oapi.SearchMasterBooks200JSONResponse{}, nil
}

func (h Core) GetMasterBookById(ctx context.Context, request oapi.GetMasterBookByIdRequestObject) (oapi.GetMasterBookByIdResponseObject, error) {
	// TODO: 実装
	return oapi.GetMasterBookById200JSONResponse{}, nil
}

func (h Core) UpdateMasterBook(ctx context.Context, request oapi.UpdateMasterBookRequestObject) (oapi.UpdateMasterBookResponseObject, error) {
	// TODO: 実装
	return oapi.UpdateMasterBook200JSONResponse{}, nil
}