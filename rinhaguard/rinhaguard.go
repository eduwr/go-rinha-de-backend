package rinhaguard

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	FieldErrors []string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("Validation errors: %v", e.FieldErrors)
}

func Check(i interface{}) error {
	validate := validator.New()
	err := validate.Struct(i)
	if err != nil {
		var fieldErrors []string
		for _, validationErr := range err.(validator.ValidationErrors) {
			fieldErrors = append(fieldErrors, validationErr.Field()+" is invalid; ")
		}
		return NewValidationError(fieldErrors...)
	}
	return nil
}

func CheckUUID(id string) error {
	validate := validator.New()
	err := validate.Var(id, "uuid")
	if err != nil {
		var fieldErrors []string
		for _, validationErr := range err.(validator.ValidationErrors) {
			fieldErrors = append(fieldErrors, validationErr.Field()+" is invalid; ")
		}
		return NewValidationError(fieldErrors...)
	}

	return nil
}

func NewValidationError(fields ...string) ValidationError {
	return ValidationError{FieldErrors: fields}
}
