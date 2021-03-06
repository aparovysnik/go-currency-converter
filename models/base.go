package models

import "time"

type Base struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"modified_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
