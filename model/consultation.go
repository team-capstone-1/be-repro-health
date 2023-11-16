package model

import (
	"time"

	"gorm.io/gorm"
)

type Consultation struct {
	gorm.Model
	DoctorID uint `gorm:"index" json:"doctor_id"`
	PatientID uint      `gorm:"index" json:"patient_id"`
	ClinicID    uint        `gorm:"index" json:"clinic_id"`
	Date        time.Time   `gorm:"type:date"`
	Session     string      `gorm:"type:enum('pagi', 'siang', 'malam')"`
	Clinic      Clinic      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Menambah relasi ke Clinic
	Transaction Transaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
