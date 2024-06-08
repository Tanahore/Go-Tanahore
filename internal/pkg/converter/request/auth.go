package request

import (
	"tanahore/internal/model/domain"
	"tanahore/internal/model/web"
)

func RegisterUserRequestToUserModel(userRequest *web.RegisterUserRequest) *domain.Users {
	return &domain.Users{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		RoleName: "user",
	}
}
