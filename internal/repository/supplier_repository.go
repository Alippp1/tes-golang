package repository

import (
	"strings"

	"github.com/Alippp1/tes-golang/internal/models"

	"gorm.io/gorm"
)

type SupplierRepository interface {
	Create(supplier *models.Supplier) error
	FindAll(name string) ([]models.Supplier, error)
	FindByID(id uint) (*models.Supplier, error)
	Update(supplier *models.Supplier) error
	Delete(id uint) error
}

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &supplierRepository{db}
}

func (r *supplierRepository) Create(supplier *models.Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *supplierRepository) FindAll(name string) ([]models.Supplier, error) {
	var suppliers []models.Supplier

	query := r.db

	if name != "" {
		keyword := "%" + strings.ToLower(name) + "%"
		query = query.Where("LOWER(name) LIKE ?", keyword)
	}

	err := query.Find(&suppliers).Error
	return suppliers, err
}

func (r *supplierRepository) FindByID(id uint) (*models.Supplier, error) {
	var supplier models.Supplier
	err := r.db.First(&supplier, id).Error
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (r *supplierRepository) Update(supplier *models.Supplier) error {
	return r.db.Save(supplier).Error
}

func (r *supplierRepository) Delete(id uint) error {
	return r.db.Delete(&models.Supplier{}, id).Error
}
