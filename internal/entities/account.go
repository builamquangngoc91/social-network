package entities

import "time"

type Account struct {
	AccountID string
	Username  string
	Password  string
	Fullname  string
	Email     string
	Status    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type Accounts []*Account

func (e *Account) FieldMap() (fields []string, values []interface{}) {
	return []string{
			"account_id",
			"username",
			"password",
			"fullname",
			"email",
			"status",
			"created_at",
			"updated_at",
		}, []interface{}{
			&e.AccountID,
			&e.Username,
			&e.Password,
			&e.Fullname,
			&e.Email,
			&e.Status,
			&e.CreatedAt,
			&e.UpdatedAt,
		}
}

func (e *Account) TableName() string {
	return "account"
}
