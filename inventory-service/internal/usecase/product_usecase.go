package usecase

import (
	"inventory-service/internal/domain"
	"inventory-service/internal/repo"
)

type ProductUsecase interface {
	CreateProduct(*domain.Product) error
	GetProductByID(int) (*domain.Product, error)
}

type productUC struct {
	repo repo.ProductRepository
}

func NewProductUsecase(r repo.ProductRepository) ProductUsecase {
	return &productUC{repo: r}
}

func (u *productUC) CreateProduct(p *domain.Product) error {
	return u.repo.Create(p)
}

func (u *productUC) GetProductByID(id int) (*domain.Product, error) {
	return u.repo.GetByID(id)
}
