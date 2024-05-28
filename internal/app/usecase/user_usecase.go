package usecase

import (
	"github.com/yourusername/flowerstore/internal/app/entity"
	"github.com/yourusername/flowerstore/internal/app/repository"
)

type UserUsecase interface {
	GetUserByID(userID int) (*entity.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) GetUserByID(userID int) (*entity.User, error) {
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
