package model

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	DoctorID uint      `gorm:"index" json:"doctor_id"`
	Title    string    `gorm:"size:255"`
	Content  string    `gorm:"size:255"`
	Date     time.Time `gorm:"type:datetime"`
	Image    string    `gorm:"size:255"`
	Comment  []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
