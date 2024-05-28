package usecase

import "github.com/yourusername/flowerstore/internal/entity"

type CartItem struct {
	FlowerID int
	Quantity int
}

type CartUsecase interface {
	AddItemToCart(userID int, flowerID, quantity int) error
	GetCartItems(userID int) ([]CartItem, error)
	RemoveItemFromCart(userID int, flowerID int) error
}

type cartUsecase struct {
	// Здесь может быть необходимость в репозитории корзины или других зависимостях
}

func NewCartUsecase() CartUsecase {
	return &cartUsecase{}
}

func (uc *cartUsecase) AddItemToCart(userID int, flowerID, quantity int) error {
	// Реализация добавления товара в корзину
	return nil
}

func (uc *cartUsecase) GetCartItems(userID int) ([]CartItem, error) {
	// Реализация получения товаров из корзины
	return nil, nil
}

func (uc *cartUsecase) RemoveItemFromCart(userID int, flowerID int) error {
	// Реализация удаления товара из корзины
	return nil
}
