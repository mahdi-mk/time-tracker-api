package models

import "gorm.io/gorm"

// Organization represents an individual organization
type Organization struct {
	gorm.Model
	Name   string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
}
