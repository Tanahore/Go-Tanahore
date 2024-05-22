package response

import (
	"tanahore/internal/model/domain"
	"tanahore/internal/model/web"
)

func PlantsDomainToSoilPlantResponse(plants []domain.Plants, soil string) web.SoilPredictPlansResponse {
	var recommendations []domain.SoilTypes
	for i := range plants {
		recommendations = append(recommendations, domain.SoilTypes(plants[i].PlantName))
	}
	res := web.SoilPredictPlansResponse{
		PredictedSoil:        soil,
		PlantRecommendations: recommendations,
	}
	return res
}
