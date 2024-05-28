package handler

import (
	"tanahore/configs"
	"tanahore/internal/app/devices/service"

	"github.com/labstack/echo/v4"
)

type DeviceHandler interface {
	CreateDevice(ctx echo.Context) error
	PredictBestPlant(ctx echo.Context) error
}

type DeviceHandlerImpl struct {
	Service  service.DeviceService
	ModelURL *configs.ModelAPIConfig
}

func NewDeviceHandler(service service.DeviceService, modelURL *configs.ModelAPIConfig) DeviceHandler {
	return &DeviceHandlerImpl{
		Service:  service,
		ModelURL: modelURL,
	}
}
