package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"social-network/internal/entities"
	"social-network/utils/golibs/database"
)

type CommentRepository struct {
	*sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{
		db,
	}
}

func (r *CommentRepository) Create(ctx context.Context, u *entities.Comment) error {
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
		return fmt.Errorf("can't insert account")
	}

	return err
}

func (r *CommentRepository) Update(ctx context.Context, u *entities.Comment) error {
	stmt := fmt.Sprintf(`UPDATE %s SET message = $1, image_url = $2, updated_at = NOW() WHERE comment_id = $3`, u.TableName())
	result, err := r.DB.ExecContext(ctx, stmt, &u.Message, &u.ImageUrl, &u.CommentID)
	if err != nil {
		return fmt.Errorf("r.DB.ExecContext: %w", err)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("result.RowsAffected: %w", err)
	}

	if rowAffected != 1 {
		return fmt.Errorf("can't update comment")
	}

	return err
}

// FindByCommentID find feed by commentID
func (r *CommentRepository) FindByCommentID(ctx context.Context, commentID string) (*entities.Comment, error) {
	comment := &entities.Comment{}
	fields, values := comment.FieldMap()

	stmt := fmt.Sprintf(`SELECT %s FROM %s WHERE feed_id = $1`, strings.Join(fields, ","), comment.TableName())
	row := r.QueryRowContext(ctx, stmt, commentID)

	if err := row.Scan(values...); err != nil {
		return nil, err
	}

	return comment, nil
}

type ListCommentsArgs struct {
	FeedID string
}

// List comments
func (r *CommentRepository) List(ctx context.Context, args *ListCommentsArgs) (cs entities.Comments, _ error) {
	comment := &entities.Comment{}
	fields, _ := comment.FieldMap()

	stmt := fmt.Sprintf(`SELECT %s 
	FROM %s 
	WHERE feed_id = $1::TEXT`, strings.Join(fields, ","), comment.TableName())
	rows, err := r.QueryContext(ctx, stmt, args.FeedID)
	if err != nil {
		return nil, fmt.Errorf("r.QueryContext: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		c := &entities.Comment{}
		_, values := c.FieldMap()
		err := rows.Scan(values...)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		cs = append(cs, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return cs, nil
}

func (r *CommentRepository) Delete(ctx context.Context, commentID, accountID string) error {
	comment := &entities.Comment{}

	stmt := fmt.Sprintf(`UPDATE %s SET deleted_at = NOW() WHERE comment_id = $1 AND account_id = $2`, comment.TableName())
	result, err := r.DB.ExecContext(ctx, stmt, &commentID, &accountID)
	if err != nil {
		return fmt.Errorf("r.DB.ExecContext: %w", err)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("result.RowsAffected: %w", err)
	}

	if rowAffected != 1 {
		return fmt.Errorf("can't delete comment")
	}

	return err
}
