package repository

import (
	"tanahore/internal/model/domain"
	"tanahore/internal/model/web"
)

func (soilRepository *SoilPredictRepositoryImpl) GetSoilPlantsBySoilType(soilType *web.SoilPredictPlansRequest) ([]domain.Plants, error) {
	var plants []domain.Plants
	result := soilRepository.DB.Where("plants.soiltype=?", soilType).Find(&plants)
	if result.Error != nil {
		return nil, result.Error
	}
	return plants, nil
}
