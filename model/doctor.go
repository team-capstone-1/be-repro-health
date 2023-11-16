package model

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	Name            string         `gorm:"size:255"`
	Email           string         `gorm:"size:255"`
	Password        string         `gorm:"size:255"`
	Price           float64        `gorm:"type:decimal(15,2)"`
	SpecialistID    uint           `gorm:"column:specialist_id;index" json:"specialist_id"`
	ClinicID        uint           `gorm:"column:clinic_id;index" json:"clinic_id"`
	DoctorProfileID uint           `gorm:"column:profile_id;index" json:"profile_id"`
	Specialist      Specialist     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Clinic          Clinic         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Consultation    []Consultation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Article         []Article      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
