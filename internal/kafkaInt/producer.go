// internal/kafka/producer.go
package kafkaInt

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func NewProducer(broker string, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func SendMessage(producer *kafka.Writer, message string) error {
	err := producer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: []byte(message),
		},
	)
	if err != nil {
		return fmt.Errorf("could not send message: %w", err)
	}
	log.Println("Message sent to Kafka")
	return nil
}
