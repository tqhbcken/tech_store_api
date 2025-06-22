package services

import (
	"api_techstore/internal/database"
	"api_techstore/internal/models"
)

func GetAllCategories() ([]models.Category, error) {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}

	//lay danh sach category
	var categories []models.Category
	err = db.DB.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoryById(id string) (models.Category, error) {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return models.Category{}, err
	}

	//lay category theo id
	var category models.Category
	err = db.DB.First(&category, id).Error
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func CreateCategory(category models.Category) (models.Category, error) {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return models.Category{}, err
	}

	//them moi category
	err = db.DB.Create(&category).Error
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func UpdateCategory(id string, category models.Category) (models.Category, error) {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return models.Category{}, err
	}

	//cap nhat category theo id
	err = db.DB.Model(&models.Category{}).Where("id = ?", id).Updates(category).Error
	if err != nil {
		return models.Category{}, err
	}

	//tra ve category da cap nhat
	var updatedCategory models.Category
	err = db.DB.First(&updatedCategory, id).Error
	if err != nil {
		return models.Category{}, err
	}
	return updatedCategory, nil
}

func DeleteCategory(id string) error {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return err
	}

	//xoa category theo id
	err = db.DB.Delete(&models.Category{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

