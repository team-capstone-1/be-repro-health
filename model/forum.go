package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Forum struct {
	gorm.Model
	ID         uuid.UUID  `json:"id" form:"id"`
	PatientID  uuid.UUID  `gorm:"index" json:"patient_id"`
	Title      string     `gorm:"size:255"`
	Content    string     `gorm:"size:255"`
	Anonymous  bool       `gorm:"default:false"`
	Date       time.Time  `gorm:"type:datetime"`
	// ForumReply ForumReply `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
