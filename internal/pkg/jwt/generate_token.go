package jwt

import (
	"os"
	"tanahore/internal/model/web"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateAccessToken(userLoginResponse *web.AuthResponse) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = userLoginResponse.ID
	claims["username"] = userLoginResponse.Username
	claims["email"] = userLoginResponse.Email
	claims["role_name"] = userLoginResponse.RoleName
	claims["created_at"] = userLoginResponse.CreatedAt
	claims["exp"] = expireTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}
