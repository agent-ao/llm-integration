package provider

import "github.com/agent-ao/llm-integration/internal/domain/queue"

type LLMClient interface {
	GenerateResponse(prompt *queue.Prompt) (string, error)
}
