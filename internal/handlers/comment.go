package handlers

import (
	"context"
	"fmt"

	"social-network/internal/entities"
	"social-network/utils/elasticsearch"
	"social-network/utils/sse"
)

type CommentHandler struct {
	esClient *elasticsearch.ElasticClient
	sse      *sse.Broker
}

func NewCommentHandler(esClient *elasticsearch.ElasticClient, sse *sse.Broker) *CommentHandler {
	return &CommentHandler{
		esClient: esClient,
		sse:      sse,
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

	// notify message
	h.sse.Notifier <- msg

	return nil
}
