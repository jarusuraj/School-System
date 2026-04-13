package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string,user_id int ,role string) (string, error) {

	claims := jwt.MapClaims{
		"email": email,
		"user_id":user_id,
		"role":role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("SECRET_KEY")
	return token.SignedString([]byte(secret))
}
