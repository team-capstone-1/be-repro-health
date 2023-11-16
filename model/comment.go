package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ArticleID uint      `gorm:"index" json:"article_id"`
	PatientID uint      `gorm:"index" json:"patient_id"` 
	Comment   string    `gorm:"size:255"`
	Date      time.Time `gorm:"type:datetime"`
}
