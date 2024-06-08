package service

import (
	"tanahore/internal/app/auth/repository"
	"tanahore/internal/model/web"

	"github.com/go-playground/validator"
)

type AuthService interface {
	RegisterUser(req *web.RegisterUserRequest) (*web.AuthResponse, error)
	LoginUser(req *web.LoginUserRequest) (*web.AuthResponse, *web.UserLoginResponse, error)
}

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	Validate       *validator.Validate
}

func NewAuthService(repo repository.AuthRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: repo,
		Validate:       validate,
	}
}
