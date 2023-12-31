package repository

import (
	"capstone-project/constant"
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
		token, errToken = middleware.CreateToken(data.ID, constant.ROLE_DOCTOR, data.Name, true)
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

	tx := database.DB.Preload("Clinic").Preload("Specialist").Preload("DoctorWorkHistories").Preload("DoctorEducations").First(&datadoctor, id)
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

	tx := database.DB.Preload("Clinic").Preload("Specialist").Preload("DoctorWorkHistories").Preload("DoctorEducations")

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

	tx := database.DB.Preload("Clinic").Preload("Specialist").Preload("DoctorWorkHistories").Preload("DoctorEducations").First(&datadoctor, id)
	if tx.Error != nil {
		return model.Doctor{}, tx.Error
	}
	return datadoctor, nil
}

func GetDoctorsBySpecialist(id uuid.UUID) ([]model.Doctor, error) {
	var datadoctors []model.Doctor

	tx := database.DB.Preload("Clinic").Preload("Specialist").Preload("DoctorWorkHistories").Preload("DoctorEducations").Where("specialist_id = ?", id).Find(&datadoctors)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datadoctors, nil
}

func GetDoctorsByClinic(id uuid.UUID) ([]model.Doctor, error) {
	var datadoctors []model.Doctor

	tx := database.DB.Preload("Clinic").Preload("Specialist").Preload("DoctorWorkHistories").Preload("DoctorEducations").Where("clinic_id = ?", id).Find(&datadoctors)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datadoctors, nil
}

func SetDoctorOTP(email, otp string) error {
	if !CheckDoctorEmail(email) {
		return errors.New("user email not found")
	}

	tx := database.DB.Model(&model.Doctor{}).Where("email = ?", email).Update("OTP", otp)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func ValidateDoctorOTP(email, otp string) (model.Doctor, string, error) {
	var data model.Doctor

	tx := database.DB.Where("email = ? AND otp = ?", email, otp).First(&data)
	if tx.Error != nil {
		return model.Doctor{}, "", errors.New("Invalid Email or OTP")
	}

	database.DB.Model(&model.Doctor{}).Where("email = ?", email).Update("OTP", nil)
	if tx.Error != nil {
		return model.Doctor{}, "", tx.Error
	}

	token, err := middleware.CreateToken(data.ID, constant.ROLE_DOCTOR, data.Name, true)
	if err != nil {
		return model.Doctor{}, "", err
	}

	return data, token, nil
}

func UpdateDoctorPassword(data model.Doctor) (model.Doctor, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.Doctor{}, err
	}

	tx := database.DB.Model(&data).Where("id = ?", data.ID).Updates(map[string]any{"password": string(hashPassword)})
	if tx.Error != nil {
		return model.Doctor{}, tx.Error
	}

	tx = database.DB.Where("id = ?", data.ID).First(&data)
	if tx.Error != nil {
		return model.Doctor{}, tx.Error
	}

	return data, nil
}
func GetDoctorsBySpecialistAndClinic(specialist_id, clinic_id uuid.UUID) ([]model.Doctor, error) {
	var datadoctors []model.Doctor

	tx := database.DB.Preload("Clinic").Preload("Specialist").Preload("DoctorWorkHistories").Preload("DoctorEducations").Where("specialist_id = ? AND clinic_id = ?", specialist_id, clinic_id).Find(&datadoctors)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datadoctors, nil
}

func GetProfileByDoctorID(id uuid.UUID) string {
	var data model.Doctor
	database.DB.Where("id = ?", id).First(&data)

	return data.ProfileImage
}