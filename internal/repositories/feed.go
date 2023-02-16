package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"social-network/internal/entities"
	"social-network/utils/golibs/database"
)

type FeedRepository struct {
	*sql.DB
}

func NewFeedRepository(db *sql.DB) *FeedRepository {
	return &FeedRepository{
		db,
	}
}

func (r *FeedRepository) Create(ctx context.Context, u *entities.Feed) error {
	fields, values := u.FieldMap()
	placeHolders := database.GeneratePlaceholders(len(fields))

	stmt := fmt.Sprintf(`INSERT INTO %s(%s) VALUES (%s)`, u.TableName(), strings.Join(fields, ","), placeHolders)
	result, err := r.DB.ExecContext(ctx, stmt, values...)
	if err != nil {
		return fmt.Errorf("r.DB.ExecContext: %w", err)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("result.RowsAffected: %w", err)
	}

	if rowAffected != 1 {
		return fmt.Errorf("can't insert feed")
	}

	return err
}

func (r *FeedRepository) Update(ctx context.Context, u *entities.Feed) error {
	stmt := fmt.Sprintf(`UPDATE %s SET message = $1, image_url = $2, updated_at = NOW() WHERE feed_id = $3 AND account_id = $4`, u.TableName())
	result, err := r.DB.ExecContext(ctx, stmt, &u.Message, &u.ImageUrl, &u.FeedID, &u.AccountID)
	if err != nil {
		return fmt.Errorf("r.DB.ExecContext: %w", err)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("result.RowsAffected: %w", err)
	}

	if rowAffected != 1 {
		return fmt.Errorf("can't update feed")
	}

	return err
}

func (r *FeedRepository) FindByFeedIDAndAccountID(ctx context.Context, feedID, accountID string) (*entities.Feed, error) {
	user := &entities.Feed{}
	fields, values := user.FieldMap()

	stmt := fmt.Sprintf(`SELECT %s FROM %s WHERE feed_id = $1 AND account_id = $2 AND deleted_at IS NULL`, strings.Join(fields, ","), user.TableName())
	row := r.QueryRowContext(ctx, stmt, &feedID, &accountID)

	if err := row.Scan(values...); err != nil {
		return nil, err
	}

	return user, nil
}

type ListFeedsArgs struct {
	AccountID *string
}

func (r *FeedRepository) List(ctx context.Context, args *ListFeedsArgs) (fs entities.Feeds, _ error) {
	feed := &entities.Feed{}
	fields, _ := feed.FieldMap()

	stmt := fmt.Sprintf(`SELECT %s 
	FROM %s 
	WHERE ($1::TEXT IS NULL OR account_id = $1::TEXT) AND deleted_at IS NULL`, strings.Join(fields, ","), feed.TableName())
	rows, err := r.QueryContext(ctx, stmt, &args.AccountID)
	if err != nil {
		return nil, fmt.Errorf("r.QueryContext: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		f := &entities.Feed{}
		_, values := f.FieldMap()
		err := rows.Scan(values...)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		fs = append(fs, f)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return fs, nil
}

func (r *FeedRepository) Search(ctx context.Context, message string) (fs entities.Feeds, _ error) {
	feed := &entities.Feed{}
	fields, _ := feed.FieldMap()

	message = "%" + message + "%"
	stmt := fmt.Sprintf(`SELECT %s 
	FROM %s 
	WHERE message LIKE $1 AND deleted_at IS NULL`, strings.Join(fields, ","), feed.TableName())
	rows, err := r.QueryContext(ctx, stmt, &message)
	if err != nil {
		return nil, fmt.Errorf("r.QueryContext: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		f := &entities.Feed{}
		_, values := f.FieldMap()
		err := rows.Scan(values...)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		fs = append(fs, f)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return fs, nil
}

func (r *FeedRepository) Delete(ctx context.Context, feedID, accountID string) error {
	feed := &entities.Feed{}

	stmt := fmt.Sprintf(`UPDATE %s SET deleted_at = NOW() WHERE feed_id = $1 AND account_id = $2`, feed.TableName())
	result, err := r.DB.ExecContext(ctx, stmt, &feedID, &accountID)
	if err != nil {
		return fmt.Errorf("r.DB.ExecContext: %w", err)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("result.RowsAffected: %w", err)
	}

	if rowAffected != 1 {
		return fmt.Errorf("can't delete feed")
	}

	return err
}
