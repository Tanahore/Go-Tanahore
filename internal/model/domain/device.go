package domain

type Device struct {
	Input        DeviceInput `json:"input"`
	DeviceStatus string      `json:"deviceStatus"`
}

type DeviceInput struct {
	PhLevel        float32 `json:"ph"`
	Humidity       float32 `json:"kelembapan"`
	Temperature    float32 `json:"suhu"`
	LightIntensity int     `json:"intensitasCahaya"`
}

type Devices struct {
	Devices map[string]Device `json:"devices"`
}
