package repository

import (
	"errors"

	"capstone-project/database"
	"capstone-project/middleware"
	"capstone-project/model"
	"capstone-project/constant"

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
		token, errToken = middleware.CreateToken(data.ID, constant.ROLE_USER, data.Name, false)
		if errToken != nil {
			return model.User{}, "", errToken
		}
	}
	return data, token, nil
}

func CheckAdmin(email string, password string) (model.User, string, error) {
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
		token, errToken = middleware.CreateToken(data.ID, constant.ROLE_ADMIN, data.Name, false)
		if errToken != nil {
			return model.User{}, "", errToken
		}
	}
	return data, token, nil
}

func CreateUser(data model.User) (model.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
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
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	tx := database.DB.Model(&data).Where("id = ?", data.ID).Updates(map[string]interface{}{"password": string(hashPassword)})
	if tx.Error != nil {
		return model.User{}, tx.Error
	}
	return data, nil
}

func CheckUserEmail(email string) bool {
	var data model.User

	tx := database.DB.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return false
	}

	return true
}

func SetOTP(email, otp string) error {
	if !CheckUserEmail(email) {
        return errors.New("user email not found")
    }

	tx := database.DB.Model(&model.User{}).Where("email = ?", email).Update("OTP", otp)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func ValidateOTP(email, otp string) (model.User, string, error) {
	var data model.User

	tx := database.DB.Where("email = ? AND otp = ?", email, otp).First(&data)
	if tx.Error != nil {
		return model.User{}, "", errors.New("Invalid Email or OTP")
	}

	database.DB.Model(&model.User{}).Where("email = ?", email).Update("OTP", nil)
	if tx.Error != nil {
		return model.User{}, "", tx.Error
	}

	token, err := middleware.CreateToken(data.ID, constant.ROLE_USER, data.Name, false)
		if err != nil {
			return model.User{}, "", err
		}

	return data, token, nil
}