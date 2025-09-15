package config

import "os"

type Config struct {
	Port      string
	RabbitURI string
	LLMAPIKey string
}

func Load() *Config {
	return &Config{
		Port:      getEnv("PORT", "8080"),
		RabbitURI: getEnv("RABBITMQ_URI", "amqp://guest:guest@localhost:5672/"),
		LLMAPIKey: getEnv("LLM_API_KEY", ""),
	}
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
