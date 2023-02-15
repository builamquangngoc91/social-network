package main

import (
	"database/sql"
	"net/http"
	"social-network/utils/cmsql"
	"social-network/utils/kafka"
	log "social-network/utils/log"

	"social-network/internal/services"

	"github.com/Shopify/sarama"
	_ "github.com/lib/pq"
)

var l = log.New()

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		l.Panic(err)
	}
	defer func() {
		if err := syncProducer.Close(); err != nil {
			l.Panic(err)
		}
	}()
	kafkaProducer := kafka.NewKafkaProducer(syncProducer)

	cfg := cmsql.ConfigPostgres{
		Protocol: "postgres",
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "postgres",
		Database: "eureka",
	}

	db, err := sql.Open(cfg.ConnectionString())
	if err != nil {
		l.Panicf("error opening db %v", err)
	}

	if err := db.Ping(); err != nil {
		l.Panicf("error when ping: %v", err)
	}

	services := services.NewServices(db, "secret", kafkaProducer)

	l.Printf("HTTP server listening at %v", "8081")

	if err := http.ListenAndServe(":8081", services); err != nil {
		l.Panicf("error when starting server %v", err)
	}
}
