package kafka

type TopicDef struct {
	Name      TopicName
	Partition int64
}

type TopicName string
