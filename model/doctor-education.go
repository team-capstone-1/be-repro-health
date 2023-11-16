package model

import (
	"time"

	"gorm.io/gorm"
)

type DoctorEducation struct {
	gorm.Model
	DoctorProfileID  uint            `gorm:"index" json:"doctor_profile_id"`
	StartingDate     time.Time       `gorm:"type:date"`
	EndingDate       time.Time       `gorm:"type:date"`
	EducationProgram string          `gorm:"size:255"`
	University       string          `gorm:"size:255"`
}
