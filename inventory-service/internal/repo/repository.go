package repo

import "inventory-service/internal/domain"

type ProductRepository interface {
	Create(p *domain.Product) error
	GetByID(id int) (*domain.Product, error)
}
