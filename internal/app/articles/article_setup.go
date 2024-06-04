package articles

import (
	HandlerPkg "tanahore/internal/app/articles/handler"
	RepositoryPkg "tanahore/internal/app/articles/repository"
	RoutePkg "tanahore/internal/app/articles/routes"
	ServicePkg "tanahore/internal/app/articles/service"
	"tanahore/internal/pkg/cloudinary"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func ArticleSetup(db *gorm.DB, validate *validator.Validate, cloudinary cloudinary.CloudinaryUploader) RoutePkg.ArticelRoutes {
	repository := RepositoryPkg.NewArticlesRepository(db)
	service := ServicePkg.NewArticleService(repository, validate, cloudinary)
	handler := HandlerPkg.NewArticleHandler(service)
	route := RoutePkg.NewArticleRoute(handler)

	return route
}
