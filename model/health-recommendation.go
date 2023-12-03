package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HealthRecommendation struct {
	gorm.Model
	ID         uuid.UUID    `json:"id" form:"id"`
	PatientID  uuid.UUID    `gorm:"index" json:"patient_id"`
	Question   string       `gorm:"size:255"`
	Answer     string       `gorm:"size:2000"`
}
