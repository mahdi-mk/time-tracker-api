package database

import "github.com/mahdi-mk/time-tracker/business/models"

func RunMigrations() {
	err := DB.AutoMigrate(
		&models.Project{},
		&models.Client{},
	)

	if err != nil {
		panic("Failed to run auto migrations")
	}
}
