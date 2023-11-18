package model


import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Consultation struct {
	gorm.Model
	ID          uuid.UUID   `json:"id" form:"id"`
	DoctorID    uuid.UUID   `gorm:"index" json:"doctor_id"`
	// PatientID   uuid.UUID   `gorm:"index" json:"patient_id"`
	ClinicID    uuid.UUID   `gorm:"index" json:"clinic_id"`
	Date        time.Time   `gorm:"type:date"`
	Session     string      `gorm:"type:enum('pagi', 'siang', 'malam')"`
	Clinic      Clinic      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Menambah relasi ke Clinic
	Transaction Transaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
