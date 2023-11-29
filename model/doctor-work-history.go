package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DoctorWorkHistory struct {
	gorm.Model
	ID           uuid.UUID `json:"id" form:"id"`
	DoctorID     uuid.UUID `gorm:"column:doctor_id;index" json:"doctor_id"`
	StartingDate time.Time `gorm:"type:date"`
	EndingDate   time.Time `gorm:"type:date"`
	Job          string    `gorm:"size:255"`
	Workplace    string    `gorm:"size:255"`
	Position     string    `gorm:"size:255"`
}
