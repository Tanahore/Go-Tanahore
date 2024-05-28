package web

type CreateDeviceRequest struct {
	DeviceID string `json:"deviceID" validate:"required"`
}
