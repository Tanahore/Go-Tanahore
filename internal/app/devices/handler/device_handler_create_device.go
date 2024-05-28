package handler

import (
	"strings"
	"tanahore/internal/model/web"
	"tanahore/internal/pkg/responses"

	"github.com/labstack/echo/v4"
)

func (handler *DeviceHandlerImpl) CreateDevice(ctx echo.Context) error {
	var device web.CreateDeviceRequest
	if err := ctx.Bind(&device); err != nil {
		return responses.StatusBadRequest(ctx, "error binding request", err)
	}
	if err := handler.Service.CreateDevice(&device); err != nil {
		if strings.Contains(err.Error(), "validation") {
			return responses.StatusBadRequest(ctx, "error validation struct", err)
		}
		if strings.Contains(err.Error(), "exists") {
			return responses.StatusAlreadyExist(ctx, "error", err)
		}
		return responses.StatusInternalServerError(ctx, "something went wrong", err)
	}
	return responses.StatusCreated(ctx, "successfully created device", nil, nil)
}
