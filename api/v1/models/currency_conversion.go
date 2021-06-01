package models

import (
	"github.com/aparovysnik/go-currency-converter/config"
	"github.com/aparovysnik/go-currency-converter/utils"
)

//ConvertCurrencyRequest represents currency conversion request data
type ConvertCurrencyRequest struct {
	Amount       float32 `form:"amount"`
	FromCurrency string  `form:"fromCurrency"`
	ToCurrency   string  `form:"toCurrency"`
}

//ConvertCurrencyResponse represents currency conversion response data
type ConvertCurrencyResponse struct {
	BaseResponse
	Amount float32 `json:"amount"`
}

// IsValid validates request model
func (req *ConvertCurrencyRequest) IsValid(config config.Config) bool {
	return req.Amount >= 0 &&
		req.FromCurrency == "EUR" &&
		utils.ContainsString(config.SupportedCurrencies, req.ToCurrency)
}
