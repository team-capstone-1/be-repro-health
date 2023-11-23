package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type DoctorWorkHistoryRequest struct {
	StartingDate time.Time `json:"start_date"`
	EndingDate   time.Time `json:"end_date"`
	Job          string    `json:"job"`
	Workplace    string    `json:"workplace"`
	Position     string    `json:"position"`
}

type DoctorEducationRequest struct {
	StartingDate     time.Time `json:"start_date"`
	EndingDate       time.Time `json:"end_date"`
	EducationProgram string    `json:"education_program"`
	University       string    `json:"university"`
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

type DoctorWorkHistoryResponse struct {
	ID              uuid.UUID `json:"id"`
	DoctorProfileID uuid.UUID `json:"doctor_profile_id"`
	StartingDate    time.Time `json:"start_date"`
	EndingDate      time.Time `json:"end_date"`
	Job             string    `json:"job"`
	Workplace       string    `json:"workplace"`
	Position        string    `json:"position"`
}

type DoctorEducationResponse struct {
	ID               uuid.UUID `json:"id"`
	DoctorProfileID  uuid.UUID `json:"doctor_profile_id"`
	StartingDate     time.Time `json:"start_date"`
	EndingDate       time.Time `json:"end_date"`
	EducationProgram string    `json:"education_program"`
	University       string    `json:"university"`
}

type DoctorCertificationResponse struct {
	ID              uuid.UUID `json:"id"`
	DoctorProfileID uuid.UUID `json:"doctor_profile_id"`
	StartingDate    time.Time `json:"start_date"`
	EndingDate      time.Time `json:"end_date"`
	Description     string    `json:"description"`
	CertificateType string    `json:"certificate_type"`
	FileSize        string    `json:"file_size"`
	Details         string    `json:"details"`
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

func ConvertToDoctorWorkHistoryModel(workHistory DoctorWorkHistoryRequest) model.DoctorWorkHistory {
	return model.DoctorWorkHistory{
		ID:              uuid.New(),
		StartingDate:    workHistory.StartingDate,
		EndingDate:      workHistory.EndingDate,
		Job:             workHistory.Job,
		Workplace:       workHistory.Workplace,
		Position:        workHistory.Position,
	}
}

func ConvertToDoctorEducationModel(education DoctorEducationRequest) model.DoctorEducation {
	return model.DoctorEducation{
		ID:               uuid.New(),
		StartingDate:     education.StartingDate,
		EndingDate:       education.EndingDate,
		EducationProgram: education.EducationProgram,
		University:       education.University,
	}
}

func ConvertToDoctorWorkHistoriesResponse(workHistory model.DoctorWorkHistory) DoctorWorkHistoryResponse {
	return DoctorWorkHistoryResponse{
		ID:              workHistory.ID,
		DoctorProfileID: workHistory.DoctorProfileID,
		StartingDate:    workHistory.StartingDate,
		EndingDate:      workHistory.EndingDate,
		Job:             workHistory.Job,
		Workplace:       workHistory.Workplace,
		Position:        workHistory.Position,
	}
}

func ConvertToDoctorEducationResponse(doctorEducation model.DoctorEducation) DoctorEducationResponse {
	return DoctorEducationResponse{
		ID:               doctorEducation.ID,
		DoctorProfileID:  doctorEducation.ID,
		StartingDate:     doctorEducation.StartingDate,
		EndingDate:       doctorEducation.EndingDate,
		EducationProgram: doctorEducation.EducationProgram,
		University:       doctorEducation.University,
	}
}

func ConvertToDoctorCertificationResponse(certification model.DoctorCertification) DoctorCertificationResponse {
	return DoctorCertificationResponse{
		ID:              certification.ID,
		DoctorProfileID: certification.DoctorProfileID,
		StartingDate:    certification.StartingDate,
		EndingDate:      certification.EndingDate,
		Description:     certification.Description,
		CertificateType: certification.CertificateType,
		FileSize:        certification.FileSize,
		Details:         certification.Details,
	}
}
