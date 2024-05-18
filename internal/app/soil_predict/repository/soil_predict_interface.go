package repository

import (
	"tanahore/internal/model/domain"
	"tanahore/internal/model/web"

	"gorm.io/gorm"
)

type SoilPredictRepository interface {
	GetSoilPlantsBySoilType(soilType *web.SoilPredictPlansRequest) ([]domain.Plants, error)
}

type SoilPredictRepositoryImpl struct {
	DB *gorm.DB
}

func NewSoilPredictRepository(db *gorm.DB) SoilPredictRepository {
	return &SoilPredictRepositoryImpl{
		DB: db,
	}
}
