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