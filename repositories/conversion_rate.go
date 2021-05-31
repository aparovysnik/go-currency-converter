package repositories

import (
	"github.com/aparovysnik/go-currency-converter/models"
	"gorm.io/gorm"
)

type conversionRate struct {
	db *gorm.DB
}

type ConversionRate interface {
	Add(conversionRate models.ConversionRate) error
}

func NewConversionRateRepository(db *gorm.DB) ConversionRate {
	return &conversionRate{
		db: db,
	}
}

// Add adds a single conversion rate to the collection
func (repo *conversionRate) Add(conversionRate models.ConversionRate) error {
	if err := repo.db.Create(&conversionRate).Error; err != nil {
		return err
	}

	return nil
}
