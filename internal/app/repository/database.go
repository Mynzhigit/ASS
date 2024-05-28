package repository

import (
	"database/sql"
	"log"

	"github.com/yourusername/flowerstore/internal/entity"
)

type FlowerRepository struct {
	DB *sql.DB
}

func NewFlowerRepository(db *sql.DB) *FlowerRepository {
	return &FlowerRepository{
		DB: db,
	}
}

func (r *FlowerRepository) GetAll() ([]entity.Flower, error) {
	// Запрос к базе данных для получения всех цветов
	rows, err := r.DB.Query("SELECT * FROM flowers")
	if err != nil {
		log.Printf("Failed to fetch flowers: %v", err)
		return nil, err
	}
	defer rows.Close()

	var flowers []entity.Flower
	for rows.Next() {
		var flower entity.Flower
		err := rows.Scan(&flower.ID, &flower.Name, &flower.Price)
		if err != nil {
			log.Printf("Failed to scan flower: %v", err)
			continue
		}
		flowers = append(flowers, flower)
	}

	return flowers, nil
}
