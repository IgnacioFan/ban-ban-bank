package repository

import (
	"go-bank-express/internal/entity"

	"gorm.io/gorm"
)

// type UserRepository interface {
// 	CreateUser(user *entity.User) error
// }

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repo *UserRepository) CreateUser(user *entity.User) error {
	return repo.DB.Create(user).Error
}
