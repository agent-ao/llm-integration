package queue

import "github.com/agent-ao/llm-integration/internal/domain/queue"

type Subscriber interface {
	Handle(envelop *queue.Envelop) error
	Support(eventType string) bool
}
