package router

import (
	"github.com/gofiber/fiber/v2"
)

func Resource(name string, group fiber.Router, controller ResourceController) {
	prefix := "/" + name + "/"

	group.Get(prefix, controller.Query)
	group.Get(prefix+":id", controller.QueryByID)
	group.Post(prefix, controller.Create)
	group.Patch(prefix+":id", controller.Update)
	group.Delete(prefix+":id", controller.Delete)
}

// ResourceController containts all methods for CRUD operations.
type ResourceController interface {
	// Query returns a paginated list of a resource
	Query(c *fiber.Ctx) error

	// QueryByID returns a specific resource by its ID
	QueryByID(c *fiber.Ctx) error

	// Create stores a new resource to the database
	Create(c *fiber.Ctx) error

	// Update updates the data of the specified resource
	Update(c *fiber.Ctx) error

	// Delete deletes the specified resource from the database
	Delete(c *fiber.Ctx) error
}
