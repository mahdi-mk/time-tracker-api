package validator

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validatorInstance *validator.Validate

// Init initializes the "validator" support.
func Init(db *gorm.DB) {
	validatorInstance = validator.New()

	registerValidationRules(db)
}
