package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	ID          uuid.UUID   `json:"id" form:"id"`
	PatientID   uuid.UUID   `gorm:"index" json:"patient_id"`
	Title       string      `gorm:"size:255"`
	Content     string      `gorm:"size:255"`
	Category    string      `gorm:"type:enum('janji_temu', 'forum', 'info')"`
	Date        time.Time   `gorm:"type:date"`
}
