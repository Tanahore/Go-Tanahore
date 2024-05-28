package soil_predict

import (
	"tanahore/configs"
	HandlerPkg "tanahore/internal/app/soil_predict/handler"
	RepositoryPkg "tanahore/internal/app/soil_predict/repository"
	RoutePkg "tanahore/internal/app/soil_predict/route"
	ServicePkg "tanahore/internal/app/soil_predict/service"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func SoilPredictSetup(db *gorm.DB, validate *validator.Validate, url *configs.ModelAPIConfig) RoutePkg.SoilPredictRoutes {
	Repository := RepositoryPkg.NewSoilPredictRepository(db)
	Service := ServicePkg.NewSoilPredictService(Repository, validate)
	Handler := HandlerPkg.NewSoilPredictHandler(Service, url)
	Route := RoutePkg.NewSoilPredictRoutes(Handler)

	return Route
}
