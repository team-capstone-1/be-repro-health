package model

import "gorm.io/gorm"

type DoctorProfile struct {
	gorm.Model
	Address             string                `gorm:"size:255"`
	Phone               string                `gorm:"size:255"`
	DoctorEducationID   []DoctorEducation     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DoctorWorkHistoryID []DoctorWorkHistory   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DoctorCertification []DoctorCertification `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Doctor              Doctor                `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
