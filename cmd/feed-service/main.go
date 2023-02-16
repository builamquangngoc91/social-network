package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"social-network/cmd/feed-service/server"
	"social-network/utils/cmsql"
	"social-network/utils/elasticsearch"
	"social-network/utils/kafka"
	log "social-network/utils/log"

	"github.com/Shopify/sarama"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var l = log.New()

func main() {
	config := LoadFeedServiceConfig()

	// create kafkaProducer
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

	// create eurekaDB
	eurekaDB, err := sql.Open(config.EurekaConfigPostgres.ConnectionString())
	if err != nil {
		l.Panicf("error opening db %v", err)
	}

	if err := eurekaDB.Ping(); err != nil {
		l.Panicf("error when ping: %v", err)
	}

	// create elasticSearchClient
	elasticClient, err := elasticsearch.NewElasticClient(config.ElasticUrls)
	if err != nil {
		l.Info(err.Error())
	}

	// create redisClient
	rd := redis.NewClient(&redis.Options{
		Addr: config.RedisAddr,
	})

	feedServer := server.NewFeedServer(eurekaDB, config.JWTKey, kafkaProducer, elasticClient, rd)
	l.Printf("HTTP server listening at %d", config.Port)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.Host, config.Port), feedServer); err != nil {
		l.Panicf("error when starting server %v", err)
	}
}

type FeedServiceConfig struct {
	KafkaConfig  *sarama.Config
	KafkaBrokers []string

	EurekaConfigPostgres cmsql.ConfigPostgres

	ElasticUrls []string

	RedisAddr string

	Host string
	Port int

	JWTKey string
}

func LoadFeedServiceConfig() *FeedServiceConfig {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Retry.Max = 5
	kafkaConfig.Producer.Return.Successes = true

	return &FeedServiceConfig{
		KafkaConfig:  kafkaConfig,
		KafkaBrokers: []string{"localhost:9092"},
		EurekaConfigPostgres: cmsql.ConfigPostgres{
			Protocol: "postgres",
			Host:     "localhost",
			Port:     5432,
			Username: "postgres",
			Password: "postgres",
			Database: "eureka",
		},
		ElasticUrls: []string{"http://localhost:9200"},
		RedisAddr:   "localhost:6379",
		Port:        8081,
		Host:        "",
		JWTKey:      "secret",
	}
}
