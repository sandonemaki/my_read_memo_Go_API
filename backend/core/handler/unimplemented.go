package handler

import (
	"context"
	"fmt"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/handler/adaptor"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/oapi"
)

// StrictServerInterface用のUnimplemented実装
type Unimplemented struct{}

func (_ Unimplemented) DeleteMe(ctx context.Context, request oapi.DeleteMeRequestObject) (oapi.DeleteMeResponseObject, error) {
	return oapi.DeleteMedefaultJSONResponse{
		Body:       adaptor.NewError((fmt.Errorf("unimplemented"))),
		StatusCode: 500,
	}, nil
}

func (_ Unimplemented) GetMe(ctx context.Context, request oapi.GetMeRequestObject) (oapi.GetMeResponseObject, error) {
	return oapi.GetMedefaultJSONResponse{
		Body:       adaptor.NewError((fmt.Errorf("unimplemented"))),
		StatusCode: 404,
	}, nil
}

func (_ Unimplemented) UpdateMe(ctx context.Context, request oapi.UpdateMeRequestObject) (oapi.UpdateMeResponseObject, error) {
	return oapi.UpdateMedefaultJSONResponse{
		Body:       adaptor.NewError((fmt.Errorf("unimplemented"))),
		StatusCode: 404,
	}, nil
}

func (_ Unimplemented) CreateAuthor(ctx context.Context, request oapi.CreateAuthorRequestObject) (oapi.CreateAuthorResponseObject, error) {
	return oapi.CreateAuthordefaultJSONResponse{
		Body:       adaptor.NewError((fmt.Errorf("unimplemented"))),
		StatusCode: 500,
	}, nil
}
func (_ Unimplemented) GetAuthorById(ctx context.Context, request oapi.GetAuthorByIdRequestObject) (oapi.GetAuthorByIdResponseObject, error) {
	return oapi.GetAuthorByIddefaultJSONResponse{
		Body:       adaptor.NewError((fmt.Errorf("unimplemented"))),
		StatusCode: 404,
	}, nil
}
func (_ Unimplemented) GetAuthors(ctx context.Context, request oapi.ListAuthorsRequestObject) (oapi.ListAuthorsResponseObject, error) {
	return oapi.ListAuthorsdefaultJSONResponse{
		Body:       adaptor.NewError((fmt.Errorf("unimplemented"))),
		StatusCode: 404,
	}, nil
}

func (_ Unimplemented) SearchAuthors(ctx context.Context, request oapi.SearchAuthorsRequestObject) (oapi.SearchAuthorsResponseObject, error) {
	return oapi.SearchAuthorsdefaultJSONResponse{
		Body:       adaptor.NewError((fmt.Errorf("unimplemented"))),
		StatusCode: 404,
	}, nil
}
func (_ Unimplemented) CreatePublisher(ctx context.Context, request oapi.CreatePublisherRequestObject) (oapi.CreatePublisherResponseObject, error) {
	return oapi.CreatePublisherdefaultJSONResponse{
		Body:       adaptor.NewError((fmt.Errorf("unimplemented"))),
		StatusCode: 500,
	}, nil
}
func (_ Unimplemented) GetPublisherById(ctx context.Context, request oapi.GetPublisherByIdRequestObject) (oapi.GetPublisherByIdResponseObject, error) {
	return oapi.GetPublisherByIddefaultJSONResponse{
		Body:       adaptor.NewError((fmt.Errorf("unimplemented"))),
		StatusCode: 404,
	}, nil
}
func (_ Unimplemented) GetPublishers(ctx context.Context, request oapi.ListPublishersRequestObject) (oapi.ListPublishersResponseObject, error) {
	return oapi.ListPublishersdefaultJSONResponse{
		Body:       adaptor.NewError((fmt.Errorf("unimplemented"))),
		StatusCode: 404,
	}, nil
}

func (_ Unimplemented) SearchPublishers(ctx context.Context, request oapi.SearchPublishersRequestObject) (oapi.SearchPublishersResponseObject, error) {
	return oapi.SearchPublishersdefaultJSONResponse{
		Body:       adaptor.NewError((fmt.Errorf("unimplemented"))),
		StatusCode: 404,
	}, nil
}

func (_ Unimplemented) CreateMasterBook(ctx context.Context, request oapi.CreateBookRequestObject) (oapi.CreateBookResponseObject, error) {
	return oapi.CreateMasterBookdefaultJSONResponse{
		Body:       adaptor.NewError((fmt.Errorf("unimplemented"))),
		StatusCode: 500,
	}, nil
}
