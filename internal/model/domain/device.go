package domain

type Device struct {
	WifiSSID     string      `json:"wifi_ssid"`
	WifiPass     string      `json:"wifi_pass"`
	Input        DeviceInput `json:"input"`
	DeviceStatus string      `json:"deviceStatus"`
}

type DeviceInput struct {
	PhLevel        float32 `json:"ph"`
	Humidity       int     `json:"kelembapan"`
	Temperature    string  `json:"suhu"`
	LightIntensity int     `json:"intensitasCahaya"`
}

type Devices struct {
	Devices map[string]Device `json:"devices"`
}
