package validatorAdaptor

import (
	"github.com/go-playground/validator/v10"
)

type ParamValidator struct {
	validator *validator.Validate
}

func NewParamValidator() (ret *ParamValidator) {
	ret = &ParamValidator{validator: validator.New()}
	return ret
}

func (cv *ParamValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
