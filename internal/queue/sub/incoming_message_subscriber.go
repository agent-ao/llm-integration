package subscriber

import (
	"encoding/json"
	"fmt"
	"log"

	domain "github.com/agent-ao/llm-integration/internal/domain/queue"
	queue "github.com/agent-ao/llm-integration/internal/queue/provider"
	"github.com/agent-ao/llm-integration/internal/service/provider"
	"github.com/agent-ao/llm-integration/pkg/commons/enums/event"
	enums "github.com/agent-ao/llm-integration/pkg/commons/enums/event"
	"google.golang.org/genai"
)

type IncomingMessageSubscriber struct {
	RabbitClient queue.QueueClient
	LLMClient    provider.LLMClient
}

func NewInvoiceMessageSubscriber(client provider.LLMClient, rabbitClient queue.QueueClient) *IncomingMessageSubscriber {
	return &IncomingMessageSubscriber{LLMClient: client, RabbitClient: rabbitClient}
}

func (s *IncomingMessageSubscriber) Handle(envelop *domain.Envelop) error {
	log.Printf("🔄 Handling %s event: %+v", enums.EventTypeIncomingMessage, envelop)

	if envelop.Type != enums.EventTypeIncomingMessage {
		return fmt.Errorf("unsupported event type: %s", envelop.Type)
	}

	var msg domain.IncomingMessageEvent

	dataBytes, err := json.Marshal(envelop.Data)
	if err != nil {
		return fmt.Errorf("failed to marshal envelop.Data: %w", err)
	}
	// Unmarshal the JSON data into the IncomingMessageEvent struct
	if err := json.Unmarshal(dataBytes, &msg); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}
	response, err := s.LLMClient.GenerateResponse(&domain.Prompt{
		History:     msg.ConversationHistory,
		Instruction: &genai.Content{Parts: msg.SystemInstruction},
	})
	if err != nil {
		return fmt.Errorf("failed to get response from LLM: %w", err)
	}

	log.Printf("✅ Received LLM response for session: %s", response)

	outgoingMsg := domain.OutgoingMessageEvent{
		AuthToken:         msg.AuthToken,
		ClientContactId:   msg.ClientContactId,
		BusinessContactId: msg.BusinessContactId,
		Content: []*genai.Content{
			{
				Parts: []*genai.Part{{Text: response}},
				Role:  string(event.RoleAgent),
			},
		},
	}

	envol := domain.Envelop{
		Type: enums.EventTypeOutgoingMessage,
		Data: outgoingMsg,
	}

	envolBytes, err := json.Marshal(envol)
	if err != nil {
		return fmt.Errorf("failed to marshal outgoing message: %w", err)
	}

	return s.RabbitClient.PublishMessage(event.MessageQueue, envolBytes)
}

func (s *IncomingMessageSubscriber) Support(eventType string) bool {
	return eventType == enums.EventTypeIncomingMessage
}
