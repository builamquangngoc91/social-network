package handlers

import (
	"context"
	"encoding/json"
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
	commentMsg, err := UnmarshalComment(msg)
	if err != nil {
		return fmt.Errorf("UnmarshalComment: %w", err)
	}

	if err := h.esClient.Index(ctx, "comment", commentMsg); err != nil {
		return fmt.Errorf("h.esClient.Index: %w", err)
	}

	// notify message
	h.sse.Notifier <- msg

	return nil
}

func MarshalComment(feed *entities.Comment) ([]byte, error) {
	msgByte, err := json.Marshal(feed)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	return msgByte, nil
}

func UnmarshalComment(val []byte) (*entities.Comment, error) {
	var msg entities.Comment
	if err := json.Unmarshal(val, &msg); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return &msg, nil
}
