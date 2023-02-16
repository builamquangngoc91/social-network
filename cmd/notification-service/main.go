package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"social-network/internal/handlers"
	"social-network/utils/cmsql"
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

	// create bobDB
	bobDB, err := sql.Open(config.BobConfigPostgres.ConnectionString())
	if err != nil {
		l.Panicf("error opening bobDB %v", err)
	}

	if err := bobDB.Ping(); err != nil {
		l.Panicf("error when ping bobDB: %v", err)
	}

	// create eurekaDB
	eurekaDB, err := sql.Open(config.EurekaConfigPostgres.ConnectionString())
	if err != nil {
		l.Panicf("error opening eurekaBD %v", err)
	}

	if err := eurekaDB.Ping(); err != nil {
		l.Panicf("error when ping eurekaDB: %v", err)
	}

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
	consumer.StartConsuming(context.Background(), kafka_config.GetTopicDefs(), getHandlers(bobDB, eurekaDB, elasticClient, sse))

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.Host, config.Port), sse); err != nil {
		l.Panicf("error when starting server %v", err)
	}
}

func getHandlers(bobDB, eurekaDB *sql.DB, esClient *elasticsearch.ElasticClient, sse *sse.Broker) map[string]kafka.Handler {
	return map[string]kafka.Handler{
		"event":   handlers.NewEventHandler(),
		"comment": handlers.NewCommentHandler(eurekaDB, esClient, sse),
		"feed":    handlers.NewFeedHandler(bobDB, esClient, sse),
	}
}

type NotificationServiceConfig struct {
	KafkaConfig  *sarama.Config
	KafkaBrokers []string

	ElasticUrls []string

	BobConfigPostgres    cmsql.ConfigPostgres
	EurekaConfigPostgres cmsql.ConfigPostgres

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
		BobConfigPostgres: cmsql.ConfigPostgres{
			Protocol: "postgres",
			Host:     "localhost",
			Port:     5432,
			Username: "postgres",
			Password: "postgres",
			Database: "bob",
		},
		EurekaConfigPostgres: cmsql.ConfigPostgres{
			Protocol: "postgres",
			Host:     "localhost",
			Port:     5432,
			Username: "postgres",
			Password: "postgres",
			Database: "eureka",
		},
		Port:   8082,
		Host:   "",
		JWTKey: "secret",
	}
}
