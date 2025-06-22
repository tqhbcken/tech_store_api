package services

import (
	"api_techstore/internal/database"
	"api_techstore/internal/models"
)

func GetAllProducts() ([]models.Product, error) {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}

	//lay danh sach product
	var products []models.Product
	err = db.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductById(id string) (models.Product, error) {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return models.Product{}, err
	}

	//lay product theo id
	var product models.Product
	err = db.DB.First(&product, id).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func CreateProduct(product models.Product) (models.Product, error) {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return models.Product{}, err
	}

	//them moi product
	err = db.DB.Create(&product).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func UpdateProduct(id string, product models.Product) (models.Product, error) {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return models.Product{}, err
	}

	//cap nhat product
	err = db.DB.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
	if err != nil {
		return models.Product{}, err
	}

	//lay product da cap nhat
	var updatedProduct models.Product
	err = db.DB.First(&updatedProduct, id).Error
	if err != nil {
		return models.Product{}, err
	}

	return updatedProduct, nil
}

func DeleteProduct(id string) error {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return err
	}

	//xoa product theo id
	err = db.DB.Delete(&models.Product{}, id).Error
	if err != nil {
		return err
	}

	return nil
}