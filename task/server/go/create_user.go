package openapi

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func createUser(db *gorm.DB, req AuthRequest) error {
	// Пользователь не найден – создаём нового
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := User{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
		Balance:      1000,
	}

	if err = db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
