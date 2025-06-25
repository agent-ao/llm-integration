package handler

import (
	"github.com/agent-ao/llm-integration/internal/service"
)

type Handlers struct {
	HealthHandler *HealthHandler
	// Add more handlers here
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		HealthHandler: NewHealthHandler(services.Health),
	}
}
