package usecase

import "github.com/yourusername/flowerstore/internal/entity"

type OrderItem struct {
	FlowerID int
	Quantity int
}

type OrderUsecase interface {
	CreateOrder(userID int, items []OrderItem) (int, error)
	GetUserOrders(userID int) ([]entity.Order, error)
	GetOrderDetails(orderID int) ([]OrderItem, error)
}

type orderUsecase struct {
	// Здесь может быть необходимость в репозитории заказов или других зависимостях
}

func NewOrderUsecase() OrderUsecase {
	return &orderUsecase{}
}

func (uc *orderUsecase) CreateOrder(userID int, items []OrderItem) (int, error) {
	// Реализация создания заказа
	return 0, nil
}

func (uc *orderUsecase) GetUserOrders(userID int) ([]entity.Order, error) {
	// Реализация получения заказов пользователя
	return nil, nil
}

func (uc *orderUsecase) GetOrderDetails(orderID int) ([]OrderItem, error) {
	// Реализация получения деталей заказа
	return nil, nil
}
