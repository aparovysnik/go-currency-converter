package repositories

import (
	"github.com/aparovysnik/go-currency-converter/models"
	"gorm.io/gorm"
)

type conversionRate struct {
	db *gorm.DB
}

type ConversionRate interface {
	AddOrUpdate(conversionRate models.ConversionRate) error
}

func NewConversionRateRepository(db *gorm.DB) ConversionRate {
	return &conversionRate{
		db: db,
	}
}

// Add adds a single conversion rate to the collection
func (repo *conversionRate) AddOrUpdate(conversionRate models.ConversionRate) error {
	var existingRate models.ConversionRate
	result := repo.db.Find(&existingRate, "from_currency = ? AND to_currency = ?",
		conversionRate.FromCurrency, conversionRate.ToCurrency)

	if result.RowsAffected > 0 {
		if err := repo.db.Save(&conversionRate).Error; err != nil {
			return err
		}
	} else if err := repo.db.Create(&conversionRate).Error; err != nil {
		return err
	}

	return nil
}
