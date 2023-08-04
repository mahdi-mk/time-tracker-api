package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/mahdi-mk/time-tracker/database"
	"github.com/mahdi-mk/time-tracker/router"
	"github.com/mahdi-mk/time-tracker/support/env"
	"github.com/mahdi-mk/time-tracker/support/validator"
)

func main() {
	// Environment Variables Support
	env.Init()

	// Validation Support
	validator.Init()

	// Initialize the application
	app := fiber.New(fiber.Config{
		AppName: env.Get("APP_NAME", ""),
	})

	// Register all application routes
	router.RegisterRoutes(app)

	// Connect to database and run migrations
	database.ConnectDB()
	database.RunMigrations()

	// Start the server
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("Failed to start the server", err)
	}
}
