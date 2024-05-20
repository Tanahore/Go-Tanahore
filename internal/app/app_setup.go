package app

import (
	articles "tanahore/internal/app/articles"
	soilPredict "tanahore/internal/app/soil_predict"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitApp(db *gorm.DB, validate *validator.Validate, e *echo.Echo) {
	apiGroupMobile := e.Group("mobile")
	soilPredictRoutes := soilPredict.SoilPredictSetup(db, validate)
	soilPredictRoutes.MobileSoilPredict(apiGroupMobile)

	articleRoutes := articles.ArticleSetup(db, validate)
	articleRoutes.MobileArticleRoutes(*apiGroupMobile)
}