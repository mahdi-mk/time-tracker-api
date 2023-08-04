package project

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/models"
	"github.com/mahdi-mk/time-tracker/app/requests"
	"github.com/mahdi-mk/time-tracker/database"
	"github.com/mahdi-mk/time-tracker/support/validator"
	"gorm.io/gorm"
)

type ProjectController struct {
	db *gorm.DB
}

func MakeController(db *gorm.DB) *ProjectController {
	return &ProjectController{
		db: db,
	}
}

// Query returns a paginated list of projects
func (cont *ProjectController) Query(c *fiber.Ctx) error {
	var projects []models.Project

	database.DB.Find(&projects)

	return c.JSON(projects)
}

// QueryByID returns a specific project by its ID
func (cont *ProjectController) QueryByID(c *fiber.Ctx) error {
	var project models.Project
	database.DB.First(&project, c.Params("id"))

	if project.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Project Not Found",
		})
	}

	return c.JSON(project)
}

// Create stores a new project to the database
func (cont *ProjectController) Create(c *fiber.Ctx) error {
	request := new(requests.CreateOrUpdateProject)

	if err := validator.ValidateRequest(c, request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if res := database.DB.First(&models.Client{}, request.ClientID); res.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Invalid Client ID",
		})
	}

	orgID, _ := strconv.ParseUint(c.Locals("OrgID").(string), 10, 32)
	project := models.Project{
		Name:           request.Name,
		Description:    request.Description,
		OrganizationID: uint(orgID),
		ClientID:       request.ClientID,
	}
	database.DB.Create(&project)

	return c.Status(fiber.StatusCreated).JSON(project)
}

// Update updates the data of the specified project
func (cont *ProjectController) Update(c *fiber.Ctx) error {
	request := new(requests.CreateOrUpdateProject)
	var project models.Project

	database.DB.First(&project, c.Params("id"))

	if project.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Project Not Found",
		})
	}

	if err := validator.ValidateRequest(c, request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	database.DB.Model(&project).Updates(request)

	return c.JSON(project)
}

// Delete deletes the specified project from the database
func (cont *ProjectController) Delete(c *fiber.Ctx) error {
	var project models.Project
	database.DB.First(&project, c.Params("id"))

	if project.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Project Not Found",
		})
	}

	database.DB.Delete(&project)

	return c.SendStatus(fiber.StatusNoContent)
}
