package adaptor

import (
	"sync"

	v "github.com/go-playground/validator/v10"
)

var _validatorOnce sync.Once
var _validator *v.Validate

func validator() *v.Validate {
	_validatorOnce.Do(func() {
		_validator = v.New()
	})
	return _validator
}
