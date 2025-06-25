package router

import (
	"github.com/agent-ao/llm-integration/internal/handler"
	"github.com/gin-gonic/gin"
)

func Setup(handlers *handler.Handlers) *gin.Engine {
	r := gin.Default()
	r.GET("/health", handlers.HealthHandler.Check)
	return r
}
