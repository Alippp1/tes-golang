package models

import "gorm.io/gorm"

type Purchasing struct {
	gorm.Model
	Date        string  `gorm:"type:date;not null"`
	SupplierID uint    `gorm:"not null"`
	UserID     uint    `gorm:"not null"`
	GrandTotal float64 `gorm:"not null"`

	Supplier Supplier `gorm:"foreignKey:SupplierID"`
	User     User     `gorm:"foreignKey:UserID"`

	Details []PurchasingDetail `gorm:"foreignKey:PurchasingID"`
}

func (Purchasing) TableName() string { return "Purchasing" }

type PurchasingDetail struct {
	gorm.Model
	PurchasingID uint    `gorm:"not null"`
	ItemID       uint    `gorm:"not null"`
	Qty          int     `gorm:"not null"`
	SubTotal     float64 `gorm:"not null"`

	Purchasing Purchasing `gorm:"foreignKey:PurchasingID"`
	Item       Item       `gorm:"foreignKey:ItemID"`
}

func (PurchasingDetail) TableName() string { return "PurchasingDetail" }