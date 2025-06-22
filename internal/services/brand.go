package services

import (
	"api_techstore/internal/database"
	"api_techstore/internal/models"
)

func GetAllBrands() ([]models.Brand, error) {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}

	//lay danh sach brand
	var brands []models.Brand
	err = db.DB.Find(&brands).Error
	if err != nil {
		return nil, err
	}
	return brands, nil
}

func GetBrandById(id string) (models.Brand, error) {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return models.Brand{}, err
	}

	//lay brand theo id
	var brand models.Brand
	err = db.DB.First(&brand, id).Error
	if err != nil {
		return models.Brand{}, err
	}
	return brand, nil
}

func CreateBrand(brand models.Brand) (models.Brand, error) {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return models.Brand{}, err
	}

	//them moi brand
	err = db.DB.Create(&brand).Error
	if err != nil {
		return models.Brand{}, err
	}
	return brand, nil
}

func UpdateBrand(id string, brand models.Brand) (models.Brand, error) {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return models.Brand{}, err
	}

	//cap nhat brand
	err = db.DB.Model(&models.Brand{}).Where("id = ?", id).Updates(brand).Error
	if err != nil {
		return models.Brand{}, err
	}

	//lay brand da cap nhat
	var updatedBrand models.Brand
	err = db.DB.First(&updatedBrand, id).Error
	if err != nil {
		return models.Brand{}, err
	}
	return updatedBrand, nil
}

func DeleteBrand(id string) error {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return err
	}

	//xoa brand theo id
	err = db.DB.Delete(&models.Brand{}, id).Error
	if err != nil {
		return err
	}
	return nil
}