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

func handleErrors(err error) error {
	if err != nil {
		var fieldErrors []string
		for _, validationErr := range err.(validator.ValidationErrors) {
			fieldErrors = append(fieldErrors, validationErr.Field()+" is invalid; ")
		}
		return NewValidationError(fieldErrors...)
	}

	return nil
}

func Check(i interface{}) error {
	validate := validator.New()
	err := validate.Struct(i)
	return handleErrors(err)
}

func CheckUUID(id string) error {
	validate := validator.New()
	err := validate.Var(id, "uuid")
	return handleErrors(err)
}

func CheckSearchTerm(t string) error {
	validate := validator.New()
	err := validate.Var(t, "required")
	return handleErrors(err)
}

func NewValidationError(fields ...string) ValidationError {
	return ValidationError{FieldErrors: fields}
}
