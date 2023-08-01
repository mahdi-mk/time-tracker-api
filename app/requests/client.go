package requests

// CreateOrUpdateClient is what needed to create or update a client
type CreateOrUpdateClient struct {
	Name string `json:"name" validate:"required,max=100"`
}
