package usecase

import "github.com/yourusername/flowerstore/internal/entity"

type CatalogUsecase interface {
	GetFlowers() ([]entity.Flower, error)
}

type catalogUsecase struct {
	FlowerRepository repository.FlowerRepository
}

func NewCatalogUsecase(flowerRepository repository.FlowerRepository) CatalogUsecase {
	return &catalogUsecase{
		FlowerRepository: flowerRepository,
	}
}

func (uc *catalogUsecase) GetFlowers() ([]entity.Flower, error) {
	// Получить список цветов из репозитория
	flowers, err := uc.FlowerRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return flowers, nil
}
