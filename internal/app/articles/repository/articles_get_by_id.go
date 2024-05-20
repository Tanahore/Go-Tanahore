package repository

import (
	"tanahore/internal/model/domain"
	"tanahore/internal/model/web"
)

func (repository *ArticlesRepositoryImpl) GetArticlesByID(article_id *web.GetArticlesByID) (*domain.Articles, error) {
	var article domain.Articles
	if err := repository.DB.Where("article_id = ?", article_id.ArticleID).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}
