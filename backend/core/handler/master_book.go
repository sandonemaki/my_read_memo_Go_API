package handler

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/oapi"
)

func (h Core) ListMasterBooks(ctx context.Context, _ oapi.ListMasterBooksRequestObject) (oapi.ListMasterBooksResponseObject, error) {
	// TODO: 実装
	return oapi.ListMasterBooks200JSONResponse{Books: []oapi.MasterBook{}}, nil
}

func (h Core) CreateMasterBook(ctx context.Context, _ oapi.CreateMasterBookRequestObject) (oapi.CreateMasterBookResponseObject, error) {
	// TODO: 実装
	return oapi.CreateMasterBook201JSONResponse{}, nil
}

func (h Core) SearchMasterBooks(ctx context.Context, request oapi.SearchMasterBooksRequestObject) (oapi.SearchMasterBooksResponseObject, error) {
	// TODO: 実装
	return oapi.SearchMasterBooks200JSONResponse{Books: []oapi.MasterBook{}}, nil
}

func (h Core) GetMasterBookById(ctx context.Context, _ oapi.GetMasterBookByIdRequestObject) (oapi.GetMasterBookByIdResponseObject, error) {
	// TODO: 実装
	return oapi.GetMasterBookById200JSONResponse{}, nil
}

func (h Core) UpdateMasterBook(ctx context.Context, _ oapi.UpdateMasterBookRequestObject) (oapi.UpdateMasterBookResponseObject, error) {
	// TODO: 実装
	return oapi.UpdateMasterBook200JSONResponse{}, nil
}
