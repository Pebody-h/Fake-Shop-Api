package models

type Product struct {
	ID          int    `gorm:"primaryKey"`
	name        string `gorm:"not null"`
	price       int    `gorm:"not null"`
	category    string `gorm:"not null"`
	description string `gorm:"not null"`
	image       string `gorm:"not null"`
}
