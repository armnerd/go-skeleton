package kafka

import (
	"context"
	"fmt"
	"github.com/armnerd/go-skeleton/config"
	"time"

	"github.com/segmentio/kafka-go"
)

func Produce(topic string, key string, value string) {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  config.KafkaConfig.Broker,
		Topic:    topic,
		Balancer: &kafka.Hash{},
		Dialer:   dialer,
	})

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		},
	)
	if err != nil {
		fmt.Printf("failed to write messages: %v\n", err)
	}

	if err := w.Close(); err != nil {
		fmt.Printf("failed to close writer: %v\n", err)
	}
}
