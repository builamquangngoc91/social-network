// +build !generator

// Code generated by generator sqlgen. DO NOT EDIT.

package model

import (
	"database/sql"
	"sync"
	time "time"

	cmsql "o.o/backend/pkg/common/sql/cmsql"
	migration "o.o/backend/pkg/common/sql/migration"
	core "o.o/backend/pkg/common/sql/sq/core"
)

var __sqlModels []interface{ SQLVerifySchema(db *cmsql.Database) }
var __sqlonce sync.Once

func SQLVerifySchema(db *cmsql.Database) {
	__sqlonce.Do(func() {
		for _, m := range __sqlModels {
			m.SQLVerifySchema(db)
		}
	})
}

type SQLWriter = core.SQLWriter

type Transactions []*Transaction

const __sqlTransaction_Table = "transaction"
const __sqlTransaction_ListCols = "\"name\",\"id\",\"amount\",\"account_id\",\"status\",\"type\",\"classify\",\"note\",\"referral_type\",\"referral_ids\",\"created_at\",\"updated_at\""
const __sqlTransaction_ListColsOnConflict = "\"name\" = EXCLUDED.\"name\",\"id\" = EXCLUDED.\"id\",\"amount\" = EXCLUDED.\"amount\",\"account_id\" = EXCLUDED.\"account_id\",\"status\" = EXCLUDED.\"status\",\"type\" = EXCLUDED.\"type\",\"classify\" = EXCLUDED.\"classify\",\"note\" = EXCLUDED.\"note\",\"referral_type\" = EXCLUDED.\"referral_type\",\"referral_ids\" = EXCLUDED.\"referral_ids\",\"created_at\" = EXCLUDED.\"created_at\",\"updated_at\" = EXCLUDED.\"updated_at\""
const __sqlTransaction_Insert = "INSERT INTO \"transaction\" (" + __sqlTransaction_ListCols + ") VALUES"
const __sqlTransaction_Select = "SELECT " + __sqlTransaction_ListCols + " FROM \"transaction\""
const __sqlTransaction_Select_history = "SELECT " + __sqlTransaction_ListCols + " FROM history.\"transaction\""
const __sqlTransaction_UpdateAll = "UPDATE \"transaction\" SET (" + __sqlTransaction_ListCols + ")"
const __sqlTransaction_UpdateOnConflict = " ON CONFLICT ON CONSTRAINT transaction_pkey DO UPDATE SET"

func (m *Transaction) SQLTableName() string  { return "transaction" }
func (m *Transactions) SQLTableName() string { return "transaction" }
func (m *Transaction) SQLListCols() string   { return __sqlTransaction_ListCols }

func (m *Transaction) SQLVerifySchema(db *cmsql.Database) {
	query := "SELECT " + __sqlTransaction_ListCols + " FROM \"transaction\" WHERE false"
	if _, err := db.SQL(query).Exec(); err != nil {
		db.RecordError(err)
	}
}

