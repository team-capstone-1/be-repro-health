package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

// AIRepository handles interactions with the AI model.
type AIRepository interface {
	GetHealthRecommendation(ctx context.Context, message, language string) (string, error)
}

type aiRepository struct{}

// NewAIRepository creates a new instance of AIRepository.
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
