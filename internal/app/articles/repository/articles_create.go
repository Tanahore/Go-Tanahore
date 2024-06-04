package repository

import (
	"tanahore/internal/model/domain"
)

func (repository *ArticlesRepositoryImpl) CreateArticle(article *domain.Articles) error {
	if res := repository.DB.Create(&article); res.Error != nil {
		return res.Error
	}
	return nil
}
