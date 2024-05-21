package service

import (
	"tanahore/internal/app/articles/repository"
	"tanahore/internal/model/web"
	"tanahore/internal/pkg/cloudinary"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ArticleService interface {
	CreateArticle(req *web.CreateArticle, ctx echo.Context) error
	GetAllArticles() ([]web.GetAllArticlesResponse, error)
	GetArticlesByID(article_id *web.GetArticlesByID) (*web.AllArticlesResponse, error)
	ArticleSearch(req *web.GetArticlesByTitle) ([]web.GetAllArticlesResponse, error)
}

type ArticleServiceImpl struct {
	Repository repository.ArticlesRepository
	Validate   *validator.Validate
	Cloudinary cloudinary.CloudinaryUploader
}

func NewArticleService(repo repository.ArticlesRepository, validate *validator.Validate, cloudinary cloudinary.CloudinaryUploader) ArticleService {
	return &ArticleServiceImpl{
		Repository: repo,
		Validate:   validate,
		Cloudinary: cloudinary,
	}
}
