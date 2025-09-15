package queue

import (
	"encoding/json"
	"fmt"
	"log"

	domain "github.com/agent-ao/llm-integration/internal/domain/queue"
	base "github.com/agent-ao/llm-integration/internal/queue"
	queue "github.com/agent-ao/llm-integration/internal/queue/provider"
	sub "github.com/agent-ao/llm-integration/internal/queue/sub"
	"github.com/agent-ao/llm-integration/internal/service/provider"
)

type MessageHandler struct {
	Subscribers []base.Subscriber
}

func NewMessageHandler(queueClient queue.QueueClient, llmClient provider.LLMClient) *MessageHandler {
	return &MessageHandler{
		Subscribers: []base.Subscriber{
			&sub.IncomingMessageSubscriber{
				RabbitClient: queueClient,
				LLMClient:    llmClient,
			},
			// Add more subscribers here
		},
	}
}

func (h *MessageHandler) Handle(body []byte) error {
	var msg domain.Envelop
	if err := json.Unmarshal(body, &msg); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	for _, subscriber := range h.Subscribers {
		if subscriber.Support(msg.Type) {
			if err := subscriber.Handle(&msg); err != nil {
				log.Printf("failed to handle message: %v", err)
			}
		}
	}

	return nil
}
