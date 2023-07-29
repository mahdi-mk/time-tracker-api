package models

import "gorm.io/gorm"

// Project contains information about an individual project
type Project struct {
	gorm.Model
	Name        string
	Description string
	ClientID    uint
	Client      Client
}

// CreateOrUpdateProject is what needed to create or update a project
type CreateOrUpdateProject struct {
	Name        string `json:"name" validate:"required,max=150"`
	Description string `json:"description" validate:"max=255"`
	ClientID    uint   `json:"client_id" validate:"required"`
}
