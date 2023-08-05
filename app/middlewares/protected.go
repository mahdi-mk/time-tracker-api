package middlewares

import (
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/models"
	"github.com/mahdi-mk/time-tracker/support/auth"
	"github.com/mahdi-mk/time-tracker/support/env"
	"gorm.io/gorm"
)

func Protected(db *gorm.DB) func(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			var token models.Token
			db.Where("expired_at < ?", time.Now()).Or("expired_at IS NULL").First(&token, "token = ?", auth.GetToken(c))

			if token.ID == 0 {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Unauthorized",
				})
			}

			return c.Next()
		},
		SigningKey: jwtware.SigningKey{Key: []byte(env.Get("JWT_SECRET", ""))},
	})
}
