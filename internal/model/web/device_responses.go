package web

type DeviceInputResponse struct {
	PhLevel        float32 `json:"ph"`
	Humidity       float32 `json:"kelembapan"`
	LightIntensity int     `json:"intensitasCahaya"`
	SoilType       string  `json:"jenisTanah"`
}
