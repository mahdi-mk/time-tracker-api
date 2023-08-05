package requests

// LoginUserRequest is what we require when a user wants to log in
type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}

// RegisterUserRequest is what we require when registering a user
type RegisterUserRequest struct {
	FirstName string `json:"first_name" validate:"required,max=150"`
	LastName  string `json:"last_name" validate:"max=150"`
	LoginUserRequest
}
