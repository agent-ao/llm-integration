package config

import "os"

type Config struct {
	Port        string
	MongoURI    string
	RabbitURI   string
	MongoDBName string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		MongoURI:    getEnv("MONGO_URI", "mongodb://localhost:27017"),
		RabbitURI:   getEnv("RABBITMQ_URI", "amqp://guest:guest@localhost:5672/"),
		MongoDBName: getEnv("MONGO_DB_NAME", "llm_integration"),
	}
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
