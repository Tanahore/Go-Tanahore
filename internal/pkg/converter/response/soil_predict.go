package response

import (
	"tanahore/internal/model/domain"
	"tanahore/internal/model/web"
)

func PlantsDomainToSoilPlantResponse(plants []domain.Plants) []web.SoilPlantsResponse {
	var recommendation []web.SoilPlantsResponse
	for i := range plants {
		soilPlantResponse := web.SoilPlantsResponse{
			PlantRecommendation: plants[i].PlantName,
		}
		recommendation = append(recommendation, soilPlantResponse)
	}
	return recommendation
}
