package kafka

import (
	"context"
	"sync"

	log "social-network/utils/log"
	"social-network/utils/sse"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	consumer sarama.Consumer
	sse      *sse.Broker
}

func NewConsumer(consumer sarama.Consumer, sse *sse.Broker) *Consumer {
	return &Consumer{
		consumer: consumer,
		sse:      sse,
	}
}

var l = log.New()

func (c *Consumer) StartConsuming(ctx context.Context, topics []*TopicDef, handlers map[string]Handler) error {
	for _, topic := range topics {
		handler := handlers[string(topic.Name)]
		if handler == nil {
			l.Infof("No handler for topic: %s", topic.Name)
			continue
		}

		c.consumeAndHandleTopic(ctx, topic, handler)
	}

	return nil
}

func (c *Consumer) consumeAndHandleTopic(ctx context.Context, topic *TopicDef, handler Handler) {
	var wg sync.WaitGroup

	wg.Add(int(topic.Partition))

	for partition := 0; partition < int(topic.Partition); partition++ {
		go func(topicName string, partition int) {
			defer wg.Done()

			pc, err := c.consumer.ConsumePartition(topicName, int32(partition), sarama.OffsetNewest)
			if err != nil {
				l.Fatalf("Error while consuming topic name(%s)/ partition(%d): %v", topicName, partition, err.Error())
			}

			go func() {
				for {
					select {
					case err := <-pc.Errors():
						l.Errorf("Error while consuming topic name(%s)/ partition(%d): %v", topicName, partition, err.Error())
					case message := <-pc.Messages():
						l.Infof("Received message: key(%s), message(%s) ", message.Key, message.Value)

						// TODO: add code to retry 5 times
						if err := handler.Handle(ctx, message.Value); err != nil {
							l.Errorf("handler.Handle: %w", err)
						}
					}
				}
			}()

		}(string(topic.Name), partition)
	}

	wg.Wait()
}
