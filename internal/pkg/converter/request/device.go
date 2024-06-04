package request

import (
	"encoding/json"
	"tanahore/internal/model/domain"
)

func CreateDeviceToDomain() *domain.Device {
	return &domain.Device{
		DeviceStatus: "off",
	}
}

func DeviceInfoToMap(device *domain.Device) (map[string]interface{}, error) {
	device.DeviceStatus = "on"
	dataJSON, err := json.Marshal(device)
	if err != nil {
		return nil, err
	}

	var dataMap map[string]interface{}
	if err := json.Unmarshal(dataJSON, &dataMap); err != nil {
		return nil, err
	}
	return dataMap, nil
}
