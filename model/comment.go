package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        uuid.UUID `json:"id" form:"id"`
	ArticleID uuid.UUID `gorm:"index" json:"article_id"`
	PatientID uuid.UUID `gorm:"index" json:"patient_id"`
	Comment   string    `gorm:"size:255"`
	Profile   string    `gorm:"size:255"`
	Date      time.Time `gorm:"type:datetime"`
}
