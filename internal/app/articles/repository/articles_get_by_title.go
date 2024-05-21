package repository

import (
	"tanahore/internal/model/domain"
)

func (repository *ArticlesRepositoryImpl) GetArticlesByTitle(title string) ([]domain.Articles, error) {
	var articles []domain.Articles
	searchPattern := "%" + title + "%"
	if err := repository.DB.Where("title LIKE ?", searchPattern).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}
