package main

import (
	"log"
	"mathalama/config"
	"mathalama/internal/handler"
	"mathalama/internal/repository"
	"mathalama/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Загрузка конфигурации
	cfg := config.Load()

	// 2. Инициализация пула соединений БД
	dbPool, err := config.InitDB(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbPool.Close()

	// 3. Инициализация слоев (Dependency Injection)

	// User flow
	userRepo := repository.NewUserRepository(dbPool)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	// System flow
	systemRepo := repository.NewSystemRepository(dbPool)
	systemSvc := service.NewSystemService(systemRepo)
	appHandler := handler.NewAppHandler(systemSvc)

	// 4. Настройка роутера
	router := gin.Default()

	// Health check через Handler и Service
	router.GET("/health", appHandler.HealthCheck)

	// Маршруты для пользователей
	router.POST("/users", userHandler.CreateUser)

	log.Printf("Server starting on port %s", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
