package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	dsn := "host=database-service.database-system.svc.cluster.local user=root password=password dbname=time_tracker_db port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}
}
