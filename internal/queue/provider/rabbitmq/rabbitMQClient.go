package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// RabbitMQClient provides a connection to a RabbitMQ instance.
type RabbitMQClient struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

// NewRabbitMQClient creates and returns a new RabbitMQClient instance.
func NewRabbitMQClient(url string) (*RabbitMQClient, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("❌ failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("❌ failed to open a channel: %w", err)
	}

	log.Println("✅ Connected to RabbitMQ")

	return &RabbitMQClient{conn: conn, ch: ch}, nil
}

// Close closes the RabbitMQ connection and channel.
func (c *RabbitMQClient) Close() {
	c.ch.Close()
	c.conn.Close()
}

// ConsumeMessages starts consuming messages from a queue.
func (c *RabbitMQClient) ConsumeMessages(queueName string, handlerFunc func(body []byte) error) error {
	// Declare the queue to ensure it exists.
	_, err := c.ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	msgs, err := c.ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	// Start a goroutine to process messages.
	go func() {
		for d := range msgs {
			log.Printf("Received a message from %s", queueName)
			err := handlerFunc(d.Body)
			if err != nil {
				log.Printf("Error processing message: %v", err)
			}
		}
	}()

	return nil
}

// PublishMessage publishes a message to a RabbitMQ queue.
func (c *RabbitMQClient) PublishMessage(queueName string, body []byte) error {
	log.Printf("Publishing message to queue %s: %s", queueName, body)
	// Declare the queue to ensure it exists.
	_, err := c.ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	return c.ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
