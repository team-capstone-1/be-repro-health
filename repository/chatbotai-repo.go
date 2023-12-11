package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
	"github.com/google/uuid"
)

type UserAIRepository interface {
	GetPreviousQuestionUser(ctx context.Context, UserSessionID uuid.UUID) (string, error)
	UserGetHealthRecommendationWithContext(ctx context.Context, currentQuestion, previousQuestion, message, language string) (string, error)
	UserGetHealthRecommendation(ctx context.Context, message, language string) (string, error)
	StoreChatToDB(data model.HealthRecommendation)
	UserGetAllHealthRecommendations(user_id uuid.UUID) ([]model.HealthRecommendation, error)
	GetSessionUserIDFromDatabase(ctx context.Context, userID uuid.UUID) (uuid.UUID, error)
	UpdateSessionUserIDToDatabase(ctx context.Context, userID, UserSessionID uuid.UUID) error
}

type UserAIRepositoryImpl struct{}

func NewUserAIRepository() UserAIRepository {
	return &UserAIRepositoryImpl{}
}

func (ar *UserAIRepositoryImpl) GetPreviousQuestionUser(ctx context.Context, userSessionID uuid.UUID) (string, error) {
	var userConversation model.HealthRecommendation

	if err := database.DB.Where("user_session_id = ?", userSessionID).Last(&userConversation).Error; err != nil {
		return "", err
	}

	if userConversation.Question == "" {
		return "Saya tidak bisa menjawab seputar hal diluar kesehatan reproduksi. Apakah ada pertanyaan lain yang berkaitan dengan kesehatan reproduksi? Atau mungkin anda bisa memakai kalimat dengan satu atau lebih kata kunci yang membuat saya bisa memahami pertanyaan anda hehehe...", nil
	}

	return userConversation.Question, nil
}

func (ar *UserAIRepositoryImpl) GetSessionUserIDFromDatabase(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	var userConversation model.HealthRecommendation

	if err := database.DB.Where("user_id = ?", userID).Order("created_at DESC").First(&userConversation).Error; err != nil {
		return uuid.Nil, err
	}

	if userConversation.UserSessionID == uuid.Nil {
		// Create a new session ID
		newUserSessionID := uuid.New()
		if err := ar.UpdateSessionUserIDToDatabase(ctx, userID, newUserSessionID); err != nil {
			return uuid.Nil, err
		}
		return newUserSessionID, nil
	}
	return userConversation.UserSessionID, nil
}

func (ar *UserAIRepositoryImpl) UpdateSessionUserIDToDatabase(ctx context.Context, userID, UserSessionID uuid.UUID) error {
	var userConversation model.HealthRecommendation

	if err := database.DB.Where("user_id = ?", userID).Last(&userConversation).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&userConversation).Update("user_session_id", UserSessionID).Error; err != nil {
		return err
	}

	return nil
}

func (ar *UserAIRepositoryImpl) UserGetHealthRecommendationWithContext(ctx context.Context, currentQuestion, previousQuestion, message, language string) (string, error) {
	client := openai.NewClient(os.Getenv("REPROHEALTH"))
	model := openai.GPT3Dot5Turbo

	var systemMessage string
	var tipsPrefix string

	if language == "id" {
		systemMessage = "Halo, Saya akan membantu Anda menemukan rekomendasi kesehatan. Bagaimana saya bisa membantu Anda?"
		tipsPrefix = ""
	} else {
		systemMessage = "Hello, I will assist you in finding health recommendations. How can I help you?"
		tipsPrefix = ""
	}

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemMessage,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: currentQuestion,
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

func (ar *UserAIRepositoryImpl) UserGetHealthRecommendation(ctx context.Context, message, language string) (string, error) {
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

func (ar *aiRepository) StoreChatToDB(data model.HealthRecommendation){
	database.DB.Save(&data)
}

func (ar *aiRepository)GetAllHealthRecommendations(patient_id uuid.UUID) ([]model.HealthRecommendation, error) {
	var datahealthRecommendations []model.HealthRecommendation

	tx := database.DB.Where("user_id = ?", UserID).Find(&userDataHealthRecommendations)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return userDataHealthRecommendations, nil
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

