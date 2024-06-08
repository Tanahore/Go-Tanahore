package handler

import (
	"strings"

	"tanahore/internal/model/web"
	conversion "tanahore/internal/pkg/converter/response"
	"tanahore/internal/pkg/jwt"
	res "tanahore/internal/pkg/responses"

	"github.com/labstack/echo/v4"
)

func (handler *AuthHandlerImpl) RegisterUser(ctx echo.Context) error {
	registerUserRequest := web.RegisterUserRequest{}
	err := ctx.Bind(&registerUserRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request not valid", err)
	}

	response, err := handler.AuthService.RegisterUser(&registerUserRequest)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return res.StatusBadRequest(ctx, "validation error", err)
		}
		if strings.Contains(err.Error(), "already exist") {
			return res.StatusAlreadyExist(ctx, "account already exist", err)
		}
		return res.StatusInternalServerError(ctx, "failed to register user, something happen", err)
	}

	tokenAccess, err := jwt.GenerateAccessToken(response)
	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to register user, something happen", err)
	}

	loginResponse := conversion.AuthResponseToLoginResponse(response, tokenAccess)

	return res.StatusCreated(ctx, "success to register user", loginResponse, nil)
}
