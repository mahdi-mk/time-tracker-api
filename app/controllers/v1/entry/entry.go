package entry

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/models"
	"github.com/mahdi-mk/time-tracker/app/requests"
	"gorm.io/gorm"
)

type EntryController struct {
	db *gorm.DB
}

func MakeController(db *gorm.DB) EntryController {
	return EntryController{
		db: db,
	}
}

// Query returns a paginated list of entries
func (cont *EntryController) Query(c *fiber.Ctx) error {
	var entries []models.Entry

	cont.db.Find(&entries)

	return c.JSON(entries)
}

// QueryByID returns a specific entry by its ID
func (cont *EntryController) QueryByID(c *fiber.Ctx) error {
	var entry models.Entry
	cont.db.First(&entry, c.Params("id"))

	if entry.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Time entry not found",
		})
	}

	return c.JSON(entry)
}

// Create stores a new entry to the database
func (cont *EntryController) Create(c *fiber.Ctx) error {
	entry, errors := new(requests.CreateOrUpdateEntryRequest).GetValidatedData(c)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	cont.db.Create(&entry)

	return c.Status(fiber.StatusCreated).JSON(entry)
}

// Update updates the data of the specified entry
func (cont *EntryController) Update(c *fiber.Ctx) error {
	entry, errors := new(requests.CreateOrUpdateEntryRequest).GetValidatedData(c)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	cont.db.Model(&entry).Updates(entry)

	return c.JSON(entry)
}

// Delete deletes the specified entry from the database
func (cont *EntryController) Delete(c *fiber.Ctx) error {
	var entry models.Entry
	cont.db.First(&entry, c.Params("id"))

	if entry.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Time entry not found",
		})
	}

	cont.db.Delete(&entry)

	return c.SendStatus(fiber.StatusNoContent)
}
