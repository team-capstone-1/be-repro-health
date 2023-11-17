package repository

import (
	"capstone-project/database"
	"capstone-project/middleware"
	"capstone-project/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func CheckDoctor(email string, password string) (model.Doctor, string, error) {
	var data model.Doctor

	tx := database.DB.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return model.Doctor{}, "", errors.New("Invalid Email or Password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(password))
	if err != nil {
		return model.Doctor{}, "", errors.New("Invalid Email or Password")
	}

	var token string
	if tx.RowsAffected > 0 {
		var errToken error
		token, errToken = middleware.CreateTokenDoctor(data.ID, "doctor")
		if errToken != nil {
			return model.Doctor{}, "", errToken
		}
	}
	return data, token, nil
}
