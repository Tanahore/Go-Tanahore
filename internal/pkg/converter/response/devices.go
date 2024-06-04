package response

import (
	"tanahore/internal/model/domain"
	"tanahore/internal/model/web"
)

func DeviceModelToDeviceInput(device *domain.Device, soilType domain.SoilTypes) *web.DeviceInputResponse {
	return &web.DeviceInputResponse{
		PhLevel:        device.Input.PhLevel,
		Humidity:       device.Input.Humidity,
		LightIntensity: device.Input.LightIntensity,
		SoilType:       string(soilType),
	}
}
