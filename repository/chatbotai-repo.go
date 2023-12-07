package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
)

type AIRepository interface {
	GetPreviousQuestion(ctx context.Context, UserSessionID uuid.UUID) (string, error)
	PatientGetHealthRecommendationWithContext(ctx context.Context, currentQuestion, previousQuestion, message, language string) (string, error)
	PatientGetHealthRecommendation(ctx context.Context, message, language string) (string, error)
	StoreChatToDB(data model.HealthRecommendation)
	GetAllHealthRecommendations(patient_id uuid.UUID) ([]model.HealthRecommendation, error)
	GetPatientSessionIDFromDatabase(ctx context.Context, PatientID uuid.UUID) (uuid.UUID, error)
}

type PatientAIRepository struct{}

func NewPatientAIRepository() PatientAIRepository {
	return PatientAIRepository{}
}

func (r *PatientAIRepository) GetPreviousQuestion(ctx context.Context, UserSessionID uuid.UUID) (string, error) {
	var patientConversation model.HealthRecommendation

	if err := database.DB.Where("session_id = ?", UserSessionID).Last(&patientConversation).Error; err != nil {
		return "", err
	}

	if patientConversation.Question == "" {
		return "Kami tidak dapat menampilkan hasil yang sesuai pencarian Anda. Silahkan pilih dari kategori berikut atau coba cari kata kunci yang berkaitan dengan kesehatan reproduksi.", nil
	}

	return patientConversation.Question, nil
}

func (r *PatientAIRepository) GetPatientSessionIDFromDatabase(ctx context.Context, PatientID uuid.UUID) (uuid.UUID, error) {
	var patientConversation model.HealthRecommendation

	if err := database.DB.Where("patient_id = ?", PatientID).Order("created_at DESC").First(&patientConversation).Error; err != nil {
		return uuid.Nil, err
	}

	if patientConversation.PatientSessionID == uuid.Nil {
		// Create a new session ID
		newSessionID := uuid.New()
		if err := r.UpdateSessionIDUserToDatabase(ctx, PatientID, newSessionID); err != nil {
			return uuid.Nil, err
		}
		return newSessionID, nil
	}

	return patientConversation.PatientSessionID, nil
}

func (r *PatientAIRepository) UpdateSessionIDUserToDatabase(ctx context.Context, PatientID, PatientSessionID uuid.UUID) error {
	var patientConversation model.HealthRecommendation

	if err := database.DB.Where("patient_id = ?", PatientID).Last(&patientConversation).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&patientConversation).Update("session_id", PatientSessionID).Error; err != nil {
		return err
	}

	return nil
}

func (ar *PatientAIRepository) PatientGetHealthRecommendationWithContext(ctx context.Context, currentQuestion, previousQuestion, message, language string) (string, error) {
	client := openai.NewClient(os.Getenv("REPROHEALTH"))
	model := openai.GPT3Dot5Turbo

	var systemMessage string
	var tipsPrefix string

	if language == "id" {
		systemMessage = "Halo, saya Emilia, Asisten Kesehatan Anda. Saya akan membantu Anda menemukan rekomendasi kesehatan. Bagaimana saya bisa membantu Anda?"
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

func (r *PatientAIRepository) PatientGetHealthRecommendation(ctx context.Context, message, language string) (string, error) {
	client := openai.NewClient(os.Getenv("REPROHEALTH"))
	model := openai.GPT3Dot5Turbo

	var systemMessage string
	var tipsPrefix string

	if language == "id" {
		systemMessage = "Halo, saya Emilia, Asisten Kesehatan Anda. Saya akan membantu Anda menemukan rekomendasi kesehatan. Bagaimana saya bisa membantu Anda?"
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

func (ar *PatientAIRepository) StoreChatToDB(data model.HealthRecommendation) {
	database.DB.Save(&data)
}

func (ar *PatientAIRepository) GetAllHealthRecommendations(patient_id uuid.UUID) ([]model.HealthRecommendation, error) {
	var datahealthRecommendations []model.HealthRecommendation

	tx := database.DB.Where("patient_id = ?", patient_id).Find(&datahealthRecommendations)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datahealthRecommendations, nil
}
