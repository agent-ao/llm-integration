package service

type Services struct {
	Health *HealthService
	// Add more services here
}

func NewServices() *Services {
	return &Services{
		Health: NewHealthService(),
	}
}
