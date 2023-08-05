package requests

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/models"
	"github.com/mahdi-mk/time-tracker/support/validator"
)

// CreateOrUpdateClient is what needed to create or update a client
type CreateOrUpdateClientRequest struct {
	Name string `json:"name" validate:"required,max=100"`
}

func (request *CreateOrUpdateClientRequest) Validated(c *fiber.Ctx) (models.Client, validator.ValidationErrors) {
	if err := validator.ValidateRequest(c, request); err != nil {
		return models.Client{}, err
	}

	orgID, _ := strconv.ParseUint(c.Locals("OrgID").(string), 10, 32)

	return models.Client{
		Name:           request.Name,
		OrganizationID: uint(orgID),
	}, nil
}
