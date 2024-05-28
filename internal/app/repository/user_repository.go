package repository

import (
	"database/sql"

	"github.com/yourusername/flowerstore/internal/app/entity"
)

type UserRepository interface {
	GetByID(userID int) (*entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetByID(userID int) (*entity.User, error) {
	// Реализация метода для получения пользователя по ID из базы данных
	return nil, nil
}
