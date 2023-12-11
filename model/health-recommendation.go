package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HealthRecommendation struct {
	gorm.Model
	ID            uuid.UUID `json:"id" form:"id"`
	UserID        uuid.UUID `gorm:"index" json:"user_id"`
	User          User      `gorm:"foreignKey:UserID;references:ID"`
	UserSessionID uuid.UUID `json:"user_session_id" form:"user_session_id"`
	Question      string    `gorm:"size:255"`
	Answer        string    `gorm:"size:2000"`
	IsAIAssistant bool      `gorm:"default:false" json:"is_ai_assistant" form:"is_ai_assistant"`
}
