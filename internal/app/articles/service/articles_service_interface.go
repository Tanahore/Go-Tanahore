package service

import (
	"tanahore/internal/app/articles/repository"
	"tanahore/internal/model/web"

	"github.com/go-playground/validator"
)

type ArticleService interface {
	CreateArticle(req *web.CreateArticle) error
}

type ArticleServiceImpl struct {
	Repository repository.ArticlesRepository
	Validate   *validator.Validate
}

func NewArticleService(repo repository.ArticlesRepository, validate *validator.Validate) ArticleService {
	return &ArticleServiceImpl{
		Repository: repo,
		Validate:   validate,
	}
}
