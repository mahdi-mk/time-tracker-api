package router

import (
	"github.com/gofiber/fiber/v2"
	clientGrpV1 "github.com/mahdi-mk/time-tracker/app/handlers/v1/client"
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
}
