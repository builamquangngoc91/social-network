package main

import (
	"context"
	"fmt"
	"math/rand"
	"social-network/utils/kafka"

	log "social-network/utils/log"

	"github.com/Shopify/sarama"
)

var l = log.New()

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	ctx := context.Background()

	syncProducer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		l.Panic(err)
	}
	defer func() {
		if err := syncProducer.Close(); err != nil {
			l.Panic(err)
		}
	}()

	producer := kafka.NewKafkaProducer(syncProducer)
	for i := 0; i < 100; i++ {
		if err := producer.SendMessage(ctx, "event", []byte(fmt.Sprintf("abc %d", rand.Int31()))); err != nil {
			l.Errorf(err.Error())
		}
	}
}
