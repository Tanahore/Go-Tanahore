package repository

import (
	"context"
	"errors"
)

func (repository *DeviceRepositoryImpl) TurnOnDevice(id string, device map[string]any) error {
	if err := repository.DB.NewRef("devices").Child(id).Update(context.Background(), device); err != nil {
		return errors.New("error, " + err.Error())
	}
	return nil
}
