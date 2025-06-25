package handler

import (
	"net/http"

	"github.com/agent-ao/llm-integration/internal/service"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	service *service.HealthService
}

func NewHealthHandler(service *service.HealthService) *HealthHandler {
	return &HealthHandler{service: service}
}

func (h *HealthHandler) Check(c *gin.Context) {
	status := h.service.GetHealthStatus()
	c.JSON(http.StatusOK, status)
}
