package repositories

import (
	"github.com/aparovysnik/go-currency-converter/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Base struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

// SetupDB initializes the database instance
func SetupDB() (*gorm.DB, error) {
	dsn := "host=postgres user=user password=password dbname=converter port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.Project{})

	return db, err
}
