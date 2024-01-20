package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Price       int    `gorm:"not null" json:"price"`
	Category    string `gorm:"not null" json:"category"`
	Description string `gorm:"not null" json:"description"`
	Image       string `gorm:"not null" json:"image"`
}
