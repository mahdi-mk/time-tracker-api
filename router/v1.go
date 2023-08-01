package router

import (
	"github.com/gofiber/fiber/v2"
	authConrtoller "github.com/mahdi-mk/time-tracker/app/controllers/v1/auth"
	clientConrtoller "github.com/mahdi-mk/time-tracker/app/controllers/v1/client"
	organizationController "github.com/mahdi-mk/time-tracker/app/controllers/v1/organization"
	projectConrtoller "github.com/mahdi-mk/time-tracker/app/controllers/v1/project"
	"github.com/mahdi-mk/time-tracker/app/middlewares"
)

func RegisterV1(group fiber.Router) {

	//======================================================
	// Public Routes

	// Authentication
	authGrp := group.Group("/auth/")
	authGrp.Post("register", authConrtoller.Register)
	authGrp.Post("login", authConrtoller.Login)

	//======================================================
	// Protected Routes

	protected := group.Use(middlewares.Protected())

	// Organizations
	orgGrp := protected.Group("/organizations/")
	orgGrp.Get("", organizationController.Query)
	orgGrp.Get(":id", organizationController.QueryByID)
	orgGrp.Post("", organizationController.Create)
	orgGrp.Patch(":id", organizationController.Update)
	orgGrp.Delete(":id", organizationController.Delete)

	// Clients
	clientGrp := protected.Group("/clients/")
	clientGrp.Get("", clientConrtoller.Query)
	clientGrp.Get(":id", clientConrtoller.QueryByID)
	clientGrp.Post("", clientConrtoller.Create)
	clientGrp.Patch(":id", clientConrtoller.Update)
	clientGrp.Delete(":id", clientConrtoller.Delete)

	// Projects
	projectGrp := protected.Group("/projects/")
	projectGrp.Get("", projectConrtoller.Query)
	projectGrp.Get(":id", projectConrtoller.QueryByID)
	projectGrp.Post("", projectConrtoller.Create)
	projectGrp.Patch(":id", projectConrtoller.Update)
	projectGrp.Delete(":id", projectConrtoller.Delete)
}
