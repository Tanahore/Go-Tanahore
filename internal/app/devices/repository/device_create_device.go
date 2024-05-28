package repository

import (
	"context"
	"tanahore/internal/model/domain"
)

func (repository *DeviceRepositoryImpl) CreateDevice(id string, device *domain.Device) error {
	ref := repository.DB.NewRef("devices").Child(id)
	err := ref.Set(context.Background(), &device)
	if err != nil {
		return err
	}
	return nil
}
