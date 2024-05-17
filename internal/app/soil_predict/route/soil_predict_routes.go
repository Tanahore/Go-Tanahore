package route

import "tanahore/internal/app/soil_predict/handler"

type SoilPredictRoutes interface {
}

type SoilPredictRoutesImpl struct {
	Handler handler.SoilPredictHandler
}

func NewSoilPredictRoutes(handler handler.SoilPredictHandler) SoilPredictRoutes {
	return &SoilPredictRoutesImpl{
		Handler: handler,
	}
}
