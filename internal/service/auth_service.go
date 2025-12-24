package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/Alippp1/tes-golang/internal/models"
	"github.com/Alippp1/tes-golang/internal/repository"
	"github.com/Alippp1/tes-golang/internal/utils"
)

type AuthService interface {
	Register(username, password, role string) error
	Login(username, password string) (string, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Register(username, password, role string) error {
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	if role == "" {
		role = "user"
	}

	user := &models.User{
		Username: username,
		Password: string(hashed),
		Role:     role,
	}

	return s.userRepo.Create(user)
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	); err != nil {
		return "", errors.New("invalid username or password")
	}

	return utils.GenerateJWT(user.ID, user.Role, user.Username)
}
