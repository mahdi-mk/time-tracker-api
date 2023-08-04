package requests

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/models"
	"github.com/mahdi-mk/time-tracker/support/validator"
)

type CreateOrUpdateEntryRequest struct {
	Description string `json:"description" validate:"max=100"`
	From        string `json:"from" validate:"required,datetime_custom"`
	To          string `json:"to" validate:"required,datetime_custom"`
	ProjectID   uint   `json:"project_id"`
}

// GetValidatedData validates the request and returns the validated data as a model
func (request *CreateOrUpdateEntryRequest) GetValidatedData(c *fiber.Ctx) (models.Entry, validator.ValidationErrors) {
	if err := validator.ValidateRequest(c, request); err != nil {
		return models.Entry{}, err
	}

	orgID, _ := strconv.ParseUint(c.Locals("OrgID").(string), 10, 32)
	from, _ := time.Parse("2006-01-02 15:04:05.999999-07", request.From)
	to, _ := time.Parse("2006-01-02 15:04:05.999999-07", request.From)

	entry := models.Entry{
		Description:    request.Description,
		From:           from,
		To:             to,
		OrganizationID: uint(orgID),
		ProjectID:      request.ProjectID,
	}

	return entry, nil
}
