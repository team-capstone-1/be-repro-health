package model

import "gorm.io/gorm"

type Specialist struct {
	gorm.Model
	Name  string `gorm:"size:255"`
	Image string `gorm:"size:255"`
	Doctors []Doctor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
