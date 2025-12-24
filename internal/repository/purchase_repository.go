package repository

import (
	"github.com/Alippp1/tes-golang/internal/models"
	"gorm.io/gorm"
)

type PurchasingRepository interface {
	Create(tx *gorm.DB, purchasing *models.Purchasing) error
	FindAll() ([]models.Purchasing, error)
	FindByID(id uint) (*models.Purchasing, error)
}

type purchasingRepository struct {
	db *gorm.DB
}

func NewPurchasingRepository(db *gorm.DB) PurchasingRepository {
	return &purchasingRepository{db}
}

func (r *purchasingRepository) Create(tx *gorm.DB, purchasing *models.Purchasing) error {
	return tx.Create(purchasing).Error
}

func (r *purchasingRepository) FindAll() ([]models.Purchasing, error) {
	var purchases []models.Purchasing

	err := r.db.
		Preload("Supplier").
		Preload("User").
		Preload("Details").
		Preload("Details.Item").
		Order("created_at DESC").
		Find(&purchases).Error

	return purchases, err
}

func (r *purchasingRepository) FindByID(id uint) (*models.Purchasing, error) {
	var purchase models.Purchasing

	err := r.db.
		Preload("Supplier").
		Preload("User").
		Preload("Details").
		Preload("Details.Item").
		First(&purchase, id).Error

	if err != nil {
		return nil, err
	}

	return &purchase, nil
}
