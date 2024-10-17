package utils

import (
	"github.com/doryngal/CourseFlow/backend/internal/auth-service/models"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func GenerateNewAccessToken(user models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	hoursCount := 24

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
