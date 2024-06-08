package response

import (
	"tanahore/internal/model/domain"
	"tanahore/internal/model/web"
)

func UserToAuthResponse(user *domain.Users) *web.AuthResponse {
	return &web.AuthResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		RoleName:  web.Role(user.RoleName),
		CreatedAt: user.CreatedAt,
	}
}

func AuthResponseToLoginResponse(auth *web.AuthResponse, token string) *web.UserLoginResponse {
	return &web.UserLoginResponse{
		ID:          auth.ID,
		Username:    auth.Username,
		AccessToken: token,
	}
}

func UserToLoginResponse(user *domain.Users) *web.UserLoginResponse {
	return &web.UserLoginResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}
