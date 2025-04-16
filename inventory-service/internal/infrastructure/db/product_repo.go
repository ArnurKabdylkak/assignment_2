package repo

import (
	"errors"
	"inventory-service/internal/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProductRepo struct{ DB *gorm.DB }

func NewProductRepo(dsn string) *ProductRepo {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&domain.Product{})
	return &ProductRepo{DB: db}
}

func (r *ProductRepo) Create(p *domain.Product) error {
	return r.DB.Create(p).Error
}

func (r *ProductRepo) GetByID(id int) (*domain.Product, error) {
	var p domain.Product
	if err := r.DB.First(&p, id).Error; err != nil {
		return nil, errors.New("not found")
	}
	return &p, nil
}
