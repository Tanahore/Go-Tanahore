package web

import "tanahore/internal/model/domain"

type CreateDeviceRequest struct {
	DeviceID string `json:"deviceID" validate:"required"`
}

type DevicePredictBestPlant struct {
	DeviceID string           `json:"deviceID" validate:"required"`
	SoilType domain.SoilTypes `json:"jenisTanah" validate:"required"`
}
