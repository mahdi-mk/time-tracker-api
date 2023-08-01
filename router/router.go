package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/middlewares"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api", middlewares.Organizable)

	// ===========================================================
	// Version 1

	v1 := api.Group("/v1/")
	RegisterV1(v1)
}
