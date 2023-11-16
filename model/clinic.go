package model

import "gorm.io/gorm"

type Clinic struct {
	gorm.Model
	Name         string         `gorm:"size:255"`
	City         string         `gorm:"size:255"`
	Location     string         `gorm:"size:255"`
	Profile      string         `gorm:"size:255"`
	Latitude     string         `gorm:"size:255"`
	Longitude    string         `gorm:"size:255"`
	Doctors      []Doctor       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Consultation []Consultation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
