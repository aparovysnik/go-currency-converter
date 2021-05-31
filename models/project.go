package models

type Project struct {
	Base
	ContactEmail string `gorm:"unique;not null;type:varchar(100);default:null"`
	ApiKeyHash   []byte `gorm:"not null;default:null"`
}
