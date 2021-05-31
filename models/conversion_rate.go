package models

type ConversionRate struct {
	Base
	FromCurrency string  `gorm:"not null;type:varchar(100);default:null"`
	ToCurrency   string  `gorm:"not null;type:varchar(100);default:null"`
	Rate         float32 `gorm:"not null;default:null"`
}
