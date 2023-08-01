package database

import "github.com/mahdi-mk/time-tracker/app/models"

func RunMigrations() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Organization{},
		&models.Project{},
		&models.Client{},
	)

	if err != nil {
		panic("Failed to run auto migrations")
	}
}
