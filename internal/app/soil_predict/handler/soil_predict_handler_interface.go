package handler

import "tanahore/internal/app/soil_predict/service"

type SoilPredictHandler interface {
}

type SoilPredictImpl struct {
	SoilPredictService service.SoilPredictService
}

func NewSoilPredictHandler(service service.SoilPredictService) SoilPredictHandler {
	return SoilPredictImpl{
		SoilPredictService: service,
	}
}
