package web

import "tanahore/internal/model/domain"

type SoilPlantsResponse struct {
	PlantRecommendation string `json:"plantRecommendation"`
}

type SoilPredictPlansResponse struct {
	PredictedSoil        string             `json:"predictedSoil"`
	PlantRecommendations []domain.SoilTypes `json:"plantRecommendations"`
}

type ModelResponseStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// soil predict
type ImageModelResponseData struct {
	JenisTanah string `json:"jenis_tanah"`
}

type ImageModelApiResponse struct {
	Data   ImageModelResponseData `json:"data"`
	Status ModelResponseStatus    `json:"status"`
}

// plants predict
type PlantModelResponseData struct {
	Plant string `json:"plantRecommendation"`
}

type PlantModelApiResponse struct {
	Data   PlantModelResponseData `json:"data"`
	Status ModelResponseStatus    `json:"status"`
}
