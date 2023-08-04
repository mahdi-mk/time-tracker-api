package database

import (
	"fmt"

	"github.com/mahdi-mk/time-tracker/support/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.Get("DB_HOST", ""),
		env.Get("DB_USER", ""),
		env.Get("DB_PASS", ""),
		env.Get("DB_NAME", ""),
		env.Get("DB_PORT", ""),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}
}
