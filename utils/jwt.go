package utils

import (
	"time"

	"github.com/idealana/go-project-management/config"

	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int64, role, email string, publicID uuid.UUID) (string, error) {
	secret := config.AppConfig.JWTSecret
	duration, _ := time.ParseDuration(config.AppConfig.JWTExpired)

	claims := jwt.MapClaims{
		"user_id": userID,
		"role": role,
		"email": email,
		"pub_id": publicID,
		"exp": time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SigningString([]byte(secret))
}

func GenerateRefreshToken(userID int64) (string, error) {
	secret := config.AppConfig.JWTSecret
	duration, _ := time.ParseDuration(config.AppConfig.JWTRefreshToken)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SigningString([]byte(secret))
}
