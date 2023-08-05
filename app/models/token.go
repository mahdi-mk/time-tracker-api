package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Token     string
	UserID    uint
	ExpiredAt sql.NullTime
}
