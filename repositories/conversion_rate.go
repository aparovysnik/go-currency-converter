package repositories

import (
	"errors"

	"github.com/aparovysnik/go-currency-converter/models"
	"gorm.io/gorm"
)

type conversionRate struct {
	db *gorm.DB
}

type ConversionRate interface {
	AddOrUpdate(conversionRate models.ConversionRate) error
	GetConversionRate(fromCurrency string, toCurrency string) (float32, error)
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

//GetConversionRate finds and returns a conversion rate for given currencies
func (repo *conversionRate) GetConversionRate(fromCurrency string, toCurrency string) (float32, error) {
	var existingRate models.ConversionRate
	result := repo.db.Where("from_currency = ? AND to_currency = ?",
		fromCurrency, toCurrency).
		Find(&existingRate)

	if result.RowsAffected == 0 {
		return 0, errors.New("rate not found")
	}

	return existingRate.Rate, nil
}
