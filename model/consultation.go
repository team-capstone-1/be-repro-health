package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Consultation struct {
	gorm.Model
	ID              uuid.UUID     `json:"id" form:"id"`
	DoctorID        uuid.UUID     `gorm:"index" json:"doctor_id"`
	PatientID       uuid.UUID     `gorm:"index" json:"patient_id"`
	ClinicID        uuid.UUID     `gorm:"index" json:"clinic_id"`
	Date            time.Time     `gorm:"type:date"`
	Session         string        `gorm:"type:enum('pagi', 'siang', 'malam')"`
	QueueNumber     string        `gorm:"size:255"`
	DoctorAvailable bool          `gorm:"default:true" json:"doctor_available"`
	Rescheduled	    bool          `gorm:"default:false" json:"rescheduled"`
	PaymentMethod   string        `gorm:"type:ENUM('manual_transfer', 'clinic_payment')"`
	Patient         Patient       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Clinic          Clinic        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Doctor          Doctor        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Transaction     []Transaction `gorm:"foreignKey:ConsultationID"`
}
