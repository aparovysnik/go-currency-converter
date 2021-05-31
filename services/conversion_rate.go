package services

import (
	"log"

	"github.com/aparovysnik/go-currency-converter/models"
	"github.com/aparovysnik/go-currency-converter/repositories"
)

func UpdateConversionRate(repository repositories.ConversionRate) {
	log.Println("update rate")
	repository.AddOrUpdate(models.ConversionRate{
		FromCurrency: "EUR",
		ToCurrency:   "USD",
		Rate:         1.33,
	})
}
