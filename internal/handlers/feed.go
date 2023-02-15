package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"social-network/internal/entities"
	"social-network/utils/elasticsearch"
)

type FeedHandler struct {
	esClient *elasticsearch.ElasticClient
}

func NewFeedHandler(esClient *elasticsearch.ElasticClient) *FeedHandler {
	return &FeedHandler{
		esClient: esClient,
	}
}

func (h *FeedHandler) Handle(ctx context.Context, msg []byte) error {
	feedMsg, err := UnmarshalFeedCreated(msg)
	if err != nil {
		return fmt.Errorf("UnmarshalFeedCreated: %w", err)
	}

	if err := h.esClient.Index(ctx, "feed", feedMsg); err != nil {
		return fmt.Errorf("h.esClient.Index: %w", err)
	}

	return nil
}

func MarshalFeedCreated(feed *entities.Feed) ([]byte, error) {
	msgByte, err := json.Marshal(feed)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	return msgByte, nil
}

func UnmarshalFeedCreated(val []byte) (*entities.Feed, error) {
	var msg entities.Feed
	if err := json.Unmarshal(val, &msg); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return &msg, nil
}
