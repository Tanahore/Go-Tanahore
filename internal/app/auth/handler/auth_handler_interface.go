package handler

import (
	"tanahore/internal/app/auth/service"

	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	RegisterUser(ctx echo.Context) error
}

type AuthHandlerImpl struct {
	AuthService service.AuthService
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return &AuthHandlerImpl{
		AuthService: service,
	}
}
