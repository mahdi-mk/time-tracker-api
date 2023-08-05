package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// GetToken retrieves the JWT token from the given context
func GetToken(c *fiber.Ctx) string {
	token := c.Locals("user").(*jwt.Token)

	return token.Raw
}

// GetID retrieves the authenticated user's ID from the given context
func GetID(c *fiber.Ctx) uint {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return uint(claims["uid"].(float64))
}
