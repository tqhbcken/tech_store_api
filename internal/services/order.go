package services

import (
	"api_techstore/internal/database"
	"api_techstore/internal/models"
)

func GetAllOrders() ([]models.Order, error) {
	
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}

	var orders []models.Order
	if err := db.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrderByID(id string) (models.Order, error) {
	db, err := database.InitDB()
	if err != nil {
		return models.Order{}, err
	}

	var order models.Order
	if err := db.DB.First(&order, id).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func CreateOrder(order models.Order) (models.Order, error) {
	db, err := database.InitDB()
	if err != nil {
		return models.Order{}, err
	}

	if err := db.DB.Create(&order).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func UpdateOrder(id string, order models.Order) (models.Order, error) {
	db, err := database.InitDB()
	if err != nil {
		return models.Order{}, err
	}

	var existingOrder models.Order
	if err := db.DB.First(&existingOrder, id).Error; err != nil {
		return models.Order{}, err
	}

	order.OrderID = existingOrder.OrderID // Ensure the ID remains the same
	if err := db.DB.Save(&order).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func DeleteOrder(id string) error {
	db, err := database.InitDB()
	if err != nil {
		return err
	}

	var order models.Order
	if err := db.DB.First(&order, id).Error; err != nil {
		return err
	}

	if err := db.DB.Delete(&order).Error; err != nil {
		return err
	}
	return nil
}	

