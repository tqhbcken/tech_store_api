package services

import (
	"api_techstore/internal/database"
	"api_techstore/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func Login() {

}

// /
// Check function
// /
func GetUserByEmail(email string) ([]models.User, error) {
	///
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}
	var user []models.User
	if err := db.DB.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
