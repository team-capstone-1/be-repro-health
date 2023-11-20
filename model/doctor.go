package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	ID           uuid.UUID      `json:"id" form:"id"`
	Name         string         `gorm:"size:255"`
	Email        string         `gorm:"size:255"`
	Password     string         `gorm:"size:255"`
	Price        float64        `gorm:"type:decimal(15,2)"`
	Address      string         `gorm:"size:255"`
	Phone        string         `gorm:"size:255"`
	ProfileImage string         `json:"profile_image" gorm:"size:255"`
	SpecialistID uuid.UUID      `gorm:"column:specialist_id;index" json:"specialist_id"`
	ClinicID     uuid.UUID      `gorm:"column:clinic_id;index" json:"clinic_id"`
	Specialist   Specialist     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Clinic       Clinic         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Consultation []Consultation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Article      []Article      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
