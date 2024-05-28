package repository

import (
	"database/sql"

	"github.com/yourusername/flowerstore/internal/app/entity"
)

type AdminRepository interface {
	GetByID(adminID int) (*entity.Admin, error)
}

type adminRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) AdminRepository {
	return &adminRepository{
		db: db,
	}
}

func (r *adminRepository) GetByID(adminID int) (*entity.Admin, error) {
	// Реализация метода для получения администратора по ID из базы данных
	return nil, nil
}
