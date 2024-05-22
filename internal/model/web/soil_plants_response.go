package web

import "tanahore/internal/model/domain"

type SoilPlantsResponse struct {
	PlantRecommendation string `json:"plantRecommendation"`
}

type SoilPredictPlansResponse struct {
	PredictedSoil        string             `json:"predictedSoil"`
	PlantRecommendations []domain.SoilTypes `json:"plantRecommendations"`
}

type ModelResponseData struct {
	JenisTanah string `json:"jenis_tanah"`
}

type ModelResponseStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ModelApiResponse struct {
	Data   ModelResponseData   `json:"data"`
	Status ModelResponseStatus `json:"status"`
}
