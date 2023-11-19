package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DoctorProfile struct {
	gorm.Model
	ID                  uuid.UUID             `json:"id" form:"id"`
	Address             string                `gorm:"size:255"`
	Phone               string                `gorm:"size:255"`
	DoctorEducationID   []DoctorEducation     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DoctorWorkHistoryID []DoctorWorkHistory   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DoctorCertification []DoctorCertification `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Doctor              Doctor                `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
