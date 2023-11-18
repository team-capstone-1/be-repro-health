package dto

import (
	"time"

	"capstone-project/model"

	"github.com/google/uuid"
)

type PatientRequest struct {
	Name     		   string    `json:"name" form:"name"`
	ProfileImage       string    `json:"profile_image" form:"profile_image"`
	DateOfBirth 	   time.Time `json:"date_of_birth" form:"date_of_birth"`
	Relation		   string    `json:"relation" form:"relation"`
	Weight			   float64   `json:"weight" form:"weight"`
	Height			   float64   `json:"height" form:"height"`
	KTPImage		   string    `json:"ktp_image" form:"ktp_image"`
	NIK				   string    `json:"nik" form:"nik"`
	NoKartuKeluarga    string    `json:"no_kartu_keluarga" form:"no_kartu_keluarga"`
	KartuKeluargaImage string    `json:"kartu_keluarga_image" form:"kartu_keluarga_image"`
}

type PatientResponse struct {
	ID    			   uuid.UUID `json:"id"`
	Name  			   string    `json:"name"`
	ProfileImage       string    `json:"profile_image"`
	DateOfBirth 	   time.Time `json:"date_of_birth"`
	Relation		   string    `json:"relation"`
	Weight			   float64   `json:"weight"`
	Height			   float64   `json:"height"`
	KTPImage		   string    `json:"ktp_image"`
	NIK				   string    `json:"nik"`
	NoKartuKeluarga    string    `json:"no_kartu_keluarga"`
	KartuKeluargaImage string    `json:"kartu_keluarga_image"`
}

func ConvertToPatientModel(patient PatientRequest) model.Patient {
	return model.Patient{
		ID:      		    uuid.New(),
		Name:    		    patient.Name,
		ProfileImage:       patient.ProfileImage,
		DateOfBirth: 	    patient.DateOfBirth,
		Relation:		    patient.Relation,
		Weight:			    patient.Weight,
		Height:			    patient.Height,
		KTPImage:		    patient.KTPImage,
		NIK:			    patient.NIK,
		NoKartuKeluarga:    patient.NoKartuKeluarga,
		KartuKeluargaImage: patient.KartuKeluargaImage,
	}
}

func ConvertToPatientResponse(patient model.Patient) PatientResponse {
	return PatientResponse{
		ID:    patient.ID,
		Name:  patient.Name,
		ProfileImage:       patient.ProfileImage,
		DateOfBirth: 	    patient.DateOfBirth,
		Relation:		    patient.Relation,
		Weight:			    patient.Weight,
		Height:			    patient.Height,
		KTPImage:		    patient.KTPImage,
		NIK:			    patient.NIK,
		NoKartuKeluarga:    patient.NoKartuKeluarga,
		KartuKeluargaImage: patient.KartuKeluargaImage,
	}
}
