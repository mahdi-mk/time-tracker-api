package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func registerValidationRules(db *gorm.DB) {
	// Register your custom rules here
	validatorInstance.RegisterValidation("datetime_custom", customDateTimeRule)

	// This rule ensures that the value is a primary key,
	// and it exists in the provied table.
	validatorInstance.RegisterValidation("exists", func(fl validator.FieldLevel) bool {
		id := fl.Field().Uint()
		table := fl.Param()

		var entity struct{ ID uint }
		db.Table(table).First(&entity, "id = ?", id)

		return entity.ID != 0
	})
}

//------------------------------------------------------------------------------------
// Custom Validation Rules
//------------------------------------------------------------------------------------

func customDateTimeRule(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	_, err := time.Parse("2006-01-02 15:04:05.999999-07", value)

	return err == nil
}
