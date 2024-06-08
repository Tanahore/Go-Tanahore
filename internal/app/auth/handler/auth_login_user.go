package handler

import (
	"strings"
	"tanahore/internal/model/web"
	"tanahore/internal/pkg/jwt"
	"tanahore/internal/pkg/responses"

	"github.com/labstack/echo/v4"
)

func (handler *AuthHandlerImpl) LoginUser(ctx echo.Context) error {
	var req web.LoginUserRequest
	if err := ctx.Bind(&req); err != nil {
		return responses.StatusBadRequest(ctx, "login failed", err)
	}
	auth, res, err := handler.AuthService.LoginUser(&req)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return responses.StatusBadRequest(ctx, "login failed", err)
		}
		if strings.Contains(err.Error(), "not found") {
			return responses.StatusNotFound(ctx, "login failed", err)
		}
		return responses.StatusInternalServerError(ctx, "login failed", err)
	}
	res.AccessToken, err = jwt.GenerateAccessToken(auth)
	if err != nil {
		return responses.StatusInternalServerError(ctx, "login failed", err)
	}
	return responses.StatusOK(ctx, "login successfull", res, nil)
}
