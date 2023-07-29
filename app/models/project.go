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
	Name        string
	Description string
	ClientID    uint `json:"client_id"`
}
