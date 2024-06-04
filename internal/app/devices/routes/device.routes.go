package routes

import "github.com/labstack/echo/v4"

func (routes *DeviceRouteImpl) MobileDeviceRoutes(apiGroup *echo.Group) {
	DeviceGroup := apiGroup.Group("/api/device")
	DeviceGroup.POST("/create", routes.Handler.CreateDevice)
	DeviceGroup.POST("/:id/predict", routes.Handler.PredictBestPlant)
}
