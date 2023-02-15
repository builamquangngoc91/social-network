package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type EventHandler struct{}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

func (h *EventHandler) Handle(ctx context.Context, msg []byte) error {
	return nil
}

type Event struct {
	EventType EventType   `json:"event_type"`
	Body      interface{} `json:"body"`
	CreatedAt time.Time   `json:"created_at"`
}

type EventType string

const (
	AccountFollowed EventType = "account_followed"
	FeedCreated     EventType = "feed_created"
	CommentCreated  EventType = "comment_created"
)

type AccountFollowedBody struct {
	AccountID  string `json:"account_id"`
	FollowerID string `json:"follower_id"`
}

func NewAccountFollowed(accountID, followerID string) ([]byte, error) {
	event := Event{
		EventType: AccountFollowed,
		Body: &AccountFollowedBody{
			AccountID:  accountID,
			FollowerID: followerID,
		},
		CreatedAt: time.Now(),
	}

	bodyBytes, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	return bodyBytes, nil
}
