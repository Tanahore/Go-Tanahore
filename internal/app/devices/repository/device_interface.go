package repository

import (
	"tanahore/internal/model/domain"

	"firebase.google.com/go/db"
)

type DeviceRepository interface {
	CreateDevice(id string, device *domain.Device) error
	IsExists(id string) bool
	GetDeviceInput(id string) (*domain.Device, error)
	TurnOnDevice(id string, device map[string]any) error
	GetDeviceInfo(id string) (*domain.Device, error)
}

type DeviceRepositoryImpl struct {
	DB *db.Client
}

func NewDeviceRepository(db *db.Client) DeviceRepository {
	return &DeviceRepositoryImpl{
		DB: db,
	}
}
