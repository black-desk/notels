package message

import (
	"context"
	"github.com/go-playground/validator/v10"
)

func Validate(v interface{}) error {
	return validate.Struct( v)
}

func ValidateCtx(ctx context.Context, v interface{}) error {
	return validate.StructCtx(ctx, v)
}

var validate = newValidator()

func newValidator() *validator.Validate {
	ret := validator.New()

	if err := ret.RegisterValidation("nullable",
		func(fl validator.FieldLevel) (ret bool) {
			return fl.Field().IsNil()
		},
	); err != nil {
		panic(err)
	}

	if err := ret.RegisterValidation("json_object",
		func(fl validator.FieldLevel) (ret bool) {
			bytes := fl.Field().Bytes()
			if len(bytes) > 0 {
				return bytes[0] == '{'
			} else {
				return false
			}
		},
	); err != nil {
		panic(err)
	}

	if err := ret.RegisterValidation("json_array",
		func(fl validator.FieldLevel) (ret bool) {
			bytes := fl.Field().Bytes()
			if len(bytes) > 0 {
				return bytes[0] == '['
			} else {
				return false
			}
		},
	); err != nil {
		panic(err)
	}

	return ret
}
