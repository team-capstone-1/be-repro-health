package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Clinic struct {
	gorm.Model
	ID           uuid.UUID      `json:"id" form:"id"`
	Name         string         `gorm:"size:255"`
	Image   	 string    		`gorm:"size:255"`
	City         string         `gorm:"size:255"`
	Location     string         `gorm:"size:255"`
	Telephone    string         `gorm:"size:255"`
	Email        string         `gorm:"size:255"`
	Profile      string         `gorm:"size:255"`
	Latitude     string         `gorm:"size:255"`
	Longitude    string         `gorm:"size:255"`
	Doctors      []Doctor       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Consultation []Consultation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
