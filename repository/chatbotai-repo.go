package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

type AIRepository interface {
	GetHealthRecommendation(ctx context.Context, message, language string) (string, error)
	StoreChatToDB(data model.HealthRecommendation)
	GetAllHealthRecommendations(patient_id uuid.UUID) ([]model.HealthRecommendation, error)
}

type aiRepository struct{}

func NewAIRepository() AIRepository {
	return &aiRepository{}
}

func (ar *aiRepository) GetHealthRecommendation(ctx context.Context, message, language string) (string, error) {
	client := openai.NewClient(os.Getenv("REPROHEALTH"))
	model := openai.GPT3Dot5Turbo

	var systemMessage string
	var tipsPrefix string

	if language == "id" {
		systemMessage = "Halo, saya Emilia, Asisten Kesehatan Anda. Saya akan membantu Anda menemukan rekomendasi kesehatan. Bagaimana saya bisa membantu Anda?"
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

func (ar *aiRepository) StoreChatToDB(data model.HealthRecommendation) {
	database.DB.Save(&data)
}

func (ar *aiRepository) GetAllHealthRecommendations(patient_id uuid.UUID) ([]model.HealthRecommendation, error) {
	var datahealthRecommendations []model.HealthRecommendation

	tx := database.DB.Where("patient_id = ?", patient_id).Find(&datahealthRecommendations)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datahealthRecommendations, nil
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
