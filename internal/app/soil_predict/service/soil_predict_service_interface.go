package service

import (
	"tanahore/internal/app/soil_predict/repository"
	"tanahore/internal/model/web"

	"github.com/go-playground/validator"
)

type SoilPredictService interface {
	GetSoilPlantsBySoilType(soilType *web.SoilPredictPlansRequest) ([]web.SoilPlantsResponse, error)
}

type SoilPredictServiceImpl struct {
	SoilPredictRepository repository.SoilPredictRepository
	Validator             *validator.Validate
}

func NewSoilPredictService(soilRepository repository.SoilPredictRepository, validator *validator.Validate) SoilPredictService {
	return &SoilPredictServiceImpl{
		SoilPredictRepository: soilRepository,
		Validator:             validator,
	}
}
