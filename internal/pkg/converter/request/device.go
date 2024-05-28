package request

import "tanahore/internal/model/domain"

func CreateDeviceToDomain() *domain.Device {
	return &domain.Device{
		DeviceStatus: "off",
	}
}
