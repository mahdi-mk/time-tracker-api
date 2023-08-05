package requests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/models"
	"github.com/mahdi-mk/time-tracker/support/auth"
	"github.com/mahdi-mk/time-tracker/support/validator"
)

// CreateOrUpdateOrganization is what we require when creating or updating an organization
type CreateOrUpdateOrganizationRequest struct {
	Name string `json:"name" validate:"required,max=100"`
}

func (request *CreateOrUpdateOrganizationRequest) Validated(c *fiber.Ctx) (models.Organization, validator.ValidationErrors) {
	if err := validator.ValidateRequest(c, request); err != nil {
		return models.Organization{}, err
	}

	return models.Organization{
		Name:   request.Name,
		UserID: auth.GetID(c),
	}, nil
}
