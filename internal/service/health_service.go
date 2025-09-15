package service

import (
	"github.com/agent-ao/llm-integration/internal/domain"
)

type HealthService struct {
}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (h *HealthService) GetHealthStatus() domain.HealthStatus {

	return domain.HealthStatus{
		Status: "ok",
		Uptime: "running",
	}
}
