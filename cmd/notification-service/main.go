package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"social-network/internal/handlers"
	"social-network/utils/kafka"

	"github.com/Shopify/sarama"

	kafka_config "social-network/config/kafka"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := []string{"localhost:9092"}

	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			log.Panic(err)
		}
	}()

	consumer := kafka.NewConsumer(master)
	consumer.StartConsuming(context.Background(), kafka_config.GetTopicDefs(), getHandlers())
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
	log.Println("Interrupt is detected")
}

func getHandlers() map[string]kafka.Handler {
	return map[string]kafka.Handler{
		"event":   handlers.NewEventHandler(),
		"comment": handlers.NewCommentHandler(),
		"feed":    handlers.NewFeedHandler(),
	}
}
