package models

import (
	"time"

	"gorm.io/gorm"
)

// Entry represents an individual time entry
type Entry struct {
	gorm.Model
	Description    string
	From           time.Time `gorm:"not null"`
	To             time.Time `gorm:"not null"`
	OrganizationID uint      `gorm:"not null"`
	ProjectID      uint
}
