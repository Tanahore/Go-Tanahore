package repository

import (
	"tanahore/internal/model/domain"

	"gorm.io/gorm"
)

func (repository *ArticlesRepositoryImpl) CheckExistingArticle(title string) bool {
	var article domain.Articles
	if err := repository.DB.Where("title = ?", title).First(&article).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
		return false
	}
	return true
}
