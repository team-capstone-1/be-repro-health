package model

import (
	"time"

	"gorm.io/gorm"
)

type DoctorWorkHistory struct {
	gorm.Model
	DoctorProfileID uint            `gorm:"index" json:"doctor_profile_id"`
	StartingDate    time.Time       `gorm:"type:date"`
	EndingDate      time.Time       `gorm:"type:date"`
	Job             string          `gorm:"size:255"`
	Workplace       string          `gorm:"size:255"`
	Position        string          `gorm:"size:255"`
}
