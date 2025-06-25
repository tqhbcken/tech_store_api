package services

import (
	"api_techstore/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type UserService interface {
	CheckUserExists(id string) (bool, error)
	GetAllUsers() ([]models.User, error)
	GetUserById(id string) (models.User, error)
	CreateUser(user models.User) error
	UpdateUser(id string, user models.User) error
	DeleteUser(id string) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (s *userService) CheckUserExists(id string) (bool, error) {
	var exists bool
	err := s.db.Model(&models.User{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).
		Error
	return exists, err
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) GetUserById(id string) (models.User, error) {
	var user models.User
	if err := s.db.First(&user, "id = ?", id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *userService) CreateUser(user models.User) error {
	if err := s.db.Create(&user).Error; err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (s *userService) UpdateUser(id string, user models.User) error {
	if err := s.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

func (s *userService) DeleteUser(id string) error {
	if err := s.db.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}
