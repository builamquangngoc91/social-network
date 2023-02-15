package up

import (
	"strings"
	"time"

	"social-network/utils/xerror"
)

type Comment struct {
	CommentID string    `json:"comment_id"`
	FeedID    string    `json:"feed_id"`
	AccountID string    `json:"account_id"`
	Message   string    `json:"message"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCommentRequest struct {
	FeedID   string `json:"feed_id"`
	Message  string `json:"message"`
	ImageUrl string `json:"image_url"`
}

func (r *CreateCommentRequest) Validate() error {
	if strings.TrimSpace(r.FeedID) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "feed_id can't be null")
	}
	if strings.TrimSpace(r.Message) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "message can't be null")
	}

	return nil
}

type CreateCommentResponse struct{}

type UpdateCommentRequest struct {
	CommentID string `json:"comment_id"`
	Message   string `json:"message"`
	ImageUrl  string `json:"image_url"`
}

func (r *UpdateCommentRequest) Validate() error {
	if strings.TrimSpace(r.CommentID) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "comment_id can't be null")
	}
	if strings.TrimSpace(r.Message) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "message can't be null")
	}

	return nil
}

type UpdateCommentResponse struct{}

type DeleteCommentRequest struct {
	CommentID string `json:"comment_id"`
}

type DeleteCommentResponse struct{}

func (r *DeleteCommentRequest) Validate() error {
	if strings.TrimSpace(r.CommentID) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "comment_id can't be null")
	}

	return nil
}

type ListCommentsRequest struct {
	FeedID string `json:"feed_id"`
}

type ListCommentsResponse struct {
	Comments []*Comment `json:"comments"`
}
