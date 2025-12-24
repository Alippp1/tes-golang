package repository

import (
	"github.com/Alippp1/tes-golang/internal/models"
	"gorm.io/gorm"
)

type PurchasingDetailRepository interface {
	BulkCreate(tx *gorm.DB, details []models.PurchasingDetail) error
}

type purchasingDetailRepository struct {
	db *gorm.DB
}

func NewPurchasingDetailRepository(db *gorm.DB) PurchasingDetailRepository {
	return &purchasingDetailRepository{db}
}

func (r *purchasingDetailRepository) BulkCreate(tx *gorm.DB, details []models.PurchasingDetail) error {
	return tx.Create(&details).Error
}