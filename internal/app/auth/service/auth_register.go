package service

import (
	"errors"
	"tanahore/internal/model/web"
	converterReq "tanahore/internal/pkg/converter/request"
	converterRes "tanahore/internal/pkg/converter/response"
)

func (service *AuthServiceImpl) RegisterUser(req *web.RegisterUserRequest) (*web.AuthResponse, error) {
	if err := service.Validate.Struct(req); err != nil {
		return nil, err
	}
	if exists := service.AuthRepository.UserExists(req.Username, req.Email); exists {
		return nil, errors.New("user exists")
	}
	user := converterReq.RegisterUserRequestToUserModel(req)
	if err := service.AuthRepository.RegisterUser(user); err != nil {
		return nil, err
	}
	res := converterRes.UserToAuthResponse(user)
	return res, nil
}
