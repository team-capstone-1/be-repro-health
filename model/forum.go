package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Forums struct {
	gorm.Model
	ID         uuid.UUID  `json:"id" form:"id"`
	title      string     `gorm:"size:255"`
	content    string     `gorm:"size:255"`
	anonymous  bool       `gorm:"default:false"`
	date       time.Time  `gorm:"type:datetime"`
	ForumReply ForumReply `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
