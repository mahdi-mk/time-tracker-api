package organization

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/models"
	"github.com/mahdi-mk/time-tracker/app/requests"
	"gorm.io/gorm"
)

type OrganizationController struct {
	db *gorm.DB
}

func MakeController(db *gorm.DB) *OrganizationController {
	return &OrganizationController{
		db: db,
	}
}

// Query returns a paginated list of organizations
func (cont *OrganizationController) Query(c *fiber.Ctx) error {
	var organizations []models.Organization

	cont.db.Find(&organizations)

	return c.JSON(organizations)
}

// QueryByID returns a specific organization by its ID
func (cont *OrganizationController) QueryByID(c *fiber.Ctx) error {
	var organization models.Organization
	cont.db.First(&organization, c.Params("id"))

	if organization.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Organization not found",
		})
	}

	return c.JSON(organization)
}

// Create stores a new organization to the database
func (cont *OrganizationController) Create(c *fiber.Ctx) error {
	organization, errors := new(requests.CreateOrUpdateOrganizationRequest).Validated(c)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	cont.db.Create(&organization)

	return c.Status(fiber.StatusCreated).JSON(organization)
}

// Update updates the data of the specified organization
func (cont *OrganizationController) Update(c *fiber.Ctx) error {
	organization, errors := new(requests.CreateOrUpdateOrganizationRequest).Validated(c)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	cont.db.Where(c.Params("id")).Updates(&organization)

	return c.JSON(organization)
}

// Delete deletes the specified organization from the database
func (cont *OrganizationController) Delete(c *fiber.Ctx) error {
	var organization models.Organization
	cont.db.First(&organization, c.Params("id"))

	if organization.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Organization Not Found",
		})
	}

	cont.db.Delete(&organization)

	return c.SendStatus(fiber.StatusNoContent)
}
