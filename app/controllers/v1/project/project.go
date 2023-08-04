package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/models"
	"github.com/mahdi-mk/time-tracker/app/requests"
	"github.com/mahdi-mk/time-tracker/database"
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
			"error": "Project not found",
		})
	}

	return c.JSON(project)
}

// Create stores a new project to the database
func (cont *ProjectController) Create(c *fiber.Ctx) error {
	project, errors := new(requests.CreateOrUpdateProjectRequest).Validated(c)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	database.DB.Create(&project)

	return c.Status(fiber.StatusCreated).JSON(project)
}

// Update updates the data of the specified project
func (cont *ProjectController) Update(c *fiber.Ctx) error {
	project, errors := new(requests.CreateOrUpdateProjectRequest).Validated(c)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	cont.db.Where(c.Params("id")).Updates(&project)

	return c.JSON(project)
}

// Delete deletes the specified project from the database
func (cont *ProjectController) Delete(c *fiber.Ctx) error {
	var project models.Project
	database.DB.First(&project, c.Params("id"))

	if project.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Project not found",
		})
	}

	database.DB.Delete(&project)

	return c.SendStatus(fiber.StatusNoContent)
}
