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

type CommentHandler struct {
	esClient *elasticsearch.ElasticClient
	sse      *sse.Broker
	feedRepo *repositories.FeedRepository
}

func NewCommentHandler(eurekaDB *sql.DB, esClient *elasticsearch.ElasticClient, sse *sse.Broker) *CommentHandler {
	return &CommentHandler{
		esClient: esClient,
		sse:      sse,
		feedRepo: repositories.NewFeedRepository(eurekaDB),
	}
}

func (h *CommentHandler) Handle(ctx context.Context, msg []byte) error {
	comment := &entities.Comment{}
	if err := comment.Unmarshal(msg); err != nil {
		return fmt.Errorf("comment.Unmarshal: %w", err)
	}

	if err := h.esClient.Index(ctx, "comment", comment); err != nil {
		return fmt.Errorf("h.esClient.Index: %w", err)
	}

	feed, err := h.feedRepo.FindByFeedID(ctx, comment.FeedID)
	if err != nil {
		return fmt.Errorf("h.feedRepo.FindByFeedID: %w", err)
	}

	if comment.AccountID != feed.AccountID {
		msgNotification := &sse.MessageNotification{
			AccountID: feed.AccountID,
			Message:   fmt.Sprintf("Account (%s) created a comment with message(%s) on feed(%s)", comment.AccountID, comment.Message, feed.FeedID),
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
