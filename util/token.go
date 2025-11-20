package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CreateToken generates a new JWT token for a user
func CreateToken(phone string, duration time.Duration, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone": phone,
		"exp":   time.Now().Add(duration).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}
