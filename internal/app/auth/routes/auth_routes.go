package routes

import "github.com/labstack/echo/v4"

func (routes *AuthRoutesImpl) MobileRoutes(apiGroup *echo.Group) {
	AuthRoutes := apiGroup.Group("/api")

	AuthRoutes.POST("/register", routes.AuthHandler.RegisterUser)
	AuthRoutes.POST("/login", routes.AuthHandler.LoginUser)
}
