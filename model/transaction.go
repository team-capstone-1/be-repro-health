package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TransactionStatus string

const (
	Processed TransactionStatus = "processed"
	Done      TransactionStatus = "done"
	Cancelled TransactionStatus = "cancelled"
)

type Transaction struct {
	gorm.Model
	ID             uuid.UUID         `json:"id" form:"id"`
	ConsultationID uuid.UUID         `gorm:"index" json:"consultation_id"`
	Date		   time.Time		 `gorm:"type:datetime"`
	Invoice        string            `gorm:"size:255"`
	Price          float64           `gorm:"type:decimal(15,2)"`
	AdminPrice     float64           `gorm:"type:decimal(15,2)"`
	Total          float64           `gorm:"type:decimal(15,2)"`
	Status         TransactionStatus `gorm:"type:ENUM('processed', 'done', 'cancelled')"`
	PaymentStatus  TransactionStatus `gorm:"type:ENUM('processed', 'done', 'cancelled')"`
	Refunds        Refunds           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaymentMethods PaymentMethods    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Consultation   Consultation      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
