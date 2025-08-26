package adaptor

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/errof"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/oapi"
)

// エラーの根本原因を取得して返却
func NewError(err error) oapi.Error {
	var code string
	switch e := errors.Cause(err).(type) {
	case errof.BadRequestErr:
		code = string(e)
	case errof.NotFoundErr:
		code = string(e)
	case errof.InternalErr:
		code = string(e)
	default:
		code = "Unknown"
	}
	return oapi.Error{
		Code:    code,
		Message: errors.Cause(err).Error(),
	}
}

func ErrorToStatusCode(err error) int {
	switch errors.Cause(err).(type) {
	case errof.BadRequestErr:
		return http.StatusBadRequest
	case errof.NotFoundErr:
		return http.StatusNotFound
	case errof.InternalErr:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
