package handler

import (
	"tanahore/configs"
	"tanahore/internal/app/articles/service"

	"github.com/labstack/echo/v4"
)

type ArticleHandler interface {
	CreateArticle(ctx echo.Context) error
	GetAllArticles(ctx echo.Context) error
	GetArticleByID(ctx echo.Context) error
	ArticleSearch(ctx echo.Context) error
	InformationRetrieval(ctx echo.Context) error
}

type ArticleHandlerImpl struct {
	Service  service.ArticleService
	ModelURL *configs.ModelAPIConfig
}

func NewArticleHandler(service service.ArticleService, modelConfig *configs.ModelAPIConfig) ArticleHandler {
	return &ArticleHandlerImpl{
		Service:  service,
		ModelURL: modelConfig,
	}
}
