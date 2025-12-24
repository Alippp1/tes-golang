package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Alippp1/tes-golang/internal/dto"
	"github.com/Alippp1/tes-golang/internal/models"
	"github.com/Alippp1/tes-golang/internal/repository"
)

type SupplierService interface {
	Create(req dto.CreateSupplierRequest) (*models.Supplier, error)
	FindAll(name string) ([]models.Supplier, error)
	FindByID(id uint) (*models.Supplier, error)
	Update(id uint, req dto.UpdateSupplierRequest) error
	Delete(id uint) error
}

type supplierService struct {
	repo repository.SupplierRepository
}

func NewSupplierService(repo repository.SupplierRepository) SupplierService {
	return &supplierService{repo}
}

func (s *supplierService) Create(req dto.CreateSupplierRequest) (*models.Supplier, error) {
	supplier := models.Supplier{
		Name:    req.Name,
		Email:  req.Email,
		Address: req.Address,
	}

	if err := s.repo.Create(&supplier); err != nil {
		return nil, err
	}

	return &supplier, nil
}

func (s *supplierService) FindAll(name string) ([]models.Supplier, error) {
	suppliers, err := s.repo.FindAll(name)
	if err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (s *supplierService) FindByID(id uint) (*models.Supplier, error) {
	supplier, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("supplier not found")
		}
		return nil, err
	}

	return supplier, nil
}

func (s *supplierService) Update(id uint, req dto.UpdateSupplierRequest) error {
	supplier, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("supplier not found")
		}
		return err
	}

	if req.Name != "" {
		supplier.Name = req.Name
	}

	if req.Email != "" {
		supplier.Email = req.Email
	}

	if req.Address != "" {
		supplier.Address = req.Address
	}

	return s.repo.Update(supplier)
}

func (s *supplierService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("supplier not found")
		}
		return err
	}

	return s.repo.Delete(id)
}
