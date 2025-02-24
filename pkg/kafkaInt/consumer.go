// internal/kafka/consumer.go
package kafkaInt

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func NewConsumer(broker string, topic string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func ConsumeMessages(consumer *kafka.Reader) {
	for {
		message, err := consumer.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error while reading message: %v", err)
			continue
		}
		fmt.Printf("Received message: %s\n", string(message.Value))
	}
}
