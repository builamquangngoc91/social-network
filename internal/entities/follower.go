package entities

import "time"

type Follower struct {
	ID             string
	AccountID      string
	FollowerID     string
	FollowedDate   *time.Time
	UnFollowedDate *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

type Followers []*Follower

func (e *Follower) FieldMap() (fields []string, values []interface{}) {
	return []string{
			"id",
			"account_id",
			"follower_id",
			"followed_date",
			"unfollowed_date",
			"created_at",
			"updated_at",
		}, []interface{}{
			&e.ID,
			&e.AccountID,
			&e.FollowerID,
			&e.FollowedDate,
			&e.UnFollowedDate,
			&e.CreatedAt,
			&e.UpdatedAt,
		}
}

func (e *Follower) TableName() string {
	return "follower"
}
