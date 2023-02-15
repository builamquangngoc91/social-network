package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"social-network/internal/entities"
	"social-network/utils/golibs/database"
	"strings"
)

type FollowerRepository struct {
	*sql.DB
}

func NewFollowerRepository(db *sql.DB) *FollowerRepository {
	return &FollowerRepository{
		db,
	}
}

func (r *FollowerRepository) Upsert(ctx context.Context, u *entities.Follower) error {
	fields, values := u.FieldMap()
	placeHolders := database.GeneratePlaceholders(len(fields))

	stmt := fmt.Sprintf(`INSERT INTO %s(%s) 
	VALUES (%s) 
	ON CONFLICT (account_id, follower_id)
	DO
		UPDATE SET followed_date = EXCLUDED.followed_date,
			unfollowed_date = EXCLUDED.unfollowed_date,
			updated_at = EXCLUDED.updated_at`, u.TableName(), strings.Join(fields, ","), placeHolders)
	result, err := r.DB.ExecContext(ctx, stmt, values...)
	if err != nil {
		return fmt.Errorf("r.DB.ExecContext: %w", err)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("result.RowsAffected: %w", err)
	}

	if rowAffected != 1 {
		return fmt.Errorf("can't upsert account")
	}

	return err
}
