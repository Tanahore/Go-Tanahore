package handler

import (
	"tanahore/configs"
	"tanahore/internal/app/soil_predict/service"

	"github.com/labstack/echo/v4"
)

type SoilPredictHandler interface {
	GetPrediction(ctx echo.Context) error
	GetExactPlant(ctx echo.Context) error
}

type SoilPredictHandlerImpl struct {
	SoilPredictService service.SoilPredictService
	URLConfig          *configs.ModelAPIConfig
}

func NewSoilPredictHandler(service service.SoilPredictService, url *configs.ModelAPIConfig) SoilPredictHandler {
	return SoilPredictHandlerImpl{
		SoilPredictService: service,
		URLConfig:          url,
	}
}
