package handlers

import (
	"context"
	"database/sql"
	"fmt"

	"social-network/internal/entities"
	"social-network/internal/repositories"
	"social-network/utils/elasticsearch"
	"social-network/utils/sse"
)

type FeedHandler struct {
	esClient     *elasticsearch.ElasticClient
	sse          *sse.Broker
	followerRepo *repositories.FollowerRepository
}

func NewFeedHandler(bobDB *sql.DB, esClient *elasticsearch.ElasticClient, sse *sse.Broker) *FeedHandler {
	return &FeedHandler{
		esClient:     esClient,
		sse:          sse,
		followerRepo: repositories.NewFollowerRepository(bobDB),
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

	followers, err := h.followerRepo.ListByAccountID(ctx, feed.AccountID)
	if err != nil {
		return fmt.Errorf("h.followerRepo.ListByAccountID: %w", err)
	}

	for _, follower := range followers {
		msgNotification := &sse.MessageNotification{
			AccountID: follower.FollowerID,
			Message:   fmt.Sprintf("Account (%s) created a feed with message(%s)", feed.AccountID, feed.Message),
		}
		msgNotificationBytes, err := msgNotification.Marshal()
		if err != nil {
			return fmt.Errorf("msgNotification.Marshal: %v", err)
		}

		// notify message
		h.sse.Notifier <- msgNotificationBytes
	}

	return nil
}
