package router

import (
	"github.com/agent-ao/llm-integration/internal/handler/rest"
	"github.com/gin-gonic/gin"
)

func Setup(handlers *rest.Handlers) *gin.Engine {
	r := gin.Default()
	r.GET("/health", handlers.HealthHandler.Check)
	return r
}
