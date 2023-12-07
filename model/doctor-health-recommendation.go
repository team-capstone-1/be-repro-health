package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DoctorHealthRecommendation struct {
	gorm.Model
	ID        uuid.UUID `json:"id" form:"id"`
	DoctorID  uuid.UUID `gorm:"index" json:"doctor_id"`
	Doctor    Doctor    `gorm:"foreignKey:DoctorID;references:ID"`
	SessionID uuid.UUID `json:"session_id" form:"session_id"`
	Question  string    `gorm:"size:255"`
	Answer    string    `gorm:"size:2000"`
}