func (m *Transaction) Migration(db *cmsql.Database) {
	var mDBColumnNameAndType map[string]string
	if val, err := migration.GetColumnNamesAndTypes(db, "transaction"); err != nil {
		db.RecordError(err)
		return
	} else {
		mDBColumnNameAndType = val
	}
	mModelColumnNameAndType := map[string]migration.ColumnDef{
		"name": {
			ColumnName:       "name",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"id": {
			ColumnName:       "id",
			ColumnType:       "dot.ID",
			ColumnDBType:     "int64",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"amount": {
			ColumnName:       "amount",
			ColumnType:       "int",
			ColumnDBType:     "int",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"account_id": {
			ColumnName:       "account_id",
			ColumnType:       "dot.ID",
			ColumnDBType:     "int64",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"status": {
			ColumnName:       "status",
			ColumnType:       "status3.Status",
			ColumnDBType:     "enum",
			ColumnTag:        "",
			ColumnEnumValues: []string{"Z", "P", "N"},
		},
		"type": {
			ColumnName:       "type",
			ColumnType:       "transaction_type.TransactionType",
			ColumnDBType:     "enum",
			ColumnTag:        "",
			ColumnEnumValues: []string{"credit", "invoice"},
		},
		"classify": {
			ColumnName:       "classify",
			ColumnType:       "service_classify.ServiceClassify",
			ColumnDBType:     "enum",
			ColumnTag:        "",
			ColumnEnumValues: []string{"shipping", "telecom", "all"},
		},
		"note": {
			ColumnName:       "note",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"referral_type": {
			ColumnName:       "referral_type",
			ColumnType:       "subject_referral.SubjectReferral",
			ColumnDBType:     "enum",
			ColumnTag:        "",
			ColumnEnumValues: []string{"unknown", "credit", "invoice", "subscription", "order"},
		},
		"referral_ids": {
			ColumnName:       "referral_ids",
			ColumnType:       "[]dot.ID",
			ColumnDBType:     "[]int64",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"created_at": {
			ColumnName:       "created_at",
			ColumnType:       "time.Time",
			ColumnDBType:     "struct",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"updated_at": {
			ColumnName:       "updated_at",
			ColumnType:       "time.Time",
			ColumnDBType:     "struct",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
	}
	if err := migration.Compare(db, "transaction", mModelColumnNameAndType, mDBColumnNameAndType); err != nil {
		db.RecordError(err)
	}
}

func init() {
	__sqlModels = append(__sqlModels, (*Transaction)(nil))
}

func (m *Transaction) SQLArgs(opts core.Opts, create bool) []interface{} {
	now := time.Now()
	return []interface{}{
		core.String(m.Name),
		m.ID,
		core.Int(m.Amount),
		m.AccountID,
		m.Status,
		m.Type,
		m.Classify,
		core.String(m.Note),
		m.ReferralType,
		core.Array{m.ReferralIDs, opts},
		core.Now(m.CreatedAt, now, create),
		core.Now(m.UpdatedAt, now, true),
	}
}

func (m *Transaction) SQLScanArgs(opts core.Opts) []interface{} {
	return []interface{}{
		(*core.String)(&m.Name),
		&m.ID,
		(*core.Int)(&m.Amount),
		&m.AccountID,
		&m.Status,
		&m.Type,
		&m.Classify,
		(*core.String)(&m.Note),
		&m.ReferralType,
		core.Array{&m.ReferralIDs, opts},
		(*core.Time)(&m.CreatedAt),
		(*core.Time)(&m.UpdatedAt),
	}
}

func (m *Transaction) SQLScan(opts core.Opts, row *sql.Row) error {
	return row.Scan(m.SQLScanArgs(opts)...)
}

func (ms *Transactions) SQLScan(opts core.Opts, rows *sql.Rows) error {
	res := make(Transactions, 0, 128)
	for rows.Next() {
		m := new(Transaction)
		args := m.SQLScanArgs(opts)
		if err := rows.Scan(args...); err != nil {
			return err
		}
		res = append(res, m)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	*ms = res
	return nil
}

func (_ *Transaction) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlTransaction_Select)
	return nil
}

func (_ *Transactions) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlTransaction_Select)
	return nil
}

func (m *Transaction) SQLInsert(w SQLWriter) error {
	w.WriteQueryString(__sqlTransaction_Insert)
	w.WriteRawString(" (")
	w.WriteMarkers(12)
	w.WriteByte(')')
	w.WriteArgs(m.SQLArgs(w.Opts(), true))
	return nil
}

func (ms Transactions) SQLInsert(w SQLWriter) error {
	w.WriteQueryString(__sqlTransaction_Insert)
	w.WriteRawString(" (")
	for i := 0; i < len(ms); i++ {
		w.WriteMarkers(12)
		w.WriteArgs(ms[i].SQLArgs(w.Opts(), true))
		w.WriteRawString("),(")
	}
	w.TrimLast(2)
	return nil
}

func (m *Transaction) SQLUpsert(w SQLWriter) error {
	m.SQLInsert(w)
	w.WriteQueryString(__sqlTransaction_UpdateOnConflict)
	w.WriteQueryString(" ")
	w.WriteQueryString(__sqlTransaction_ListColsOnConflict)
	return nil
}

func (ms Transactions) SQLUpsert(w SQLWriter) error {
	ms.SQLInsert(w)
	w.WriteQueryString(__sqlTransaction_UpdateOnConflict)
	w.WriteQueryString(" ")
	w.WriteQueryString(__sqlTransaction_ListColsOnConflict)
	return nil
}

func (m *Transaction) SQLUpdate(w SQLWriter) error {
	now, opts := time.Now(), w.Opts()
	_, _ = now, opts // suppress unuse error
	var flag bool
	w.WriteRawString("UPDATE ")
	w.WriteName("transaction")
	w.WriteRawString(" SET ")
	if m.Name != "" {
		flag = true
		w.WriteName("name")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Name)
	}
	if m.ID != 0 {
		flag = true
		w.WriteName("id")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.ID)
	}
	if m.Amount != 0 {
		flag = true
		w.WriteName("amount")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Amount)
	}
	if m.AccountID != 0 {
		flag = true
		w.WriteName("account_id")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.AccountID)
	}
	if m.Status != 0 {
		flag = true
		w.WriteName("status")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Status)
	}
	if m.Type != 0 {
		flag = true
		w.WriteName("type")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Type)
	}
	if m.Classify != 0 {
		flag = true
		w.WriteName("classify")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Classify)
	}
	if m.Note != "" {
		flag = true
		w.WriteName("note")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Note)
	}
	if m.ReferralType != 0 {
		flag = true
		w.WriteName("referral_type")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.ReferralType)
	}
	if m.ReferralIDs != nil {
		flag = true
		w.WriteName("referral_ids")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(core.Array{m.ReferralIDs, opts})
	}
	if !m.CreatedAt.IsZero() {
		flag = true
		w.WriteName("created_at")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.CreatedAt)
	}
	if true { // always update time
		flag = true
		w.WriteName("updated_at")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(core.Now(m.UpdatedAt, time.Now(), true))
	}
	if !flag {
		return core.ErrNoColumn
	}
	w.TrimLast(1)
	return nil
}

