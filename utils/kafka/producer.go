package kafka

import (
	"context"

	"github.com/Shopify/sarama"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer(producer sarama.SyncProducer) *KafkaProducer {
	return &KafkaProducer{
		producer: producer,
	}
}

func (p *KafkaProducer) SendMessage(ctx context.Context, topic string, message []byte) error {
	partition, offset, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	})

	l.Infof("Message is stored at topic(%s)/ partition(%d)/ offset(%d)", topic, partition, offset)

	if err != nil {
		l.Errorf("Error when sending message error: %v", err.Error())
	}

	return err
}
