package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"social-network/internal/entities"
	"social-network/utils/elasticsearch"
	"social-network/utils/sse"
)

type FeedHandler struct {
	esClient *elasticsearch.ElasticClient
	sse      *sse.Broker
}

func NewFeedHandler(esClient *elasticsearch.ElasticClient, sse *sse.Broker) *FeedHandler {
	return &FeedHandler{
		esClient: esClient,
		sse:      sse,
	}
}

func (h *FeedHandler) Handle(ctx context.Context, msg []byte) error {
	feedMsg, err := UnmarshalFeed(msg)
	if err != nil {
		return fmt.Errorf("UnmarshalFeed: %w", err)
	}

	if err := h.esClient.Index(ctx, "feed", feedMsg); err != nil {
		return fmt.Errorf("h.esClient.Index: %w", err)
	}

	// notify message
	h.sse.Notifier <- msg

	return nil
}

func MarshalFeed(feed *entities.Feed) ([]byte, error) {
	msgByte, err := json.Marshal(feed)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	return msgByte, nil
}

func UnmarshalFeed(val []byte) (*entities.Feed, error) {
	var msg entities.Feed
	if err := json.Unmarshal(val, &msg); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return &msg, nil
}
