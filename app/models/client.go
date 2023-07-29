package models

import "gorm.io/gorm"

// Client represents an individual client
type Client struct {
	gorm.Model
	Name     string
	Projects []Project
}

// CreateOrUpdateClient is what needed to create or update a client
type CreateOrUpdateClient struct {
	Name string `json:"name" validate:"required,max=100"`
}
