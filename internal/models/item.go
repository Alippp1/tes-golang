package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(100);not null"`
	Stock int     `gorm:"not null;default:0"`
	Price float64 `gorm:"not null"`
}

func (Item) TableName() string { return "Item" }