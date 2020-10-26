package controllers

import (
	"fmt"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

// PayloadValidation validates request payload
func PayloadValidation(_request interface{}) error {
	validate = validator.New()

	err := validate.Struct(_request)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println("Validation error: ", err)
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.StructField())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
		}
	}

	return nil
}
