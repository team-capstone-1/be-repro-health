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

type DoctorSignUpRequest struct {
	Name         string    `json:"name" form:"name"`
	Email        string    `json:"email" form:"email"`
	Password     string    `json:"password" form:"password"`
	Price        float64   `json:"price" form:"price"`
	Address      string    `json:"address" form:"address"`
	Phone        string    `json:"phone" form:"phone"`
	ProfileImage string    `json:"profile_image"`
	SpecialistID uuid.UUID `json:"specialist_id" form:"specialist_id"`
	ClinicID     uuid.UUID `json:"clinic_id" form:"clinic_id"`
}

type DoctorSignUpResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	SpecialistID uuid.UUID `json:"specialist_id"`
	ClinicID     uuid.UUID `json:"clinic_id"`
}

type DoctorResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Price        float64   `json:"price"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	SpecialistID uuid.UUID `json:"specialist_id"`
	ClinicID     uuid.UUID `json:"clinic_id"`
}

func ConvertToDoctorSignUpResponse(doctor model.Doctor) DoctorSignUpResponse {
	return DoctorSignUpResponse{
		ID:           doctor.ID,
		Name:         doctor.Name,
		Email:        doctor.Email,
		SpecialistID: doctor.SpecialistID,
		ClinicID:     doctor.ClinicID,
	}
}

func ConvertToDoctorLoginResponse(doctor model.Doctor) DoctorLoginResponse {
	return DoctorLoginResponse{
		ID:    doctor.ID,
		Email: doctor.Email,
	}
}

func ConvertToDoctorModel(doctor DoctorSignUpRequest) model.Doctor {
	return model.Doctor{
		ID:           uuid.New(),
		Name:         doctor.Name,
		Email:        doctor.Email,
		Password:     doctor.Password,
		Price:        doctor.Price,
		Address:      doctor.Address,
		Phone:        doctor.Phone,
		ProfileImage: doctor.ProfileImage,
		SpecialistID: doctor.SpecialistID,
		ClinicID:     doctor.ClinicID,
	}
}

func ConvertToDoctorResponse(doctor model.Doctor) DoctorResponse {
	return DoctorResponse{
		ID:           doctor.ID,
		Name:         doctor.Name,
		Email:        doctor.Email,
		Price:        doctor.Price,
		Address:      doctor.Address,
		Phone:        doctor.Phone,
		SpecialistID: doctor.SpecialistID,
		ClinicID:     doctor.ClinicID,
	}
}
