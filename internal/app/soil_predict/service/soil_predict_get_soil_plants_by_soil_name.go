package service

import (
	"errors"
	"tanahore/internal/model/web"
	converter "tanahore/internal/pkg/converter/response"
)

func (soilPredictService *SoilPredictServiceImpl) GetSoilPlantsBySoilType(soilType *web.SoilPredictPlansRequest) (*web.SoilPredictPlansResponse, error) {
	validate := soilPredictService.Validator
	if err := validate.Struct(soilType); err != nil {
		return nil, err
	}
	plants, err := soilPredictService.SoilPredictRepository.GetSoilPlantsBySoilType(soilType)
	if err != nil {
		return nil, err
	}
	if len(plants) == 0 {
		return nil, errors.New("plants not found")
	}
	res := converter.PlantsDomainToSoilPlantResponse(plants, string(soilType.SoilType))
	return &res, nil
}
