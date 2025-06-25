package config

import (
	"log"

	"github.com/streadway/amqp"
)

var RabbitConn *amqp.Connection

func InitRabbitMQ(uri string) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Fatalf("❌ Failed to connect to RabbitMQ: %v", err)
	}
	RabbitConn = conn
	log.Println("✅ Connected to RabbitMQ")
}
