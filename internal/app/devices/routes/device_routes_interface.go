package routes

import (
	"tanahore/internal/app/devices/handler"

	"github.com/labstack/echo/v4"
)

type DeviceRoute interface {
	MobileDeviceRoutes(apiGroup *echo.Group)
}

type DeviceRouteImpl struct {
	Handler handler.DeviceHandler
}

func NewDeviceRoutes(handler handler.DeviceHandler) DeviceRoute {
	return &DeviceRouteImpl{
		Handler: handler,
	}
}
