package service

import (
	"context"
	"errors"
	"mathalama/internal/models"
	"mathalama/internal/repository"
)

// UserService описывает бизнес-логику для работы с пользователями
type UserService interface {
	RegisterUser(ctx context.Context, username, email string) error
	GetUserProfile(ctx context.Context, username string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(ctx context.Context, username, email string) error {
	// Здесь живет бизнес-логика. Например:
	// 1. Проверить, не занято ли имя (можно вызвать repo.GetUserByUsername)
	// 2. Дополнительная валидация
	// 3. Хеширование пароля (если бы он был)

	if username == "admin" {
		return errors.New("имя 'admin' зарезервировано")
	}

	return s.repo.CreateUser(ctx, username, email)
}

func (s *userService) GetUserProfile(ctx context.Context, username string) (*models.User, error) {
	return s.repo.GetUserByUsername(ctx, username)
}
