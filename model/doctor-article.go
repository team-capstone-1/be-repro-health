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
	Tags      string    `gorm:"size:255"`
	Reference string    `gorm:"size:255"`
	Content   string    `gorm:"size:5000"`
	Image     string    `gorm:"size:255"`
	ImageDesc string    `gorm:"size:255"`
	Date      time.Time `gorm:"type:datetime"`
	Published bool      `gorm:"default:false"`
	View      int       `gorm:"type:int;default:0"`
	Comment   []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
