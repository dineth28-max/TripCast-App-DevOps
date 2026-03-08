package services

import (
	"auth-service/models"
	"auth-service/repository"
	"auth-service/utils"
	"errors"

	"github.com/google/uuid"
)

type AuthService struct {
	Repo *repository.UserRepo
}

func NewAuthService(repo *repository.UserRepo) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) RegisterUser(email, password, name string) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := models.User{
		ID:       uuid.New().String(),
		Email:    email,
		Password: hashedPassword,
		Name:     name,
	}

	if err := s.Repo.CreateUser(user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) LoginUser(email, password string) (string, *models.User, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", nil, errors.New("failed to generate token")
	}

	return token, user, nil
}
