package entry

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/models"
	"github.com/mahdi-mk/time-tracker/app/requests"
	"github.com/mahdi-mk/time-tracker/database"
)

// Query returns a paginated list of entries
func Query(c *fiber.Ctx) error {
	var entries []models.Entry

	database.DB.Find(&entries)

	return c.JSON(entries)
}

// QueryByID returns a specific entry by its ID
func QueryByID(c *fiber.Ctx) error {
	var entry models.Entry
	database.DB.First(&entry, c.Params("id"))

	if entry.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Time entry not found",
		})
	}

	return c.JSON(entry)
}

// Create stores a new entry to the database
func Create(c *fiber.Ctx) error {
	entry, errors := new(requests.CreateOrUpdateEntryRequest).GetValidatedData(c)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	database.DB.Create(&entry)

	return c.Status(fiber.StatusCreated).JSON(entry)
}

// Update updates the data of the specified entry
func Update(c *fiber.Ctx) error {
	entry, errors := new(requests.CreateOrUpdateEntryRequest).GetValidatedData(c)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	database.DB.Model(&entry).Updates(entry)

	return c.JSON(entry)
}

// Delete deletes the specified entry from the database
func Delete(c *fiber.Ctx) error {
	var entry models.Entry
	database.DB.First(&entry, c.Params("id"))

	if entry.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Time entry not found",
		})
	}

	database.DB.Delete(&entry)

	return c.SendStatus(fiber.StatusNoContent)
}
