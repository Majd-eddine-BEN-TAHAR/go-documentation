package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Setup initializes the validator
func Setup() {
    validate = validator.New()
}

// Validate performs the actual validation of structs
func Validate(i interface{}) error {
    return validate.Struct(i)
}