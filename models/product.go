package models

import "gorm.io/gorm"

// Product representa la estructura del producto en la base de datos
type Product struct {
	gorm.Model
	Name       string   `gorm:"size:100;not null" json:"name"`
	Price      float64  `gorm:"not null" json:"price"`
	Quantity   int      `gorm:"not null" json:"quantity"`
	CategoryID uint      `gorm:"not null" json:"category_id"` // Tipo de dato INT para la relación
	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`
}
