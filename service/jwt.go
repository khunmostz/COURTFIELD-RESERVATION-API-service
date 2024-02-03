package service

import (
	"errors"
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

func (service *JWTService) ValidateTokenExpire(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return errors.New("Invalid token")
	}

	if time.Now().Unix() > claims.ExpiresAt {
		return errors.New("Token has expired")
	}

	return nil
}
