package jsonrpc2

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

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

func messageAs(msg json.RawMessage, target interface{}) bool {
	err := json.Unmarshal(msg, target)
	if err != nil {
		log.Debugf(
			"try unmarshal %v to %v failed: %v",
			msg, target, err)
		return false
	}

	err = validate.Struct(target)
	if err != nil {
		log.Debugf("%v is not valid: %v", target, err)
		return false
	}

	return true
}
