package models

import (
	"github.com/aparovysnik/go-currency-converter/config"
	"github.com/aparovysnik/go-currency-converter/utils"
)

//ConvertCurrencyRequest represents currency conversion request data
type ConvertCurrencyRequest struct {
	Amount       float32
	FromCurrency string
	ToCurrency   string
}

//ConvertCurrencyResponse represents currency conversion response data
type ConvertCurrencyResponse struct {
	BaseResponse
	Amount float32
}

// IsValid validates request model
func (req *ConvertCurrencyRequest) IsValid(config config.Config) bool {
	return req.Amount >= 0 &&
		utils.ContainsString(config.SupportedCurrencies, req.FromCurrency) &&
		utils.ContainsString(config.SupportedCurrencies, req.ToCurrency)
}
