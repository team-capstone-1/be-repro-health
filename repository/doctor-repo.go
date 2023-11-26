package repository

import (
	"capstone-project/database"
	"capstone-project/middleware"
	"capstone-project/model"
	"errors"

	"github.com/google/uuid"
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
		token, errToken = middleware.CreateToken(data.ID, "doctor", data.Name)
		if errToken != nil {
			return model.Doctor{}, "", errToken
		}
	}
	return data, token, nil
}

func CheckDoctorEmail(email string) bool {
	var data model.Doctor

	tx := database.DB.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return false
	}

	return true
}

func CreateDoctor(data model.Doctor) (model.Doctor, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.Doctor{}, err
	}
	data.Password = string(hashPassword)

	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Doctor{}, tx.Error
	}
	return data, nil
}

func GetClinicByDoctorID(id uuid.UUID) (model.Clinic, error) {
	var datadoctor model.Doctor
	var dataclinic model.Clinic

	tx := database.DB.First(&datadoctor, id)
	if tx.Error != nil {
		return model.Clinic{}, tx.Error
	}

	tx = database.DB.First(&dataclinic, datadoctor.ClinicID)
	if tx.Error != nil {
		return model.Clinic{}, tx.Error
	}
	return dataclinic, nil
}

func GetAllDoctors(name string) ([]model.Doctor, error) {
	var datadoctors []model.Doctor

	tx := database.DB

	if name != "" {
		tx = tx.Where("name LIKE ?", "%"+name+"%")
	}

	tx.Find(&datadoctors)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datadoctors, nil
}

func GetDoctorByID(id uuid.UUID) (model.Doctor, error) {
	var datadoctor model.Doctor

	tx := database.DB.First(&datadoctor, id)
	if tx.Error != nil {
		return model.Doctor{}, tx.Error
	}
	return datadoctor, nil
}

func GetDoctorsBySpecialist(id uuid.UUID) ([]model.Doctor, error) {
	var datadoctors []model.Doctor

	tx := database.DB.Where("specialist_id = ?", id).Find(&datadoctors)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datadoctors, nil
}

func GetDoctorsByClinic(id uuid.UUID) ([]model.Doctor, error) {
	var datadoctors []model.Doctor

	tx := database.DB.Where("clinic_id = ?", id).Find(&datadoctors)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datadoctors, nil
}
