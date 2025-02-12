package openapi

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

func AuthOrRegisterUser(db *gorm.DB, req AuthRequest, jwtSecret []byte) (string, error) {
	var user User
	err := db.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("No user found, creating new user.")
			if err := createUser(db, req); err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	} else {
		log.Printf("User found, authenticating with username %s.", req.Username)
		if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			return "", errors.New("invalid credentials")
		}
	}

	token, err := generateJWT(user, jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
