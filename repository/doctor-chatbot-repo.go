package repository

import (
	"context"
	"fmt"
	"os"

	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

type DoctorAIRepository interface {
	GetPreviousQuestion(ctx context.Context, sessionID uuid.UUID) (string, error)
	DoctorGetHealthRecommendationWithContext(ctx context.Context, currentQuestion, previousQuestion, message, language string) (string, error)
	DoctorGetHealthRecommendation(ctx context.Context, message, language string) (string, error)
	DoctorStoreChatToDB(data model.DoctorHealthRecommendation)
	DoctorGetAllHealthRecommendations(doctorID uuid.UUID) ([]model.DoctorHealthRecommendation, error)
	DoctorGetAllHealthRecommendationsBySession(sessionID uuid.UUID) ([]model.DoctorHealthRecommendation, error)
	DoctorGetAllHealthRecommendationsByDoctorID(doctorID uuid.UUID) ([]model.DoctorHealthRecommendation, error)
	GetSessionIDFromDatabase(ctx context.Context, doctorID uuid.UUID) (uuid.UUID, error)
	UpdateSessionIDToDatabase(ctx context.Context, doctorID, sessionID uuid.UUID) error
}

type DoctorAIRepositoryImpl struct{}

func NewDoctorAIRepository() DoctorAIRepository {
	return &DoctorAIRepositoryImpl{}
}

func (r *DoctorAIRepositoryImpl) GetPreviousQuestion(ctx context.Context, sessionID uuid.UUID) (string, error) {
	var doctorConversation model.DoctorHealthRecommendation

	if err := database.DB.Where("session_id = ?", sessionID).Last(&doctorConversation).Error; err != nil {
		return "", err
	}

	if doctorConversation.Question == "" {
		return "Saya Emilia tidak bisa menjawab seputar hal diluar kesehatan reproduksi. Apakah ada pertanyaan lain yang berkaitan dengan kesehatan reproduksi? Atau mungkin anda bisa memakai kalimat dengan satu atau lebih kata kunci yang membuat saya bisa memahami pertanyaan anda hehehe...", nil
	}

	return doctorConversation.Question, nil
}

func (r *DoctorAIRepositoryImpl) GetSessionIDFromDatabase(ctx context.Context, doctorID uuid.UUID) (uuid.UUID, error) {
	var doctorConversation model.DoctorHealthRecommendation

	if err := database.DB.Where("doctor_id = ?", doctorID).Order("created_at DESC").First(&doctorConversation).Error; err != nil {
		return uuid.Nil, err
	}

	if doctorConversation.SessionID == uuid.Nil {
		// Create a new session ID
		newSessionID := uuid.New()
		if err := r.UpdateSessionIDToDatabase(ctx, doctorID, newSessionID); err != nil {
			return uuid.Nil, err
		}
		return newSessionID, nil
	}

	return doctorConversation.SessionID, nil
}

func (r *DoctorAIRepositoryImpl) UpdateSessionIDToDatabase(ctx context.Context, doctorID, sessionID uuid.UUID) error {
	var doctorConversation model.DoctorHealthRecommendation

	if err := database.DB.Where("doctor_id = ?", doctorID).Last(&doctorConversation).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&doctorConversation).Update("session_id", sessionID).Error; err != nil {
		return err
	}

	return nil
}

func (r *DoctorAIRepositoryImpl) DoctorGetHealthRecommendationWithContext(ctx context.Context, currentQuestion, previousQuestion, message, language string) (string, error) {
	client := openai.NewClient(os.Getenv("REPROHEALTH"))
	model := openai.GPT3Dot5Turbo

	var systemMessage string
	var tipsPrefix string

	if language == "id" {
		systemMessage = "Halo, saya Emilia, Asisten Dokter Anda. Saya akan membantu Anda menemukan rekomendasi kesehatan untuk kesehatan pasien. Bagaimana saya bisa membantu Anda?"
		tipsPrefix = ""
	} else {
		systemMessage = "Hello, I'm Emilia, your Health Assistant. I will assist you in finding health recommendations. How can I help you?"
		tipsPrefix = ""
	}

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemMessage,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: message,
		},
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s", tipsPrefix, resp.Choices[0].Message.Content), nil
}

func (r *DoctorAIRepositoryImpl) DoctorGetHealthRecommendation(ctx context.Context, message, language string) (string, error) {
	client := openai.NewClient(os.Getenv("REPROHEALTH"))
	model := openai.GPT3Dot5Turbo

	var systemMessage string
	var tipsPrefix string

	if language == "id" {
		systemMessage = "Halo, saya Emilia, Asisten Dokter Anda. Saya akan membantu Anda menemukan rekomendasi kesehatan untuk kesehatan pasien. Bagaimana saya bisa membantu Anda?"
		tipsPrefix = "Tips menjaga kesehatan reproduksi: "
	} else {
		systemMessage = "Hello, I'm Emilia, your Health Assistant. I will assist you in finding health recommendations. How can I help you?"
		tipsPrefix = "Tips for maintaining reproductive health: "
	}

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemMessage,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: message,
		},
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s", tipsPrefix, resp.Choices[0].Message.Content), nil
}

func (r *DoctorAIRepositoryImpl) DoctorStoreChatToDB(data model.DoctorHealthRecommendation) {
	database.DB.Save(&data)
}

func (r *DoctorAIRepositoryImpl) DoctorGetAllHealthRecommendations(doctorID uuid.UUID) ([]model.DoctorHealthRecommendation, error) {
	var doctorDataHealthRecommendations []model.DoctorHealthRecommendation

	tx := database.DB.Where("doctor_id = ?", doctorID).Find(&doctorDataHealthRecommendations)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return doctorDataHealthRecommendations, nil
}

func (r *DoctorAIRepositoryImpl) DoctorGetAllHealthRecommendationsBySession(sessionID uuid.UUID) ([]model.DoctorHealthRecommendation, error) {
	var doctorDataHealthRecommendations []model.DoctorHealthRecommendation

	tx := database.DB.Where("session_id = ?", sessionID).Find(&doctorDataHealthRecommendations)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return doctorDataHealthRecommendations, nil
}

func (r *DoctorAIRepositoryImpl) DoctorGetAllHealthRecommendationsByDoctorID(doctorID uuid.UUID) ([]model.DoctorHealthRecommendation, error) {
	var doctorDataHealthRecommendations []model.DoctorHealthRecommendation

	tx := database.DB.Where("doctor_id = ?", doctorID).Find(&doctorDataHealthRecommendations)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return doctorDataHealthRecommendations, nil
}

func GetDoctorByIDForAI(doctorID uuid.UUID) *model.Doctor {
	var doctor model.Doctor
	result := database.DB.Where("id = ?", doctorID).First(&doctor)

	if result.Error == gorm.ErrRecordNotFound {
		// Doctor with the provided ID not found
		return nil
	}

	if result.Error != nil {
		// Handle other errors if needed
		panic(result.Error)
	}

	return &doctor
}
