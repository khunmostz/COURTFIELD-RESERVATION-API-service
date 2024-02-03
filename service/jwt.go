package service

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

type JWTService struct {
}

func (service *JWTService) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (service *JWTService) ValidateToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("invalid token signing method")
		}

		// Return the secret key
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Return the token claims
	claims, ok := token.Claims.(jwt.Claims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
