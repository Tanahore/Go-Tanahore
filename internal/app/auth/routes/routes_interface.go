package routes

import (
	"tanahore/internal/app/auth/handler"

	"github.com/labstack/echo/v4"
)

type AuthRoutes interface {
	MobileRoutes(apiGroup *echo.Group)
}

type AuthRoutesImpl struct {
	AuthHandler handler.AuthHandler
}

func NewAuthRoutes(handler handler.AuthHandler) AuthRoutes {
	return &AuthRoutesImpl{
		AuthHandler: handler,
	}
}
