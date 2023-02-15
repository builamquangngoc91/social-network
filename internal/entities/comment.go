package entities

import "time"

type Comment struct {
	CommentID string
	FeedID    string
	AccountID string
	Message   string
	ImageUrl  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type Comments []*Comment

func (e *Comment) FieldMap() (fields []string, values []interface{}) {
	return []string{
			"comment_id",
			"feed_id",
			"account_id",
			"message",
			"image_url",
			"created_at",
			"updated_at",
			"deleted_at",
		}, []interface{}{
			&e.CommentID,
			&e.FeedID,
			&e.AccountID,
			&e.Message,
			&e.ImageUrl,
			&e.CreatedAt,
			&e.UpdatedAt,
			&e.DeletedAt,
		}
}

func (e *Comment) TableName() string {
	return "comment"
}
