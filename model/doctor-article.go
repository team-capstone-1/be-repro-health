package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID        uuid.UUID `json:"id" form:"id"`
	DoctorID  uuid.UUID `gorm:"column:doctor_id;index" json:"doctor_id"`
	Title     string    `gorm:"size:255"`
	Content   string    `gorm:"size:255"`
	Date      time.Time `gorm:"type:datetime"`
	Image     string    `gorm:"size:255"`
	Published bool      `gorm:"default:false"` // Field untuk status published
	// Comment  []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
