package dto

import (
	"capstone-project/model"

	"github.com/google/uuid"
)

type SpecialistRequest struct {
	Name  string `json:"name" form:"name"`
	Image string `json:"image" form:"image"`
}

type SpecialistResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Image string    `json:"image"`
}

func ConvertToSpecialistModel(specialist SpecialistRequest) model.Specialist {
	return model.Specialist{
		ID:    uuid.New(),
		Name:  specialist.Name,
		Image: specialist.Image,
	}
}

func ConvertToSpecialistResponse(specialist model.Specialist) SpecialistResponse {
	return SpecialistResponse{
		ID:    specialist.ID,
		Name:  specialist.Name,
		Image: specialist.Image,
	}
}
