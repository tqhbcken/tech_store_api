package services

import (
	"api_techstore/internal/database"
	"api_techstore/internal/models"
	"errors"
	"fmt"
)

func GetAllUsers() ([]models.User, error) {
	//open db
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("DB Init Error:", err)
		return nil, err
	}

	//tao moi user
	users := []models.User{}
	db.DB.Find(&users)

	//neu khong tim thay user thi tra ve mang rong (khong phai loi)
	if len(users) == 0 {
		return []models.User{}, nil
	}
	//neu tim thay user thi tra ve user
	return users, nil
}

func GetUserById(id string) (models.User, error) {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return models.User{}, err
	}

	//tim user theo id
	var user models.User
	db.DB.First(&user, id).Where("id = ?", id)

	//neu khong tim thay user thi tra ve loi
	if user.ID == 0 {
		return models.User{}, errors.New("user not found")
	}

	//neu tim thay user thi tra ve user
	return user, nil
}

func CreateUser(user models.User) error {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return err
	}

	// //debug: in ra user trước khi tạo
	// fmt.Println("Creating user:", user)

	//tao moi user
	db.DB.Create(&user)

	//neu co loi thi tra ve loi
	if db.DB.Error != nil {
		fmt.Println("DB Error:", db.DB.Error)
		return db.DB.Error
	}

	//neu khong co loi thi tra ve user
	if user.ID == 0 {
		return errors.New("failed to create user")
	}
	//neu co loi thi tra ve loi
	if db.DB.Error != nil {
		//modify theo response trong pkg
		return nil
	}
	//neu khong co loi thi tra ve user
	return nil
}

func UpdateUser(id string, user models.User) error {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return err
	}

	//update user
	db.DB.Model(&models.User{}).Where("id = ?", id).Updates(user)

	//neu co loi thi tra ve loi
	if db.DB.Error != nil {
		return db.DB.Error
	}

	//neu khong co loi thi tra ve user
	return nil
}

func DeleteUser(id string) error {
	//open db
	db, err := database.InitDB()
	if err != nil {
		return err
	}

	//xoa user
	db.DB.Delete(&models.User{}, id)

	//neu co loi thi tra ve loi
	if db.DB.Error != nil {
		return db.DB.Error
	}

	//neu khong co loi thi tra ve user
	return nil
}


