package dto

import (
	"capstone-project/model"
)

type SpecialistResponse struct {
	ID 		 uint `json:"id"`
	Name    string `json:"name"`
	Image    string `json:"image"`
}

func ConvertToSpecialistResponse(specialist model.Specialist) SpecialistResponse {
	return SpecialistResponse{
		ID:    specialist.ID,
		Name:  specialist.Name,
		Image: specialist.Image,
	}
}