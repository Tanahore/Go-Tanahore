package routes

import (
	"tanahore/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (routes *DeviceRouteImpl) MobileDeviceRoutes(apiGroup *echo.Group) {
	DeviceGroup := apiGroup.Group("/api/device")
	DeviceGroup.POST("/create", routes.Handler.CreateDevice, middleware.UserMiddleware)
	DeviceGroup.POST("/:id/predict", routes.Handler.PredictBestPlant)
}
