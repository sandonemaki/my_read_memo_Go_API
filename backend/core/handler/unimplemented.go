package handler

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/oapi"
)

// StrictServerInterface用のUnimplemented実装
type Unimplemented struct{}

func (_ Unimplemented) DeleteMe(ctx context.Context, request oapi.DeleteMeRequestObject) (oapi.DeleteMeResponseObject, error) {
	return oapi.DeleteMe500JSONResponse{
		InternalServerErrorJSONResponse: oapi.InternalServerErrorJSONResponse{
			Message: "unimplemented",
		},
	}, nil
}

func (_ Unimplemented) GetMe(ctx context.Context, request oapi.GetMeRequestObject) (oapi.GetMeResponseObject, error) {
	return oapi.GetMe500JSONResponse{
		InternalServerErrorJSONResponse: oapi.InternalServerErrorJSONResponse{
			Message: "unimplemented",
		},
	}, nil
}

func (_ Unimplemented) UpdateMe(ctx context.Context, request oapi.UpdateMeRequestObject) (oapi.UpdateMeResponseObject, error) {
	return oapi.UpdateMe500JSONResponse{
		InternalServerErrorJSONResponse: oapi.InternalServerErrorJSONResponse{
			Message: "unimplemented",
		},
	}, nil
}

func (_ Unimplemented) CreateAuthor(ctx context.Context, request oapi.CreateAuthorRequestObject) (oapi.CreateAuthorResponseObject, error) {
	return oapi.CreateAuthor500JSONResponse{
		InternalServerErrorJSONResponse: oapi.InternalServerErrorJSONResponse{
			Message: "unimplemented",
		},
	}, nil
}
func (_ Unimplemented) GetAuthorById(ctx context.Context, request oapi.GetAuthorByIdRequestObject) (oapi.GetAuthorByIdResponseObject, error) {
	return oapi.GetAuthorById500JSONResponse{
		InternalServerErrorJSONResponse: oapi.InternalServerErrorJSONResponse{
			Message: "unimplemented",
		},
	}, nil
}
func (_ Unimplemented) GetAuthors(ctx context.Context, request oapi.ListAuthorsRequestObject) (oapi.ListAuthorsResponseObject, error) {
	return oapi.GetAuthors500JSONResponse{
		InternalServerErrorJSONResponse: oapi.InternalServerErrorJSONResponse{
			Message: "unimplemented",
		},
	}, nil
}
func (_ Unimplemented) CreatePublisher(ctx context.Context, request oapi.CreatePublisherRequestObject) (oapi.CreatePublisherResponseObject, error) {
	return oapi.CreatePublisher500JSONResponse{
		InternalServerErrorJSONResponse: oapi.InternalServerErrorJSONResponse{
			Message: "unimplemented",
		},
	}, nil
}
func (_ Unimplemented) GetPublisherById(ctx context.Context, request oapi.GetPublisherByIdRequestObject) (oapi.GetPublisherByIdResponseObject, error) {
	return oapi.GetPublisherById500JSONResponse{
		InternalServerErrorJSONResponse: oapi.InternalServerErrorJSONResponse{
			Message: "unimplemented",
		},
	}, nil
}
func (_ Unimplemented) GetPublishers(ctx context.Context, request oapi.ListPublishersRequestObject) (oapi.ListPublishersResponseObject, error) {
	return oapi.GetPublishers500JSONResponse{
		InternalServerErrorJSONResponse: oapi.InternalServerErrorJSONResponse{
			Message: "unimplemented",
		},
	}, nil
}
