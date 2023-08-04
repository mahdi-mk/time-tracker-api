package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validatorInstance *validator.Validate

// RegisterValidator creates a new instance of the validator package,
// and registers all custom validation rules.
func RegisterValidator() {
	validatorInstance = validator.New()

	// Register your custom rules here
	validatorInstance.RegisterValidation("datetime_custom", customDateTimeRule)
}

//------------------------------------------------------------------------------------
// Custom Validation Rules
//------------------------------------------------------------------------------------

func customDateTimeRule(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	_, err := time.Parse("2006-01-02 15:04:05.999999-07", value)

	return err == nil
}
