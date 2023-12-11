package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DoctorHoliday struct {
	gorm.Model
	ID       uuid.UUID `json:"id" form:"id"`
	DoctorID uuid.UUID `gorm:"index" json:"doctor_id"`
	Date     time.Time `gorm:"type:date" json:"date"`
	Session  string    `gorm:"type:enum('pagi', 'siang', 'malam')"`
	Doctor   Doctor    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
