package controller

import (
	"capstone-project/constant"
	"capstone-project/database"
	"capstone-project/middleware"
	"capstone-project/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func InsertDataDoctor() string {
	doctor := model.Doctor{
		ID:       uuid.New(),
		Name:     "Andi",
		Email:    "andicahyo@gmail.com",
		Password: "Doctor@123",
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(doctor.Password), bcrypt.DefaultCost)
	doctor.Password = string(hashPassword)

	database.DB.Create(&doctor)

	token, _ := middleware.CreateToken(doctor.ID, constant.ROLE_DOCTOR, doctor.Name, false)
	return token
}
