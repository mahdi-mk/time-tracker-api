package controllers

import (
	"github.com/gofiber/fiber/v2"
	v1 "github.com/mahdi-mk/time-tracker/app/controllers/v1"
	"github.com/mahdi-mk/time-tracker/app/middlewares"
	"gorm.io/gorm"
)

// Register registers all application routes so that controllers can be visible to the client.
func Register(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api", middlewares.Organizable)

	// ---------------------------------------------------------------------
	// Version 1
	// ---------------------------------------------------------------------

	v1.Register(
		api.Group("/v1/"),
		db,
	)

	// ---------------------------------------------------------------------
	// Version 2
	// ---------------------------------------------------------------------

	// Register "Version 2" routes here
	// ...
}
