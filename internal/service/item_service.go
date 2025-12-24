package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Alippp1/tes-golang/internal/dto"
	"github.com/Alippp1/tes-golang/internal/models"
	"github.com/Alippp1/tes-golang/internal/repository"
)

type ItemService interface {
	Create(req dto.CreateItemRequest) (*models.Item, error)
	FindAll(name string) ([]models.Item, error)
	FindByID(id uint) (*models.Item, error)
	Update(id uint, req dto.UpdateItemRequest) error
	Delete(id uint) error
}

type itemService struct {
	repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) ItemService {
	return &itemService{repo}
}

func (s *itemService) Create(req dto.CreateItemRequest) (*models.Item, error) {
	item := models.Item{
		Name:  req.Name,
		Stock: req.Stock,
		Price: req.Price,
	}

	if err := s.repo.Create(&item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemService) FindAll(name string) ([]models.Item, error) {
	items, err := s.repo.FindAll(name)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *itemService) FindByID(id uint) (*models.Item, error) {
	item, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("item not found")
		}
		return nil, err
	}

	return item, nil
}

func (s *itemService) Update(id uint, req dto.UpdateItemRequest) error {
	item, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("item not found")
		}
		return err
	}

	if req.Name != "" {
		item.Name = req.Name
	}

	if req.Stock != 0 {
		item.Stock = req.Stock
	}

	if req.Price != 0 {
		item.Price = req.Price
	}

	return s.repo.Update(item)
}

func (s *itemService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("item not found")
		}
		return err
	}

	return s.repo.Delete(id)
}
