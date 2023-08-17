package rinhaguard

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func Check(i interface{}) error {
	validate := validator.New()
	err := validate.Struct(i)
	if err != nil {
		var errMsg string
		for _, err := range err.(validator.ValidationErrors) {
			errMsg += err.Field() + " is invalid; "
		}
		return errors.New(errMsg)
	}

	return nil
}
