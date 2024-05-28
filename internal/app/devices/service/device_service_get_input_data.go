package service

import (
	"errors"
	"tanahore/internal/model/web"
	converter "tanahore/internal/pkg/converter/request"
	converterRes "tanahore/internal/pkg/converter/response"
)

func (service *DeviceServiceImpl) GetDeviceInputData(req *web.DevicePredictBestPlant) (*web.DeviceInputResponse, error) {
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
	res := converterRes.DeviceModelToDeviceInput(inputData, req.SoilType)
	return res, nil
}
