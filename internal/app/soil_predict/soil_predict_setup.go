package soil_predict

import (
	HandlerPkg "tanahore/internal/app/soil_predict/handler"
	RepositoryPkg "tanahore/internal/app/soil_predict/repository"
	RoutePkg "tanahore/internal/app/soil_predict/route"
	ServicePkg "tanahore/internal/app/soil_predict/service"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func SoilPredictSetup(db *gorm.DB, validate *validator.Validate) RoutePkg.SoilPredictRoutes {
	Repository := RepositoryPkg.NewSoilPredictRepository(db)
	Service := ServicePkg.NewSoilPredictService(Repository, validate)
	Handler := HandlerPkg.NewSoilPredictHandler(Service)
	Route := RoutePkg.NewSoilPredictRoutes(Handler)

	return Route
}
