package route

import "github.com/labstack/echo/v4"

func (route *SoilPredictRoutesImpl) MobileSoilPredict(apiGroup *echo.Group) {
	SoilPredictGroup := apiGroup.Group("/api/predict")

	SoilPredictGroup.POST("", route.Handler.GetPrediction)
}
