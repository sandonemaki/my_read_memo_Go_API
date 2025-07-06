package util

import (
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/volatiletech/null"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterCustomTypeFunc(NullValidator, null.String{}, null.Int{}, null.Int64{}, null.Bool{}, null.Float32{}, null.Float64{}, null.Byte{}, null.Bytes{}, null.Time{})
	return validate
}

func NullValidator(field reflect.Value) interface{} {
	switch field.Interface().(type) {
	case null.String:
		v := field.Interface().(null.String)
		if v.Valid {
			return v.String
		}
		return ""
	case null.Int64:
		v := field.Interface().(null.Int64)
		if v.Valid {
			return v.Int64
		}
		return 0
	case null.Bool:
		v := field.Interface().(null.Bool)
		if v.Valid {
			return v.Bool
		}
		return false
	case null.Float64:
		v := field.Interface().(null.Float64)
		if v.Valid {
			return v.Float64
		}
		return 0
	case null.Int:
		v := field.Interface().(null.Int)
		if v.Valid {
			return v.Int
		}
		return 0
	case null.Byte:
		v := field.Interface().(null.Byte)
		if v.Valid {
			return v.Byte
		}
		return nil
	case null.Bytes:
		v := field.Interface().(null.Bytes)
		if v.Valid {
			return v.Bytes
		}
		return nil
	case null.Time:
		v := field.Interface().(null.Time)
		if v.Valid {
			return v.Time
		}
		return time.Time{}
	default:
	}
	return nil
}
