package main

import (
	"fmt"
	"net/http"

	"github.com/yourusernam/flowerstore/internal/app/usecase"
	"github.com/yourusername/flowerstore/internal/delivery/http"
	"github.com/yourusername/flowerstore/internal/repository"
	"github.com/yourusername/flowerstore/pkg/database"
	"github.com/yourusername/flowerstore/pkg/logger"
)

func main() {
	// Подключение к базе данных
	db, err := database.ConnectDB()
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Инициализация репозиториев
	userRepo := repository.NewUserRepository(db)
	adminRepo := repository.NewAdminRepository(db)

	// Инициализация usecase
	userUsecase := usecase.NewUserUsecase(userRepo)
	adminUsecase := usecase.NewAdminUsecase(adminRepo)

	// Инициализация HTTP сервера
	userHandler := http.NewUserHandler(userUsecase)
	adminHandler := http.NewAdminHandler(adminUsecase)

	http.HandleFunc("/user", userHandler.GetUserByID)
	http.HandleFunc("/admin", adminHandler.GetAdminByID)

	logger.InfoLogger.Println("Server is running on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to start server: %v", err)
	}
}
