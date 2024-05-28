package app

import (
	"tanahore/configs"
	articles "tanahore/internal/app/articles"
	soilPredict "tanahore/internal/app/soil_predict"
	"tanahore/internal/pkg/cloudinary"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitApp(db *gorm.DB, validate *validator.Validate, e *echo.Echo, cloudinary *cloudinary.CloudinaryUploader, url *configs.ModelAPIConfig) {
	apiGroupMobile := e.Group("mobile")
	soilPredictRoutes := soilPredict.SoilPredictSetup(db, validate, url)
	soilPredictRoutes.MobileSoilPredict(apiGroupMobile)

	articleRoutes := articles.ArticleSetup(db, validate, *cloudinary)
	articleRoutes.MobileArticleRoutes(*apiGroupMobile)

	apiGroupWeb := e.Group("web")
	articleRoutes.WebArticleRoutes(*apiGroupWeb)
}
