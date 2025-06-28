package services

import (
	"api_techstore/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(db *gorm.DB, email, password string) (*models.User, error) {
	users, err := GetUserByEmail(db, email)
    if err != nil {
        return nil, err
    }
    if len(users) == 0 {
        return nil, gorm.ErrRecordNotFound
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

// HashPassword hashes a plain password (auth-specific)
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CreateUser creates a new user in the database (auth-specific)
func CreateUser(db *gorm.DB, user models.User) error {
	return db.Create(&user).Error
}
