package services

type CurrencyConverter interface {
	Convert(amount float32, fromCurrency string, toCurrency string) float32
}

type currencyConverter struct {
}

func NewCurrencyConverter() CurrencyConverter {
	return &currencyConverter{}
}

func (srv *currencyConverter) Convert(amount float32, fromCurrency string, toCurrency string) float32 {
	return amount
}
