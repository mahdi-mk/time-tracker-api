package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string
	Description string
	ClientID    uint
	Client      Client
}
