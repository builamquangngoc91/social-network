package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"social-network/internal/entities"
	"social-network/utils/golibs/database"

	"github.com/lib/pq"
)

type AccountRepository struct {
	*sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		db,
	}
}

func (r *AccountRepository) Create(ctx context.Context, u *entities.Account) error {
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

// FindByUsername find user by username
func (r *AccountRepository) FindByUsername(ctx context.Context, username string) (*entities.Account, error) {
	user := &entities.Account{}
	fields, values := user.FieldMap()

	stmt := fmt.Sprintf(`SELECT %s FROM %s WHERE username = $1`, strings.Join(fields, ","), user.TableName())
	row := r.QueryRowContext(ctx, stmt, username)

	if err := row.Scan(values...); err != nil {
		return nil, err
	}

	return user, nil
}

// FindByID find user by id
func (r *AccountRepository) FindByID(ctx context.Context, id string) (*entities.Account, error) {
	user := &entities.Account{}
	fields, values := user.FieldMap()

	stmt := fmt.Sprintf(`SELECT %s FROM %s WHERE account_id = $1`, strings.Join(fields, ","), user.TableName())
	row := r.QueryRowContext(ctx, stmt, id)

	if err := row.Scan(values...); err != nil {
		return nil, err
	}

	return user, nil
}

type ListUsersArgs struct {
	IDs []string
}

// List find accounts
func (r *AccountRepository) List(ctx context.Context, args *ListUsersArgs) (us entities.Accounts, _ error) {
	user := &entities.Account{}
	fields, _ := user.FieldMap()

	stmt := fmt.Sprintf(`SELECT %s FROM %s 
	WHERE id = ANY($1::_TEXT)`, strings.Join(fields, ","), user.TableName())
	rows, err := r.QueryContext(ctx, stmt, pq.StringArray(args.IDs))
	if err != nil {
		return nil, fmt.Errorf("r.QueryContext: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		u := &entities.Account{}
		_, values := u.FieldMap()
		err := rows.Scan(values...)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		us = append(us, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return us, nil
}
