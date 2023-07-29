package utils

import "github.com/gofiber/fiber/v2"

func ParseBodyAndValidate(c *fiber.Ctx, body interface{}) error {
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	if errorsMap := Validate(body); errorsMap != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorsMap)
	}

	return nil
}
