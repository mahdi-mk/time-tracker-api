package router

import (
	"github.com/gofiber/fiber/v2"
	clientGrpV1 "github.com/mahdi-mk/time-tracker/app/handlers/v1/client"
	projectGrpV1 "github.com/mahdi-mk/time-tracker/app/handlers/v1/project"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// ===========================================================
	// Version 1

	v1 := api.Group("/v1")

	// Client group
	clientGrp := v1.Group("/clients")
	clientGrp.Get("/", clientGrpV1.Query)
	clientGrp.Get("/:id", clientGrpV1.QueryByID)
	clientGrp.Post("/", clientGrpV1.Create)
	clientGrp.Patch("/:id", clientGrpV1.Update)
	clientGrp.Delete("/:id", clientGrpV1.Delete)

	// Project group
	projectGrp := v1.Group("/projects")
	projectGrp.Get("/", projectGrpV1.Query)
	projectGrp.Get("/:id", projectGrpV1.QueryByID)
	projectGrp.Post("/", projectGrpV1.Create)
	projectGrp.Patch("/:id", projectGrpV1.Update)
	projectGrp.Delete("/:id", projectGrpV1.Delete)
}
