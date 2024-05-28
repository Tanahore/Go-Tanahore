package service

import (
	"tanahore/internal/app/devices/repository"
	"tanahore/internal/model/web"

	"github.com/go-playground/validator"
)

type DeviceService interface {
	CreateDevice(req *web.CreateDeviceRequest) error
}

type DeviceServiceImpl struct {
	DeviceRepository repository.DeviceRepository
	Validate         *validator.Validate
}

func NewDeviceService(repo repository.DeviceRepository, validate *validator.Validate) DeviceService {
	return &DeviceServiceImpl{
		DeviceRepository: repo,
		Validate:         validate,
	}
}
