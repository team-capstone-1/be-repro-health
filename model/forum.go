package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Forum struct {
	gorm.Model
	ID         uuid.UUID    `json:"id" form:"id"`
	PatientID  uuid.UUID    `gorm:"index" json:"patient_id"`
	Patient    Patient      `gorm:"foreignKey:PatientID" json:"patient"`
	Title      string       `gorm:"size:255"`
	Content    string       `gorm:"size:255"`
	Anonymous  bool         `gorm:"default:false"`
	Date       time.Time    `gorm:"type:datetime"`
	View       int          `gorm:"type:int;default:0"`
	Status     bool         `json:"status" gorm:"default:false"`
	ForumReply []ForumReply `gorm:"foreignKey:ForumsID" json:"forum_reply"`
}
