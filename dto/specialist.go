package dto

import (
	"github.com/google/uuid"
	"capstone-project/model"
)

type SpecialistResponse struct {
	ID 		 uuid.UUID `json:"id"`
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