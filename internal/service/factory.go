package service

import "github.com/agent-ao/llm-integration/internal/repository"

type Services struct {
	Health *HealthService
	// Add more services here
}

func NewServices(repos *repository.Repos) *Services {
	return &Services{
		Health: NewHealthService(repos.HealthRepo),
	}
}
