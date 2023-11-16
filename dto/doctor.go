package dto

import (
	"capstone-project/model"

	"github.com/google/uuid"
)

type DoctorLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type DoctorLoginResponse struct {
	ID    uuid.UUID `json:"id" form:"id"`
	Email string    `json:"email" form:"email"`
	Token string    `json:"token"`
}

func ConvertToDoctorLoginResponse(doctor model.Doctor) DoctorLoginResponse {
	return DoctorLoginResponse{
		ID:    doctor.ID,
		Email: doctor.Email,
	}
}
