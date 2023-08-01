package requests

// CreateOrUpdateOrganization is what we require when creating or updating an organization
type CreateOrUpdateOrganization struct {
	Name string `json:"name" validate:"required,max=100"`
}
