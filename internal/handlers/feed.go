package handlers

import (
	"context"
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
	feed := &entities.Feed{}
	if err := feed.Unmarshal(msg); err != nil {
		return fmt.Errorf("feed.Unmarshal: %w", err)
	}

	if err := h.esClient.Index(ctx, "feed", feed); err != nil {
		return fmt.Errorf("h.esClient.Index: %w", err)
	}

	msgNotification := &sse.MessageNotification{
		AccountID: feed.AccountID,
		Message:   fmt.Sprintf("Account (%s) created a feed with message(%s)", feed.FeedID, feed.Message),
	}
	msgNotificationBytes, err := msgNotification.Marshal()
	if err != nil {
		return fmt.Errorf("msgNotification.Marshal: %v", err)
	}

	// notify message
	h.sse.Notifier <- msgNotificationBytes

	return nil
}
