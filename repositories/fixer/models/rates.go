package fixer

//ConvertCurrencyResponse represents currency conversion response data
type Rates struct {
	Success bool               `json:"success"`
	Amount  float32            `json:"amount"`
	Base    string             `json:"base"`
	Rates   map[string]float32 `json:"rates"`
}
