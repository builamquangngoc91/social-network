package kafka

import "social-network/utils/kafka"

func GetTopicDefs() []*kafka.TopicDef {
	return []*kafka.TopicDef{
		{
			Name: "event",
			Partition: 16,
		},
		{
			Name: "feed",
			Partition: 16,
		},
		{
			Name: "comment",
			Partition: 16,
		},
	}
}
