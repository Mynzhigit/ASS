package usecase

type PaymentUsecase interface {
	ProcessPayment(orderID int, amount int) error
}

type paymentUsecase struct {
	// Здесь может быть необходимость во взаимодействии с платежной системой или других зависимостях
}

func NewPaymentUsecase() PaymentUsecase {
	return &paymentUsecase{}
}

func (uc *paymentUsecase) ProcessPayment(orderID int, amount int) error {
	// Реализация обработки платежа
	return nil
}
