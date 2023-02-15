package entities

import "time"

type Comment struct {
	CommentID string     `json:"comment_id"`
	FeedID    string     `json:"feed_id"`
	AccountID string     `json:"account_id"`
	Message   string     `json:"message"`
	ImageUrl  string     `json:"image_url"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
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
