package main

import (
	"log"

	"github.com/agent-ao/llm-integration/internal/config"
	"github.com/agent-ao/llm-integration/internal/handler/queue"
	"github.com/agent-ao/llm-integration/internal/queue/provider/rabbitmq"
	"github.com/agent-ao/llm-integration/internal/service/provider/gemini"
	"github.com/agent-ao/llm-integration/pkg/commons/enums/event"
	"github.com/agent-ao/llm-integration/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()
	logger.Init()

	rabbitClient, err := rabbitmq.NewRabbitMQClient(cfg.RabbitURI)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ client: %v", err)
	}
	defer rabbitClient.Close()

	llmClient := gemini.NewGeminiClient(cfg.LLMAPIKey)

	messageHandler := queue.NewMessageHandler(rabbitClient, llmClient)

	go func() {
		if err := rabbitClient.ConsumeMessages(event.MessageQueue, messageHandler.Handle); err != nil {
			log.Fatalf("Failed to start consuming messages: %v", err)
		}
	}()

	// Block main goroutine until interrupted (e.g., Ctrl+C)
	select {}
}
