package openapi

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func generateJWT(user User, jwtSecret []byte) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
