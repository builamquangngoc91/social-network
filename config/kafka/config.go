package kafka

import "social-network/utils/kafka"

func GetTopicDefs() []*kafka.TopicDef {
	return []*kafka.TopicDef{
		{
			Name:      EventTopic,
			Partition: 16,
		},
		{
			Name:      FeedTopic,
			Partition: 16,
		},
		{
			Name:      CommentTopic,
			Partition: 16,
		},
	}
}

const (
	EventTopic   kafka.TopicName = "event"
	FeedTopic    kafka.TopicName = "feed"
	CommentTopic kafka.TopicName = "comment"
)
