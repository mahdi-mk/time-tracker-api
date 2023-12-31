package database

import "github.com/mahdi-mk/time-tracker/app/models"

func RunMigrations() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Token{},
		&models.Organization{},
		&models.Client{},
		&models.Project{},
		&models.Entry{},
	)

	if err != nil {
		panic("Failed to run auto migrations")
	}
}
