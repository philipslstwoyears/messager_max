package service

import (
	"awesomeProject2/models"
	"awesomeProject2/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.Storage
}

func NewUserService(repo *repository.Storage) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(login, password string) (int, error) {
	if len(password) < 6 {
		return 0, errors.New("password too short")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user := models.User{
		Login:    login,
		Password: string(hashedPassword),
	}

	return s.repo.SaveUser(user)
}
