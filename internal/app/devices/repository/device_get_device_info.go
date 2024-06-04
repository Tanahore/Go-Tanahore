package repository

import (
	"context"
	"errors"
	"tanahore/internal/model/domain"
)

func (repostory *DeviceRepositoryImpl) GetDeviceInfo(id string) (*domain.Device, error) {
	var device domain.Device
	if err := repostory.DB.NewRef("devices").Child(id).Get(context.Background(), &device); err != nil {
		return nil, errors.New("error, " + err.Error())
	}
	return &device, nil
}
