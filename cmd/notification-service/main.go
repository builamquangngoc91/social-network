package main

import (
	"context"
	"fmt"
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
	config := LoadNotificationServiceConfig()

	master, err := sarama.NewConsumer(config.KafkaBrokers, config.KafkaConfig)
	if err != nil {
		l.Panic(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			l.Panic(err)
		}
	}()

	elasticClient, err := elasticsearch.NewElasticClient(config.ElasticUrls)
	if err != nil {
		l.Info(err.Error())
	}

	sse := sse.NewServer(config.JWTKey)

	consumer := kafka.NewConsumer(master, sse)
	consumer.StartConsuming(context.Background(), kafka_config.GetTopicDefs(), getHandlers(elasticClient, sse))

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.Host, config.Port), sse); err != nil {
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

type NotificationServiceConfig struct {
	KafkaConfig  *sarama.Config
	KafkaBrokers []string

	ElasticUrls []string

	Host string
	Port int

	JWTKey string
}

func LoadNotificationServiceConfig() *NotificationServiceConfig {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Retry.Max = 5
	kafkaConfig.Producer.Return.Successes = true

	return &NotificationServiceConfig{
		KafkaConfig:  kafkaConfig,
		KafkaBrokers: []string{"localhost:9092"},
		ElasticUrls:  []string{"http://localhost:9200"},
		Port:         8082,
		Host:         "",
		JWTKey:       "secret",
	}
}
