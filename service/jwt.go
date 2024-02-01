package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService struct {
}

type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (service *JWTService) GenerateToken(username string) (string, error) {
	claims := &CustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
