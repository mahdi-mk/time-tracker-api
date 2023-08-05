package models

import "gorm.io/gorm"

// Project contains information about an individual project
type Project struct {
	gorm.Model
	Name           string `gorm:"not null"`
	Description    string
	OrganizationID uint `gorm:"not null"`
	ClientID       uint `gorm:"not null"`
}
