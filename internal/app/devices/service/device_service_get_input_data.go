package service

import (
	"errors"
	"tanahore/internal/model/domain"
	"tanahore/internal/model/web"
	converter "tanahore/internal/pkg/converter/request"

	"github.com/sirupsen/logrus"
)

func (service *DeviceServiceImpl) GetDeviceInputData(req *web.DevicePredictBestPlant) (*domain.Device, error) {
	if err := service.Validate.Struct(req); err != nil {
		return nil, errors.New("error, " + err.Error())
	}
	deviceData, err := service.DeviceRepository.GetDeviceInfo(req.DeviceID)
	if err != nil {
		return nil, errors.New("error, " + err.Error())
	}
	mapData, err := converter.DeviceInfoToMap(deviceData)
	if err != nil {
		return nil, errors.New("error, " + err.Error())
	}
	if err = service.DeviceRepository.TurnOnDevice(req.DeviceID, mapData); err != nil {
		return nil, errors.New("error, " + err.Error())
	}
	inputData, err := service.DeviceRepository.GetDeviceInput(req.DeviceID)
	if err != nil {
		return nil, errors.New("error, " + err.Error())
	}
	logrus.Info("input data : ", inputData)
	// res := converterRes.DeviceModelToDeviceInput(inputData, req.SoilType)
	return inputData, nil
}
