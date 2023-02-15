package entities

import (
	"time"
)

type Feed struct {
	FeedID    string
	AccountID string
	Message   string
	ImageUrl  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type Feeds []*Feed

func (e *Feed) FieldMap() (fields []string, values []interface{}) {
	return []string{
			"feed_id",
			"account_id",
			"message",
			"image_url",
			"created_at",
			"updated_at",
			"deleted_at",
		}, []interface{}{
			&e.FeedID,
			&e.AccountID,
			&e.Message,
			&e.ImageUrl,
			&e.CreatedAt,
			&e.UpdatedAt,
			&e.DeletedAt,
		}
}

func (e *Feed) TableName() string {
	return "feed"
}
