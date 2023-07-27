package client

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/models"
	"github.com/mahdi-mk/time-tracker/database"
)

// Query returns a paginated list of clients
func Query(c *fiber.Ctx) error {
	var clients []models.Client

	database.DB.Find(&clients)

	return c.JSON(clients)
}

// QueryByID returns a specific client by its ID
func QueryByID(c *fiber.Ctx) error {
	var client models.Client
	database.DB.First(&client, c.Params("id"))

	if client.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Client Not Found",
		})
	}

	return c.JSON(client)
}

// Create stores a new client to the database
func Create(c *fiber.Ctx) error {
	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	database.DB.Create(&client)

	return c.Status(fiber.StatusCreated).JSON(client)
}

// Update updates the data of the specified client
func Update(c *fiber.Ctx) error {
	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	database.DB.Where("id = ?", c.Params("id")).Updates(&client)

	return c.JSON(client)
}

// Delete deletes the specified client from the database
func Delete(c *fiber.Ctx) error {
	database.DB.Delete(&models.Client{}, c.Params("id"))

	return c.SendStatus(fiber.StatusNoContent)
}
