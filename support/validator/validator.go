package validator

import (
	"github.com/go-playground/validator/v10"
)

var validatorInstance *validator.Validate

// Init initializes the "validator" support.
func Init() {
	validatorInstance = validator.New()

	registerValidationRules()
}
