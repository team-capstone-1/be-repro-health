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

type DoctorProfileResponse struct {
	ID           uuid.UUID                       `json:"id"`
	Name         string                          `json:"name"`
	Address      string                          `json:"address"`
	Email        string                          `json:"email"`
	Phone        string                          `json:"phone"`
	ProfileImage string                          `json:"profile_image"`
	SpecialistID uuid.UUID                       `json:"specialist_id"`
	Specialist   DoctorProfileSpecialistResponse `json:"specialist"`
	ClinicID     uuid.UUID                       `json:"clinic_id"`
	Clinic       DoctorProfileClinicResponse     `json:"clinic"`
}

type DoctorProfileClinicResponse struct {
	Name     string `json:"name"`
	City     string `json:"city"`
	Location string `json:"location"`
	Profile  string `json:"profile"`
}

type DoctorProfileSpecialistResponse struct {
	Name  string `json:"name"`
	Image string `json:"image"`
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

func ConvertToDoctorProfileResponse(doctor model.Doctor) DoctorProfileResponse {
	return DoctorProfileResponse{
		ID:           doctor.ID,
		Name:         doctor.Name,
		Address:      doctor.Address,
		Email:        doctor.Email,
		Phone:        doctor.Phone,
		ProfileImage: doctor.ProfileImage,
		SpecialistID: doctor.SpecialistID,
		ClinicID:     doctor.ClinicID,
		Specialist:   DoctorProfileSpecialistResponse{Name: doctor.Specialist.Name, Image: doctor.Specialist.Image},
		Clinic:       DoctorProfileClinicResponse{Name: doctor.Clinic.Name, City: doctor.Clinic.City, Location: doctor.Clinic.Location, Profile: doctor.Clinic.Location},
	}
}
