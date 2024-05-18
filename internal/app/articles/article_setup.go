package articles

import (
	HandlerPkg "tanahore/internal/app/articles/handler"
	RepositoryPkg "tanahore/internal/app/articles/repository"
	RoutePkg "tanahore/internal/app/articles/routes"
	ServicePkg "tanahore/internal/app/articles/service"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func ArticleSetup(db *gorm.DB, validate *validator.Validate) RoutePkg.ArticelRoutes {
	repository := RepositoryPkg.NewArticlesRepository(db)
	service := ServicePkg.NewArticleService(repository, validate)
	handler := HandlerPkg.NewArticleHandler(service)
	route := RoutePkg.NewArticleRoute(handler)

	return route
}
