package config
import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConfig struct {
	URL string
}

func NewRabbitMQConfig() *RabbitMQConfig {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "amqp://guest:guest@localhost:5672/"
	}
	return &RabbitMQConfig{URL: url}
}

func (c *RabbitMQConfig) Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial(c.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	return conn, nil
}
