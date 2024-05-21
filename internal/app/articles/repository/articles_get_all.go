package repository

import "tanahore/internal/model/domain"

func (repository *ArticlesRepositoryImpl) GetAllArticles() ([]domain.Articles, error) {
	var articles []domain.Articles
	if err := repository.DB.Find(&articles).Order("created_at DESC").Error; err != nil {
		return nil, err
	}
	return articles, nil
}
