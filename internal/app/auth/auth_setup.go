package auth

import (
	handlerPkg "tanahore/internal/app/auth/handler"
	repositoryPkg "tanahore/internal/app/auth/repository"
	routesPkg "tanahore/internal/app/auth/routes"
	servicePkg "tanahore/internal/app/auth/service"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func AuthSetup(db *gorm.DB, validate *validator.Validate) routesPkg.AuthRoutes {
	AuthRepo := repositoryPkg.NewAuthRepository(db)
	AuthService := servicePkg.NewAuthService(AuthRepo, validate)
	AuthHandler := handlerPkg.NewAuthHandler(AuthService)
	AuthRoute := routesPkg.NewAuthRoutes(AuthHandler)

	return AuthRoute
}
