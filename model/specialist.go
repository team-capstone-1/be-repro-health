package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Specialist struct {
	gorm.Model
	ID      uuid.UUID `json:"id" form:"id"`
	Name    string    `gorm:"size:255"`
	Image   string    `gorm:"size:255"`
	Doctors []Doctor  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
