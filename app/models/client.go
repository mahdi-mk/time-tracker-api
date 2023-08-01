package models

import "gorm.io/gorm"

// Client represents an individual client
type Client struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Projects []Project
}
