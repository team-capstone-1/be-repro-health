package dto

import (
	"github.com/google/uuid"
	"capstone-project/model"
)

type ClinicResponse struct {
	ID 		 uuid.UUID `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Location    string `json:"location"`
	Profile    string `json:"profile"`
	Latitude    string `json:"latitude"`
	Longitude    string `json:"longitude"`
}

func ConvertToClinicResponse(clinic model.Clinic) ClinicResponse {
	return ClinicResponse{
		ID:    clinic.ID,
		Name:  clinic.Name,
		City: clinic.City,
		Location: clinic.Location,
		Profile: clinic.Profile,
		Latitude: clinic.Latitude,
		Longitude: clinic.Longitude,
	}
}