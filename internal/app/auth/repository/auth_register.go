package repository

import "tanahore/internal/model/domain"

func (repo *AuthRepositoryImpl) RegisterUser(user *domain.Users) error {
	if res := repo.DB.Create(&user); res.Error != nil {
		return res.Error
	}
	return nil
}
