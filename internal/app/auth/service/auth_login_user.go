package service

import (
	"errors"
	"tanahore/internal/model/web"
	converter "tanahore/internal/pkg/converter/response"
)

func (service *AuthServiceImpl) LoginUser(req *web.LoginUserRequest) (*web.AuthResponse, *web.UserLoginResponse, error) {
	if err := service.Validate.Struct(req); err != nil {
		return nil, nil, err
	}
	user := service.AuthRepository.UserExists("", req.Email)
	if user == nil {
		return nil, nil, errors.New("user not found")
	}
	if user.Password != req.Password {
		return nil, nil, errors.New("wrong password")
	}
	auth := converter.UserToAuthResponse(user)
	loginRes := converter.UserToLoginResponse(user)
	return auth, loginRes, nil
}
