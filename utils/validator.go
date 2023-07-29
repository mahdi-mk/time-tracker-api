package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validatorInstance = validator.New()

// Validate validates the given data, and returns validation errors if any
func Validate(data interface{}) fiber.Map {
	err := validatorInstance.Struct(data)

	if err != nil {
		var errors = make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("doesn't satisfy the `%v` constraint", err.Tag())
		}

		return fiber.Map{
			"status": fiber.StatusBadRequest,
			"errors": errors,
		}
	}

	return nil
}
