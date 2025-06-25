package services

import (
	"api_techstore/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(db *gorm.DB, email, password string) (*models.User, error) {
	
	users, err := GetUserByEmail(db, email)
	if err != nil || len(users) == 0 {
		return nil, err
	}

	user := users[0]
	
	if !CheckPasswordHash(password, user.PasswordHash) {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}


func GetUserByEmail(db *gorm.DB, email string) ([]models.User, error) {
	var user []models.User
	if err := db.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
