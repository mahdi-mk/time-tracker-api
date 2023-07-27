package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/database"
	"github.com/mahdi-mk/time-tracker/router"
)

func main() {
	app := fiber.New()

	router.RegisterRoutes(app)

	database.ConnectDB()
	database.RunMigrations()

	app.Listen(":8080")
}
