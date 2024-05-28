package web

type DeviceInputResponse struct {
	PhLevel        float32 `json:"ph"`
	Humidity       int     `json:"kelembapan"`
	Temperature    string  `json:"suhu"`
	LightIntensity int     `json:"intensitasCahaya"`
	SoilType       string  `json:"jenisTanah"`
}
