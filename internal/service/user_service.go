package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/Alippp1/tes-golang/internal/dto"
	"github.com/Alippp1/tes-golang/internal/models"
	"github.com/Alippp1/tes-golang/internal/repository"
)

type UserService interface {
	FindAll(username string) ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	Update(id uint, req dto.UpdateUserRequest) error
	Delete(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) FindAll(username string) ([]models.User, error) {
	users, err := s.repo.FindAll(username)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) FindByID(id uint) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (s *userService) Update(id uint, req dto.UpdateUserRequest) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	if req.Username != "" {
		existingUser, err := s.repo.FindByUsername(req.Username)
		if err == nil && existingUser.ID != user.ID {
			return errors.New("username already exists")
		}
		user.Username = req.Username
	}

	if req.Role != "" {
		if req.Role != "admin" && req.Role != "user" {
			return errors.New("invalid role: must be 'admin' or 'user'")
		}
		user.Role = req.Role
	}

	if req.Password != "" {
		if len(req.Password) < 6 {
			return errors.New("password must be at least 6 characters")
		}

		hashed, err := bcrypt.GenerateFromPassword(
			[]byte(req.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			return errors.New("failed to hash password")
		}
		user.Password = string(hashed)
	}

	return s.repo.Update(user)
}

func (s *userService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	return s.repo.Delete(id)
}