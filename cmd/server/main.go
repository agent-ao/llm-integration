package main

import (
	"log"

	"github.com/agent-ao/llm-integration/internal/config"
	"github.com/agent-ao/llm-integration/internal/handler"
	"github.com/agent-ao/llm-integration/internal/repository"
	"github.com/agent-ao/llm-integration/internal/router"
	"github.com/agent-ao/llm-integration/internal/service"
	"github.com/agent-ao/llm-integration/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()
	logger.Init()

	client := config.InitMongo(cfg.MongoURI)
	config.InitRabbitMQ(cfg.RabbitURI)

	// Initialize repositories
	repos := repository.NewRepos(client.Database(cfg.MongoDBName))
	services := service.NewServices(repos)
	handlers := handler.NewHandlers(services)

	r := router.Setup(handlers)

	log.Printf("✅ Starting server on port %s...", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
