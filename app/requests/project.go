package requests

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/models"
	"github.com/mahdi-mk/time-tracker/support/validator"
)

// CreateOrUpdateProject is what needed to create or update a project
type CreateOrUpdateProjectRequest struct {
	Name        string `json:"name" validate:"required,max=150"`
	Description string `json:"description" validate:"max=255"`
	ClientID    uint   `json:"client_id" validate:"required,exists=clients"`
}

func (request *CreateOrUpdateProjectRequest) Validated(c *fiber.Ctx) (models.Project, validator.ValidationErrors) {
	if err := validator.ValidateRequest(c, request); err != nil {
		return models.Project{}, err
	}

	orgID, _ := strconv.ParseUint(c.Locals("OrgID").(string), 10, 32)

	return models.Project{
		Name:           request.Name,
		Description:    request.Description,
		OrganizationID: uint(orgID),
		ClientID:       request.ClientID,
	}, nil
}
