package domain

type HealthStatus struct {
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}
