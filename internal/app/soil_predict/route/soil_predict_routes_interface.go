package route

import (
	"tanahore/internal/app/soil_predict/handler"

	"github.com/labstack/echo/v4"
)

type SoilPredictRoutes interface {
	MobileSoilPredict(apiGroup *echo.Group)
}

type SoilPredictRoutesImpl struct {
	Handler handler.SoilPredictHandler
}

func NewSoilPredictRoutes(handler handler.SoilPredictHandler) SoilPredictRoutes {
	return &SoilPredictRoutesImpl{
		Handler: handler,
	}
}
