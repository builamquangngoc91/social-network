package up

import (
	"strings"
	"time"

	"social-network/utils/xerror"
)

type Feed struct {
	FeedID    string    `json:"feed_id"`
	AccountID string    `json:"account_id"`
	Message   string    `json:"message"`
	ImageUrl  string    `json:"image_url"`
	Tag       string    `json:"tag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateFeedRequest struct {
	Message  string `json:"message"`
	ImageUrl string `json:"image_url"`
	Tag      string `json:"tag"`
}

func (r *CreateFeedRequest) Validate() error {
	if strings.TrimSpace(r.Message) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "message can't be null")
	}

	return nil
}

type CreateFeedResponse struct {
	FeedID string `json:"feed_id"`
}

type UpdateFeedRequest struct {
	FeedID   string `json:"feed_id"`
	Message  string `json:"message"`
	ImageUrl string `json:"image_url"`
}

type UpdateFeedResponse struct{}

func (r *UpdateFeedRequest) Validate() error {
	if strings.TrimSpace(r.FeedID) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "feed_id can't be null")
	}
	if strings.TrimSpace(r.Message) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "message can't be null")
	}

	return nil
}

type DeleteFeedRequest struct {
	FeedID string `json:"feed_id"`
}

type DeleteFeedResponse struct{}

func (r *DeleteFeedRequest) Validate() error {
	if strings.TrimSpace(r.FeedID) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "feed_id can't be null")
	}

	return nil
}

type GetFeedRequest struct {
	FeedID string `json:"feed_id"`
}

func (r *GetFeedRequest) Validate() error {
	if strings.TrimSpace(r.FeedID) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "feed_id can't be null")
	}

	return nil
}

type GetFeedResponse struct {
	*Feed
}

type ListFeedsRequest struct {
	OnlyCurrentAccount bool `json:"only_current_account"`
}

type ListFeedsResponse struct {
	Feeds []*Feed `json:"feeds"`
}

type SearchFeedsRequest struct {
	Field        string        `json:"field"`
	Value        string        `json:"value"`
	OffsetPaging *OffsetPaging `json:"paging"`
}

func (r *SearchFeedsRequest) Validate() error {
	if strings.TrimSpace(r.Field) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "field can't be null")
	}
	if strings.TrimSpace(r.Value) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "value can't be null")
	}

	return nil
}

type SearchFeedsResponse struct {
	Feeds []*Feed `json:"feeds"`
}
