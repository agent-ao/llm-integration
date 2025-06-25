package service

import (
	"time"

	"github.com/agent-ao/llm-integration/internal/model"
	queue "github.com/agent-ao/llm-integration/internal/queue/pub"
	"github.com/agent-ao/llm-integration/internal/repository"
)

type HealthService struct {
	repo *repository.HealthRepo
}

func NewHealthService(repo *repository.HealthRepo) *HealthService {
	return &HealthService{repo: repo}
}

func (h *HealthService) GetHealthStatus() model.HealthStatus {
	now := time.Now()
	_ = h.repo.InsertHeartbeat(now) // Log to MongoDB
	_ = queue.PublishHealthCheck()  // Push to RabbitMQ

	return model.HealthStatus{
		Status: "ok",
		Uptime: "running",
	}
}
