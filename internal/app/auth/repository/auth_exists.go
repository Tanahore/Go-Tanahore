package repository

import "tanahore/internal/model/domain"

func (repo *AuthRepositoryImpl) UserExists(username string, email string) bool {
	user := domain.Users{}
	if res := repo.DB.Where("username = ? OR email = ?", username, email).First(&user); res.Error != nil {
		return false
	}
	return true
}
