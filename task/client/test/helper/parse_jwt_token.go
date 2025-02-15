package helper

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		err := godotenv.Load()
		if err != nil {
			log.Printf("Error loading .env file. If you're in Docker, then it's fine")
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			return nil, errors.New("JWT_SECRET is not set")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if username, ok := claims["username"].(string); ok {
			return username, nil
		}
		return "", errors.New("user_id not found in token claims")
	}

	return "", errors.New("invalid token")
}
