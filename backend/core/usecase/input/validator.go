package input

import (
	"github.com/go-playground/validator/v10"
	util "github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/util"
)

var validate *validator.Validate

func init() {
	validate = util.NewValidator()
}
