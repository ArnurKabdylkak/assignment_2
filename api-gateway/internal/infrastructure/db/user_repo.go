package db

import (
	"../order-service/internal/domain"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (r *UserRepo) CreateUser(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepo) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
