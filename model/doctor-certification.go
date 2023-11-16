package model

import (
	"time"

	"gorm.io/gorm"
)

type DoctorCertification struct {
	gorm.Model
	DoctorProfileID uint   `gorm:"index" json:"doctor_profile_id"`
	CertificateType string `gorm:"size:255"`
	Description     string `gorm:"size:255"`
	StartingDate    time.Time
	EndingDate      time.Time
	FileSize        string `gorm:"size:255"`
	Details         string `gorm:"size:255"`
}
