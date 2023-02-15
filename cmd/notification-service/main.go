package main

import (
	"context"
	"net/http"
	"social-network/internal/handlers"
	"social-network/utils/elasticsearch"
	"social-network/utils/kafka"
	"social-network/utils/sse"

	"github.com/Shopify/sarama"

	kafka_config "social-network/config/kafka"
	log "social-network/utils/log"
)

var l = log.New()

func main() {
	sse := sse.NewServer("secret")

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

	consumer := kafka.NewConsumer(master, sse)
	consumer.StartConsuming(context.Background(), kafka_config.GetTopicDefs(), getHandlers(elasticClient, sse))

	if err := http.ListenAndServe(":8082", sse); err != nil {
		l.Panicf("error when starting server %v", err)
	}
}

func getHandlers(esClient *elasticsearch.ElasticClient, sse *sse.Broker) map[string]kafka.Handler {
	return map[string]kafka.Handler{
		"event":   handlers.NewEventHandler(),
		"comment": handlers.NewCommentHandler(esClient, sse),
		"feed":    handlers.NewFeedHandler(esClient, sse),
	}
}
