package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aparovysnik/go-currency-converter/config"
	"github.com/aparovysnik/go-currency-converter/models"
	"github.com/aparovysnik/go-currency-converter/repositories"
	fixer "github.com/aparovysnik/go-currency-converter/repositories/fixer/models"
)

const fixerRatesUrl = "http://data.fixer.io/api/latest?access_key=%s"

func UpdateConversionRate(repository repositories.ConversionRate, config config.Config) {
	// Fixer free API only supports EUR as base currency
	log.Println("update rates")
	resp, err := http.Get(fmt.Sprintf(fixerRatesUrl, config.FixerAccessKey))
	if err != nil {
		log.Println(fmt.Sprintf("error retrieving rates for %s: %s.", "EUR", err.Error()))
		return
	}

	var rates fixer.Rates

	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		log.Println(fmt.Sprintf("error decoding rates for %s: %s.", "EUR", err.Error()))
		return
	}

	if !rates.Success {
		log.Println(fmt.Sprintf("Rates API returned an error for %s.", "EUR"))
		return
	}

	// Iterate through other supported currencies and storee conversion rates
	for _, toCurrency := range config.SupportedCurrencies {
		// Store the retrieved rate
		if rate, ok := rates.Rates[toCurrency]; ok {
			repository.AddOrUpdate(models.ConversionRate{
				FromCurrency: "EUR",
				ToCurrency:   toCurrency,
				Rate:         rate,
			})
		} else {
			log.Println(fmt.Sprintf("Rates API did not return a conversion rate from EUR to %s.", toCurrency))
			continue
		}
	}
}
