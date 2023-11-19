package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DoctorEducation struct {
	gorm.Model
	ID               uuid.UUID `json:"id" form:"id"`
	DoctorProfileID  uuid.UUID `gorm:"index" json:"doctor_profile_id"`
	StartingDate     time.Time `gorm:"type:date"`
	EndingDate       time.Time `gorm:"type:date"`
	EducationProgram string    `gorm:"size:255"`
	University       string    `gorm:"size:255"`
}
