package service

import (
	"errors"
	"tanahore/internal/model/web"
	converter "tanahore/internal/pkg/converter/request"

	"github.com/sirupsen/logrus"
)

func (service *DeviceServiceImpl) CreateDevice(req *web.CreateDeviceRequest) error {
	if err := service.Validate.Struct(req); err != nil {
		return errors.New("validation error : " + err.Error())
	}
	if exists := service.DeviceRepository.IsExists(req.DeviceID); exists {
		return errors.New("error, device exists")
	}
	device := converter.CreateDeviceToDomain()
	if err := service.DeviceRepository.CreateDevice(req.DeviceID, device); err != nil {
		logrus.Info("error : " + err.Error())
		return err
	}
	return nil
}
