package service

import (
	"context"
	"mathalama/internal/repository"
)

type SystemService interface {
	CheckHealth(ctx context.Context) error
}

type systemService struct {
	repo repository.SystemRepository
}

func NewSystemService(repo repository.SystemRepository) SystemService {
	return &systemService{repo: repo}
}

func (s *systemService) CheckHealth(ctx context.Context) error {
	// Здесь может быть сложная логика проверки здоровья:
	// - Проверить БД
	// - Проверить Redis (если есть)
	// - Проверить свободное место на диске и т.д.
	return s.repo.Ping(ctx)
}
