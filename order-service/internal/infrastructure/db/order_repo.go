package db

import (
	"errors"
	"order-service/internal/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type OrderRepo struct{ DB *gorm.DB }

// Конструктор для создания репозитория
func NewOrderRepo(dsn string) *OrderRepo {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&domain.Order{}, &domain.OrderItem{})
	return &OrderRepo{DB: db}
}

// Реализация метода Create для репозитория
func (r *OrderRepo) Create(o *domain.Order) error {
	return r.DB.Create(o).Error
}

// Реализация метода GetByID для репозитория
func (r *OrderRepo) GetByID(id int) (*domain.Order, error) {
	var o domain.Order
	if err := r.DB.Preload("OrderItem").First(&o, id).Error; err != nil {
		return nil, errors.New("not found")
	}
	return &o, nil
}
