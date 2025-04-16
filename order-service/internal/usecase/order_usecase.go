package usecase

import "order-service/internal/domain"

type OrderRepository interface {
	Create(o *domain.Order) error
	GetByID(id int) (*domain.Order, error)
}

// Интерфейс бизнес-логики
type OrderUsecase interface {
	CreateOrder(o *domain.Order) error
	GetOrderByID(id int) (*domain.Order, error)
}

type orderUC struct {
	repo OrderRepository
}

// Конструктор для создания OrderUsecase
func NewOrderUsecase(r OrderRepository) OrderUsecase {
	return &orderUC{repo: r}
}

// Реализация методов бизнес-логики
func (u *orderUC) CreateOrder(o *domain.Order) error {
	return u.repo.Create(o) // вызов метода Create репозитория
}

func (u *orderUC) GetOrderByID(id int) (*domain.Order, error) {
	return u.repo.GetByID(id) // вызов метода GetByID репозитория
}
