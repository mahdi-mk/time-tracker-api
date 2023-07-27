package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/handlers/v1/greeting"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// ===========================================================
	// Version 1

	v1 := api.Group("/v1")

	v1.Get("/hello", greeting.SayHello)
}
