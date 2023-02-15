package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"social-network/internal/entities"
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

type FeedCreatedBody struct {
	FeedID    string `json:"feed_id"`
	AccountID string `json:"account_id"`
	Message   string `json:"message"`
	ImageUrl  string `json:"image_url"`
}

func NewFeedCreated(feed *entities.Feed) ([]byte, error) {
	event := Event{
		EventType: FeedCreated,
		Body: &FeedCreatedBody{
			FeedID:    feed.FeedID,
			AccountID: feed.AccountID,
			Message:   feed.Message,
			ImageUrl:  feed.ImageUrl,
		},
		CreatedAt: time.Now(),
	}

	bodyBytes, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	return bodyBytes, nil
}
