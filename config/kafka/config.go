package kafka

type TopicConfig struct {
	Name      string
	Partition int64
}

func GetKafkaTopicConfigs() []*TopicConfig {
	return []*TopicConfig{
		{
			Name:      "event",
			Partition: 16,
		},
		{
			Name:      "feed",
			Partition: 16,
		},
		{
			Name:      "comment",
			Partition: 16,
		},
	}
}
