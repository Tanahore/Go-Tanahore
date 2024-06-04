package repository

import (
	"context"
	"errors"
	"tanahore/internal/model/domain"
	"time"
)

func (repository *DeviceRepositoryImpl) GetDeviceInput(id string) (*domain.Device, error) {
	for {
		var device domain.Device
		if err := repository.DB.NewRef("devices").Child(id).Get(context.Background(), &device); err != nil {
			return nil, errors.New("something happened, " + err.Error())
		}
		if device.DeviceStatus == "off" {
			return &device, nil
		}
		time.Sleep(2 * time.Second)
	}
}
