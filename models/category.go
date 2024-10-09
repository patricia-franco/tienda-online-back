package models

type Category struct {
	ID uint `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"size:100;not null" json:"name"`
}
