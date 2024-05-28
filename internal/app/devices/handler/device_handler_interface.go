package handler

import (
	"tanahore/internal/app/devices/service"

	"github.com/labstack/echo/v4"
)

type DeviceHandler interface {
	CreateDevice(ctx echo.Context) error
}

type DeviceHandlerImpl struct {
	Service service.DeviceService
}

func NewDeviceHandler(service service.DeviceService) DeviceHandler {
	return &DeviceHandlerImpl{
		Service: service,
	}
}
