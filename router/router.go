package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/middlewares"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api", middlewares.Organizable)

	// ===========================================================
	// Version 1

	v1 := api.Group("/v1/")
	registerV1(v1, db)
}
