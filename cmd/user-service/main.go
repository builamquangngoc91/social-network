package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"social-network/cmd/user-service/server"
	"social-network/utils/cmsql"
	"social-network/utils/elasticsearch"
	"social-network/utils/kafka"
	log "social-network/utils/log"

	"github.com/Shopify/sarama"
	_ "github.com/lib/pq"
)

var l = log.New()

func main() {
	config := LoadUserServiceConfig()

	// create KafkaProducer
	syncProducer, err := sarama.NewSyncProducer(config.KafkaBrokers, config.KafkaConfig)
	if err != nil {
		l.Panic(err)
	}
	defer func() {
		if err := syncProducer.Close(); err != nil {
			l.Panic(err)
		}
	}()
	kafkaProducer := kafka.NewKafkaProducer(syncProducer)

	// create bobDB
	bobDB, err := sql.Open(config.BobConfigPostgres.ConnectionString())
	if err != nil {
		l.Panicf("error opening db %v", err)
	}
	if err := bobDB.Ping(); err != nil {
		l.Panicf("error when ping: %v", err)
	}

	// create elasticSearchClient
	elasticClient, err := elasticsearch.NewElasticClient(config.ElasticUrls)
	if err != nil {
		l.Info(err.Error())
	}

	userServer := server.NewUserServer(bobDB, "secret", kafkaProducer, elasticClient)
	l.Printf("HTTP server listening at %d", config.Port)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.Host, config.Port), userServer); err != nil {
		l.Panicf("error when starting server %v", err)
	}
}

type UserServiceConfig struct {
	KafkaConfig  *sarama.Config
	KafkaBrokers []string

	BobConfigPostgres cmsql.ConfigPostgres

	ElasticUrls []string

	Host string
	Port int
}

func LoadUserServiceConfig() *UserServiceConfig {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Retry.Max = 5
	kafkaConfig.Producer.Return.Successes = true

	return &UserServiceConfig{
		KafkaConfig:  kafkaConfig,
		KafkaBrokers: []string{"localhost:9092"},
		BobConfigPostgres: cmsql.ConfigPostgres{
			Protocol: "postgres",
			Host:     "localhost",
			Port:     5432,
			Username: "postgres",
			Password: "postgres",
			Database: "bob",
		},
		ElasticUrls: []string{"http://localhost:9200"},
		Port:        8080,
		Host:        "",
	}
}
