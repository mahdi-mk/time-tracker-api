package client

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/models"
	"github.com/mahdi-mk/time-tracker/app/requests"
	"gorm.io/gorm"
)

type ClientController struct {
	db *gorm.DB
}

func MakeController(db *gorm.DB) *ClientController {
	return &ClientController{
		db: db,
	}
}

// Query returns a paginated list of clients
func (cont *ClientController) Query(c *fiber.Ctx) error {
	var clients []models.Client

	cont.db.Find(&clients)

	return c.JSON(clients)
}

// QueryByID returns a specific client by its ID
func (cont *ClientController) QueryByID(c *fiber.Ctx) error {
	var client models.Client
	cont.db.First(&client, c.Params("id"))

	if client.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Client not found",
		})
	}

	return c.JSON(client)
}

// Create stores a new client to the database
func (cont *ClientController) Create(c *fiber.Ctx) error {
	client, errors := new(requests.CreateOrUpdateClientRequest).Validated(c)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	cont.db.Create(&client)

	return c.Status(fiber.StatusCreated).JSON(client)
}

// Update updates the data of the specified client
func (cont *ClientController) Update(c *fiber.Ctx) error {
	client, errors := new(requests.CreateOrUpdateClientRequest).Validated(c)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	cont.db.Where(c.Params("id")).Updates(&client)

	return c.JSON(client)
}

// Delete deletes the specified client from the database
func (cont *ClientController) Delete(c *fiber.Ctx) error {
	var client models.Client
	cont.db.First(&client, c.Params("id"))

	if client.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Client not found",
		})
	}

	cont.db.Delete(&client)

	return c.SendStatus(fiber.StatusNoContent)
}
