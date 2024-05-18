package routes

import (
	"tanahore/internal/app/articles/handler"

	"github.com/labstack/echo/v4"
)

type ArticelRoutes interface {
	MobileArticleRoutes(apiGroup echo.Group)
}

type ArticleRoutesImpl struct {
	Handler handler.ArticleHandler
}

func NewArticleRoute(handler handler.ArticleHandler) ArticelRoutes {
	return &ArticleRoutesImpl{
		Handler: handler,
	}
}
