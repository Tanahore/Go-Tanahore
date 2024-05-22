package web

import "tanahore/internal/model/domain"

type SoilPlantRecommendationRequest struct {
	Kelembapan       int              `json:"kelembapan" validate:"required"`
	IntensitasCahaya int              `json:"intensitasCahaya" validate:"required"`
	Ph               float32          `json:"ph" validate:"required"`
	JenisTanah       domain.SoilTypes `json:"jenisTanah" validate:"required"`
	Vulkanik         domain.Volcanics `json:"vulkanik" validate:"required"`
}
