package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validatorInstance = validator.New()

// ValidationErrors represents the map containing the errors
type ValidationErrors map[string]string

func ValidateRequest(c *fiber.Ctx, request interface{}) ValidationErrors {
	errors := make(ValidationErrors)

	if err := c.BodyParser(request); err != nil {
		errors["error"] = "Invalid body request"

		return errors
	}

	if err := validatorInstance.Struct(request); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("doesn't satisfy the `%v` constraint", err.Tag())
		}

		return errors
	}

	return nil
}
