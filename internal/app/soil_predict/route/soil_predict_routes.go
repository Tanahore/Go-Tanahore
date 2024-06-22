package route

import (
	"tanahore/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (route *SoilPredictRoutesImpl) MobileSoilPredict(apiGroup *echo.Group) {
	SoilPredictGroup := apiGroup.Group("/api/predict")

	SoilPredictGroup.POST("/soil", route.Handler.GetPrediction, middleware.UserMiddleware)
	SoilPredictGroup.POST("/plant", route.Handler.GetExactPlant, middleware.UserMiddleware)
}
