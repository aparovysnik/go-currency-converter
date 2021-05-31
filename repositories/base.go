package repositories

import (
	"github.com/aparovysnik/go-currency-converter/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupDB initializes the database instance
func SetupDB() (*gorm.DB, error) {
	dsn := "host=postgres user=user password=password dbname=converter port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.Project{})

	return db, err
}
