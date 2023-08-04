package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/controllers/v1/auth"
	"github.com/mahdi-mk/time-tracker/app/controllers/v1/client"
	"github.com/mahdi-mk/time-tracker/app/controllers/v1/entry"
	"github.com/mahdi-mk/time-tracker/app/controllers/v1/organization"
	"github.com/mahdi-mk/time-tracker/app/controllers/v1/project"
	"github.com/mahdi-mk/time-tracker/app/middlewares"
	"gorm.io/gorm"
)

func registerV1(group fiber.Router, db *gorm.DB) {

	//======================================================
	// Public Routes

	// Authentication
	authController := auth.MakeController(db)
	group.Post("/auth/register/", authController.Register)
	group.Post("/auth/login/", authController.Login)

	//======================================================
	// Protected Routes

	protected := group.Use(middlewares.Protected())

	// Organizations
	organizationController := organization.MakeController(db)
	protected.Get("/organizations/", organizationController.Query)
	protected.Get("/organizations/:id", organizationController.QueryByID)
	protected.Post("/organizations/", organizationController.Create)
	protected.Patch("/organizations/:id", organizationController.Update)
	protected.Delete("/organizations/:id", organizationController.Delete)

	// Clients
	clientConrtoller := client.MakeController(db)
	protected.Get("/clients/", clientConrtoller.Query)
	protected.Get("/clients/:id", clientConrtoller.QueryByID)
	protected.Post("/clients/", clientConrtoller.Create)
	protected.Patch("/clients/:id", clientConrtoller.Update)
	protected.Delete("/clients/:id", clientConrtoller.Delete)

	// Projects
	projectController := project.MakeController(db)
	protected.Get("/projects/", projectController.Query)
	protected.Get("/projects/:id", projectController.QueryByID)
	protected.Post("/projects/", projectController.Create)
	protected.Patch("/projects/:id", projectController.Update)
	protected.Delete("/projects/:id", projectController.Delete)

	// Entries
	entryController := entry.MakeController(db)
	protected.Get("/entries/", entryController.Query)
	protected.Get("/entries/:id", entryController.QueryByID)
	protected.Post("/entries/", entryController.Create)
	protected.Patch("/entries/:id", entryController.Update)
	protected.Delete("/entries/:id", entryController.Delete)
}
