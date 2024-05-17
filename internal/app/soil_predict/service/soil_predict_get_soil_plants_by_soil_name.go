package service

import (
	"tanahore/internal/model/web"
	converter "tanahore/internal/pkg/converter/response"
)

func (soilPredictService *SoilPredictServiceImpl) GetSoilPlantsBySoilType(soilType *web.SoilPredictPlansRequest) ([]web.SoilPlantsResponse, error) {
	validate := soilPredictService.Validator
	if err := validate.Struct(&soilType); err != nil {
		return nil, err
	}
	plants, err := soilPredictService.SoilPredictRepository.GetSoilPlantsBySoilType(soilType)
	if err != nil {
		return nil, err
	}
	res := converter.PlantsDomainToSoilPlantResponse(plants)
	return res, nil
}
