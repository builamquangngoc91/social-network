package main

import (
	"context"
	"os"
	"os/signal"
	"social-network/internal/handlers"
	"social-network/utils/elasticsearch"
	"social-network/utils/kafka"

	"github.com/Shopify/sarama"

	kafka_config "social-network/config/kafka"
	log "social-network/utils/log"
)

var l = log.New()

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := []string{"localhost:9092"}

	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		l.Panic(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			l.Panic(err)
		}
	}()

	elasticClient, err := elasticsearch.NewElasticClient([]string{"http://localhost:9200"})
	if err != nil {
		l.Info(err.Error())
	}

	consumer := kafka.NewConsumer(master)
	consumer.StartConsuming(context.Background(), kafka_config.GetTopicDefs(), getHandlers(elasticClient))

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
	l.Println("Interrupt is detected")
}

func getHandlers(esClient *elasticsearch.ElasticClient) map[string]kafka.Handler {
	return map[string]kafka.Handler{
		"event":   handlers.NewEventHandler(),
		"comment": handlers.NewCommentHandler(),
		"feed":    handlers.NewFeedHandler(esClient),
	}
}
