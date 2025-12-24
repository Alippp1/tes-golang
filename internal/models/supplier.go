package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null"`
	Email   string `gorm:"type:varchar(100)"`
	Address string `gorm:"type:text"`
}

func (Supplier) TableName() string { return "Supplier" }