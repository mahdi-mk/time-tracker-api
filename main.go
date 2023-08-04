package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"github.com/mahdi-mk/time-tracker/database"
	"github.com/mahdi-mk/time-tracker/router"
	"github.com/mahdi-mk/time-tracker/support/env"
	"github.com/mahdi-mk/time-tracker/support/validator"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("An error occurred while loading environment variables", err)
	}

	// Initialize the application
	app := fiber.New(fiber.Config{
		AppName: env.Get("APP_NAME", ""),
	})

	// Validator Support
	validator.RegisterValidator()

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
