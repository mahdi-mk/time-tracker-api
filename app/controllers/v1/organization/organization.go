package organization

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/models"
	"github.com/mahdi-mk/time-tracker/app/requests"
	"github.com/mahdi-mk/time-tracker/database"
	"github.com/mahdi-mk/time-tracker/utils"
	"github.com/mahdi-mk/time-tracker/utils/auth"
)

// Query returns a paginated list of organizations
func Query(c *fiber.Ctx) error {
	var organizations []models.Organization

	database.DB.Find(&organizations)

	return c.JSON(organizations)
}

// QueryByID returns a specific organization by its ID
func QueryByID(c *fiber.Ctx) error {
	var organization models.Organization
	database.DB.First(&organization, c.Params("id"))

	if organization.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Organization Not Found",
		})
	}

	return c.JSON(organization)
}

// Create stores a new organization to the database
func Create(c *fiber.Ctx) error {
	request := new(requests.CreateOrUpdateOrganization)

	if err := utils.ValidateRequest(c, request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	organization := models.Organization{
		Name:   request.Name,
		UserID: auth.GetID(c),
	}
	database.DB.Create(&organization)

	return c.Status(fiber.StatusCreated).JSON(organization)
}

// Update updates the data of the specified organization
func Update(c *fiber.Ctx) error {
	request := new(requests.CreateOrUpdateOrganization)
	var organization models.Organization

	database.DB.First(&organization, c.Params("id"))

	if organization.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Organization Not Found",
		})
	}

	if err := utils.ValidateRequest(c, request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	database.DB.Model(&organization).Updates(request)

	return c.JSON(organization)
}

// Delete deletes the specified organization from the database
func Delete(c *fiber.Ctx) error {
	var organization models.Organization
	database.DB.First(&organization, c.Params("id"))

	if organization.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Organization Not Found",
		})
	}

	database.DB.Delete(&organization)

	return c.SendStatus(fiber.StatusNoContent)
}
