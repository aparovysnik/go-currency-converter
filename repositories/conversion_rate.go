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
func (repo *conversionRate) AddOrUpdate(newRate models.ConversionRate) error {
	var existingRate models.ConversionRate
	result := repo.db.Where("from_currency = ? AND to_currency = ?",
		newRate.FromCurrency, newRate.ToCurrency).
		Find(&existingRate)

	if result.RowsAffected > 0 {
		existingRate.Rate = newRate.Rate
		if err := repo.db.Save(&existingRate).Error; err != nil {
			return err
		}
	} else if err := repo.db.Create(&newRate).Error; err != nil {
		return err
	}

	return nil
}
