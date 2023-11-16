package repository

import (
	"errors"

	"capstone-project/database"
	"capstone-project/middleware"
	"capstone-project/model"

	"golang.org/x/crypto/bcrypt"
)

func CheckUser(email string, password string) (model.User, string, error) {
	var data model.User
	
	tx := database.DB.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return model.User{}, "", errors.New("Invalid Email or Password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(password))
	if err != nil {
		return model.User{}, "", errors.New("Invalid Email or Password")
	}

	var token string
	if tx.RowsAffected > 0 {
		var errToken error
		token, errToken = middleware.CreateToken(int(data.ID), "user")
		if errToken != nil {
			return model.User{}, "", errToken
		}
	}
	return data, token, nil
}

func CreateUser(data model.User) (model.User, error) {
	hashPassword,err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}
	data.Password = string(hashPassword)

	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.User{}, tx.Error
	}
	return data, nil
}

func UpdateUserPassword(data model.User) (model.User, error) {
	hashPassword,err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	tx := database.DB.Model(&data).Where("email = ?", data.Email).Updates(map[string]interface{}{"password": string(hashPassword)})
	if tx.Error != nil {
		return model.User{}, tx.Error
	}
	return data, nil
}

func CheckUserEmail(email string) (bool) {
	var data model.User
	
	tx := database.DB.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return false
	}

	return true
}