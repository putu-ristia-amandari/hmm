package common

import (
	"strings"

	validator "github.com/go-playground/validator/v10"
)

func ErrorResponse(err error) interface{} {
	validationErrors, ok := err.(validator.ValidationErrors)
	errors := interface{}(nil)
	errors = err.Error()
	if ok {
		errorValidate := map[string]string{}

		for _, ve := range validationErrors {
			errorValidate[strings.ToLower(ve.Field())] = strings.Split(ve.Error(), ":")[2]
		}

		errors = errorValidate
	}

	return errors
}
