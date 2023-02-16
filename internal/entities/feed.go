package entities

import (
	"encoding/json"
	"fmt"
	"time"
)

type Feed struct {
	FeedID    string     `json:"feed_id"`
	AccountID string     `json:"account_id"`
	Message   string     `json:"message"`
	ImageUrl  string     `json:"image_url"`
	Tag       string     `json:"tag"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Feeds []*Feed

func (e *Feed) FieldMap() (fields []string, values []interface{}) {
	return []string{
			"feed_id",
			"account_id",
			"message",
			"image_url",
			"tag",
			"created_at",
			"updated_at",
			"deleted_at",
		}, []interface{}{
			&e.FeedID,
			&e.AccountID,
			&e.Message,
			&e.ImageUrl,
			&e.Tag,
			&e.CreatedAt,
			&e.UpdatedAt,
			&e.DeletedAt,
		}
}

func (e *Feed) TableName() string {
	return "feed"
}

func (e *Feed) Marshal() ([]byte, error) {
	msgByte, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	return msgByte, nil
}

func (e *Feed) Unmarshal(val []byte) error {
	if err := json.Unmarshal(val, e); err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	return nil
}
