package router

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// ===========================================================
	// Version 1

	v1 := api.Group("/v1/")
	RegisterV1(v1)
}
