package utils

import (
	"os"
	"test_marketplace/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(users models.User) (string, error) {
	claims := &models.JWTClaims{
		UserID: users.ID,
		Email:  users.Email,
		Role:   users.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "marketplace",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret-key"
	}
	return token.SignedString([]byte(secret))
}

func ValidateJWT(tokenString string) (*models.JWTClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret-key"
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
