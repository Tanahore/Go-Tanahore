package repository

import (
	"tanahore/internal/model/domain"
)

func (repository *ArticlesRepositoryImpl) CheckExistingArticle(title string) *domain.Articles {
	var article domain.Articles
	if err := repository.DB.Where("title = ?", title).First(&article).Error; err != nil {
		return nil
	}
	return &article
}
