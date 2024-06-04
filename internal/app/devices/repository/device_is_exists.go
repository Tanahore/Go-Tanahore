package repository

import (
	"context"
	"tanahore/internal/model/domain"

	"github.com/sirupsen/logrus"
)

func (repository *DeviceRepositoryImpl) IsExists(id string) bool {
	var device domain.Device
	if err := repository.DB.NewRef("devices").Child(id).Get(context.Background(), &device); err != nil {
		logrus.Info("error : " + err.Error())
		return false
	}
	if device.DeviceStatus == "" {
		return false
	}
	return true
}
