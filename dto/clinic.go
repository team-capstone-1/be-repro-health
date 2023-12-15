package dto

import (
	"capstone-project/model"
	"github.com/google/uuid"
)

type ClinicResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	City      string    `json:"city"`
	Location  string    `json:"location"`
	Telephone string    `json:"telephone"`
	Email     string    `json:"email"`
	Profile   string    `json:"profile"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
}

func ConvertToClinicResponse(clinic model.Clinic) ClinicResponse {
	return ClinicResponse{
		ID:        clinic.ID,
		Name:      clinic.Name,
		Image:     clinic.Image,
		City:      clinic.City,
		Location:  clinic.Location,
		Telephone: clinic.Telephone,
		Email:     clinic.Email,
		Profile:   clinic.Profile,
		Latitude:  clinic.Latitude,
		Longitude: clinic.Longitude,
	}
}
