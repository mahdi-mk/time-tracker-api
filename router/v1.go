package router

import (
	"github.com/gofiber/fiber/v2"
	authGrpV1 "github.com/mahdi-mk/time-tracker/app/controllers/v1/auth"
	clientGrpV1 "github.com/mahdi-mk/time-tracker/app/controllers/v1/client"
	organizationController "github.com/mahdi-mk/time-tracker/app/controllers/v1/organization"
	projectGrpV1 "github.com/mahdi-mk/time-tracker/app/controllers/v1/project"
	"github.com/mahdi-mk/time-tracker/app/middlewares"
)

func RegisterV1(group fiber.Router) {

	//======================================================
	// Protected Routes

	protected := group.Use(middlewares.Protected())

	// Authentication
	authGrp := protected.Group("/auth/")
	authGrp.Post("register", authGrpV1.Register)
	authGrp.Post("login", authGrpV1.Login)

	// Organizations
	orgGrp := protected.Group("/organizations/")
	orgGrp.Get("", organizationController.Query)
	orgGrp.Get(":id", organizationController.QueryByID)
	orgGrp.Post("", organizationController.Create)
	orgGrp.Patch(":id", organizationController.Update)
	orgGrp.Delete(":id", organizationController.Delete)

	// Clients
	clientGrp := protected.Group("/clients/")
	clientGrp.Get("", clientGrpV1.Query)
	clientGrp.Get(":id", clientGrpV1.QueryByID)
	clientGrp.Post("", clientGrpV1.Create)
	clientGrp.Patch(":id", clientGrpV1.Update)
	clientGrp.Delete(":id", clientGrpV1.Delete)

	// Projects
	projectGrp := protected.Group("/projects/")
	projectGrp.Get("", projectGrpV1.Query)
	projectGrp.Get(":id", projectGrpV1.QueryByID)
	projectGrp.Post("", projectGrpV1.Create)
	projectGrp.Patch(":id", projectGrpV1.Update)
	projectGrp.Delete(":id", projectGrpV1.Delete)

	//======================================================
	// Public Routes

	// ...
}