func (m *Transaction) SQLUpdateAll(w SQLWriter) error {
	w.WriteQueryString(__sqlTransaction_UpdateAll)
	w.WriteRawString(" = (")
	w.WriteMarkers(12)
	w.WriteByte(')')
	w.WriteArgs(m.SQLArgs(w.Opts(), false))
	return nil
}

type TransactionHistory map[string]interface{}
type TransactionHistories []map[string]interface{}

func (m *TransactionHistory) SQLTableName() string  { return "history.\"transaction\"" }
func (m TransactionHistories) SQLTableName() string { return "history.\"transaction\"" }

func (m *TransactionHistory) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlTransaction_Select_history)
	return nil
}

func (m TransactionHistories) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlTransaction_Select_history)
	return nil
}

func (m TransactionHistory) Name() core.Interface         { return core.Interface{m["name"]} }
func (m TransactionHistory) ID() core.Interface           { return core.Interface{m["id"]} }
func (m TransactionHistory) Amount() core.Interface       { return core.Interface{m["amount"]} }
func (m TransactionHistory) AccountID() core.Interface    { return core.Interface{m["account_id"]} }
func (m TransactionHistory) Status() core.Interface       { return core.Interface{m["status"]} }
func (m TransactionHistory) Type() core.Interface         { return core.Interface{m["type"]} }
func (m TransactionHistory) Classify() core.Interface     { return core.Interface{m["classify"]} }
func (m TransactionHistory) Note() core.Interface         { return core.Interface{m["note"]} }
func (m TransactionHistory) ReferralType() core.Interface { return core.Interface{m["referral_type"]} }
func (m TransactionHistory) ReferralIDs() core.Interface  { return core.Interface{m["referral_ids"]} }
func (m TransactionHistory) CreatedAt() core.Interface    { return core.Interface{m["created_at"]} }
func (m TransactionHistory) UpdatedAt() core.Interface    { return core.Interface{m["updated_at"]} }

func (m *TransactionHistory) SQLScan(opts core.Opts, row *sql.Row) error {
	data := make([]interface{}, 12)
	args := make([]interface{}, 12)
	for i := 0; i < 12; i++ {
		args[i] = &data[i]
	}
	if err := row.Scan(args...); err != nil {
		return err
	}
	res := make(TransactionHistory, 12)
	res["name"] = data[0]
	res["id"] = data[1]
	res["amount"] = data[2]
	res["account_id"] = data[3]
	res["status"] = data[4]
	res["type"] = data[5]
	res["classify"] = data[6]
	res["note"] = data[7]
	res["referral_type"] = data[8]
	res["referral_ids"] = data[9]
	res["created_at"] = data[10]
	res["updated_at"] = data[11]
	*m = res
	return nil
}

func (ms *TransactionHistories) SQLScan(opts core.Opts, rows *sql.Rows) error {
	data := make([]interface{}, 12)
	args := make([]interface{}, 12)
	for i := 0; i < 12; i++ {
		args[i] = &data[i]
	}
	res := make(TransactionHistories, 0, 128)
	for rows.Next() {
		if err := rows.Scan(args...); err != nil {
			return err
		}
		m := make(TransactionHistory)
		m["name"] = data[0]
		m["id"] = data[1]
		m["amount"] = data[2]
		m["account_id"] = data[3]
		m["status"] = data[4]
		m["type"] = data[5]
		m["classify"] = data[6]
		m["note"] = data[7]
		m["referral_type"] = data[8]
		m["referral_ids"] = data[9]
		m["created_at"] = data[10]
		m["updated_at"] = data[11]
		res = append(res, m)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	*ms = res
	return nil
}