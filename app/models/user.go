package models

import "gorm.io/gorm"

// User represents an individual user
type User struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"uniqueIndex;not null"`
	Password  string `json:"password" gorm:"not null"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}

type RegisterUserRequest struct {
	FirstName string `json:"first_name" validate:"required,max=150"`
	LastName  string `json:"last_name" validate:"max=150"`
	LoginUserRequest
}
