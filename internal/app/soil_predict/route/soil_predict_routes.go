package route

import "github.com/labstack/echo/v4"

func (route *SoilPredictRoutesImpl) MobileSoilPredict(apiGroup *echo.Group) {
	SoilPredictGroup := apiGroup.Group("/api/predict")

	SoilPredictGroup.POST("/soil", route.Handler.GetPrediction)
	SoilPredictGroup.POST("/plant", route.Handler.GetExactPlant)
}
