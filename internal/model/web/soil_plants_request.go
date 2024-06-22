package web

import "tanahore/internal/model/domain"

type SoilPlantRecommendationRequest struct {
	Kelembapan       int              `json:"kelembapan" validate:"required"`
	IntensitasCahaya int              `json:"intensitasCahaya" validate:"required"`
	Ph               float32          `json:"ph" validate:"required"`
	JenisTanah       domain.SoilTypes `json:"jenisTanah" validate:"required"`
}

type SoilTypeRequest struct {
	SoilType domain.SoilTypes `json:"jenisTanah" validate:"required" form:"jenisTanah"`
}
