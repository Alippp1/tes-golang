package repository

import (
	"strings"

	"github.com/Alippp1/tes-golang/internal/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(item *models.Item) error
	FindAll(name string) ([]models.Item, error)
	FindByID(id uint) (*models.Item, error)

	Update(item *models.Item) error
	Updatetx(tx *gorm.DB,item *models.Item) error
	Delete(id uint) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db}
}

func (r *itemRepository) Create(item *models.Item) error {
	return r.db.Create(item).Error
}

func (r *itemRepository) FindAll(name string) ([]models.Item, error) {
	var items []models.Item
	query := r.db

	if name != "" {
		words := strings.Fields(name)
		for _, w := range words {
			query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(w)+"%")
		}
	}

	err := query.Find(&items).Error
	return items, err
}

func (r *itemRepository) FindByID(id uint) (*models.Item, error) {
	var item models.Item
	err := r.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) Update(item *models.Item) error {
	return r.db.Save(item).Error
}

func (r *itemRepository) Updatetx( tx *gorm.DB, item *models.Item) error {
	return tx.Save(item).Error
}

func (r *itemRepository) Delete(id uint) error {
	return r.db.Delete(&models.Item{}, id).Error
}
