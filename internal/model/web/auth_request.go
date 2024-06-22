package web

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"required,email" form:"email"`
	Username string `json:"username" validate:"required" form:"username"`
	Password string `json:"password" validate:"required,min=8,max=255" form:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email" form:"email"`
	Password string `json:"password" validate:"required,min=8,max=255" form:"password"`
}
