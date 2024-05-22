package handler

import (
	"tanahore/internal/app/soil_predict/service"

	"github.com/labstack/echo/v4"
)

type SoilPredictHandler interface {
	GetPrediction(ctx echo.Context) error
	GetExactPlant(ctx echo.Context) error
}

type SoilPredictHandlerImpl struct {
	SoilPredictService service.SoilPredictService
}

func NewSoilPredictHandler(service service.SoilPredictService) SoilPredictHandler {
	return SoilPredictHandlerImpl{
		SoilPredictService: service,
	}
}
