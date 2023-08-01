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
