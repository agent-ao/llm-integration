package gemini

import (
	"context"
	"fmt"
	"log"

	"github.com/agent-ao/llm-integration/internal/domain/queue"
	"google.golang.org/genai"
)

type GeminiClient struct {
	client *genai.Client
	model  string
}

func NewGeminiClient(apiKey string) *GeminiClient {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}
	return &GeminiClient{client: client, model: "gemini-2.5-pro"}
}

func (c *GeminiClient) GenerateResponse(prompt *queue.Prompt) (string, error) {
	ctx := context.Background()

	resp, err := c.client.Models.GenerateContent(ctx, c.model,
		prompt.History, &genai.GenerateContentConfig{
			SystemInstruction: prompt.Instruction,
		})

	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		part := resp.Candidates[0].Content.Parts[0]
		if part != nil && part.Text != "" {
			return part.Text, nil
		}
	}

	return "", fmt.Errorf("no valid response from LLM")
}
