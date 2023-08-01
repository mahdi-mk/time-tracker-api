package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

// This middleware ensures that every request has the "OrganizationID" header.
func Organizable(c *fiber.Ctx) error {
	id := c.Get("OrganizationID")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing Organization ID",
		})
	}

	c.Locals("OrgID", id)

	return c.Next()
}
