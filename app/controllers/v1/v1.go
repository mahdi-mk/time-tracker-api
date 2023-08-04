package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/controllers/v1/auth"
	"github.com/mahdi-mk/time-tracker/app/controllers/v1/client"
	"github.com/mahdi-mk/time-tracker/app/controllers/v1/entry"
	"github.com/mahdi-mk/time-tracker/app/controllers/v1/organization"
	"github.com/mahdi-mk/time-tracker/app/controllers/v1/project"
	"github.com/mahdi-mk/time-tracker/app/middlewares"
	"github.com/mahdi-mk/time-tracker/support/router"
	"gorm.io/gorm"
)

func Register(group fiber.Router, db *gorm.DB) {

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
	router.Resource("organizations", protected, organization.MakeController(db))

	// Clients
	router.Resource("clients", protected, client.MakeController(db))

	// Projects
	router.Resource("projects", protected, project.MakeController(db))

	// Entries
	router.Resource("entries", protected, entry.MakeController(db))
}
