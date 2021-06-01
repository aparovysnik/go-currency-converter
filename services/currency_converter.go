package services

import "github.com/aparovysnik/go-currency-converter/repositories"

type CurrencyConverter interface {
	Convert(amount float32, fromCurrency string, toCurrency string) (float32, error)
}

type currencyConverter struct {
	repository repositories.ConversionRate
}

func NewCurrencyConverter(repository repositories.ConversionRate) CurrencyConverter {
	return &currencyConverter{
		repository: repository,
	}
}

func (srv *currencyConverter) Convert(amount float32, fromCurrency string, toCurrency string) (float32, error) {
	rate, err := srv.repository.GetConversionRate(fromCurrency, toCurrency)

	if err != nil {
		return 0, err
	}

	return amount * rate, nil
}
