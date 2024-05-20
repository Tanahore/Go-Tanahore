package repository

import (
	"tanahore/internal/model/domain"
	"tanahore/internal/model/web"

	"gorm.io/gorm"
)

type ArticlesRepository interface {
	CreateArticle(article *domain.Articles) error
	CheckExistingArticle(title string) bool
	GetAllArticles() ([]domain.Articles, error)
	GetArticlesByID(id *web.GetArticlesByID) (*domain.Articles, error)
}

type ArticlesRepositoryImpl struct {
	DB gorm.DB
}

func NewArticlesRepository(db *gorm.DB) ArticlesRepository {
	return &ArticlesRepositoryImpl{
		DB: *db,
	}
}
