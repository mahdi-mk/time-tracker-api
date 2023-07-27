package database

import (
	"github.com/mahdi-mk/time-tracker/business/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := "host=database-service.database-system.svc.cluster.local user=root password=password dbname=time_tracker_db port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}

	// ===========================================================
	// Run Migrations

	err = DB.AutoMigrate(
		&models.Project{},
		&models.Client{},
	)

	if err != nil {
		panic("Failed to auto migrate")
	}
}
