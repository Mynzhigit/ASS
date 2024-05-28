package grpc

import (
	"context"

	"github.com/yourusername/flowerstore/internal/app/usecase"
	"github.com/yourusername/flowerstore/pkg/logger"
)

type UserService struct {
	userUsecase usecase.UserUsecase
}

func NewUserService(userUsecase usecase.UserUsecase) *UserService {
	return &UserService{
		userUsecase: userUsecase,
	}
}

func (s *UserService) GetUserByID(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	logger.InfoLogger.Println("Received GetUserByID gRPC request")

	// Реализация метода для получения пользователя по ID
	return nil, nil
}
