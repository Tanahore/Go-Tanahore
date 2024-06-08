package repository

import (
	"tanahore/internal/model/domain"

	"gorm.io/gorm"
)

type AuthRepository interface {
	RegisterUser(user *domain.Users) error
	UserExists(username string, email string) *domain.Users
}

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		DB: db,
	}
}
