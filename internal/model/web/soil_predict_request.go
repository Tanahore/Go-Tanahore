package web

import (
	"mime/multipart"
	"tanahore/internal/model/domain"
)

type SoilPredictRequest struct {
	SoilImage *multipart.FileHeader `json:"image" `
}

type SoilPredictPlansRequest struct {
	SoilType domain.SoilTypes `json:"soilType" validate:"required"`
}
