package web

import "tanahore/internal/model/domain"

type CreateDeviceRequest struct {
	DeviceID string `json:"deviceID" validate:"required"`
}

type DevicePredictBestPlant struct {
	DeviceID string           `json:"deviceID" validate:"required"`
	SoilType domain.SoilTypes `json:"jenisTanah" validate:"required"`
}

type ExactPlantRequest struct {
	SoilType       string  `json:"jenisTanah"`
	LightIntensity int     `json:"intensitasCahaya"`
	PH             float32 `json:"ph"`
	Humidity       float32 `json:"kelembapan"`
	Temperature    float32 `json:"suhu"`
}
