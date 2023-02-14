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

type PurchaseRefunds []*PurchaseRefund

const __sqlPurchaseRefund_Table = "purchase_refund"
const __sqlPurchaseRefund_ListCols = "\"id\",\"shop_id\",\"purchase_order_id\",\"code\",\"code_norm\",\"note\",\"lines\",\"created_at\",\"updated_at\",\"cancelled_at\",\"confirmed_at\",\"created_by\",\"updated_by\",\"cancel_reason\",\"status\",\"supplier_id\",\"total_amount\",\"basket_value\",\"rid\""
const __sqlPurchaseRefund_ListColsOnConflict = "\"id\" = EXCLUDED.\"id\",\"shop_id\" = EXCLUDED.\"shop_id\",\"purchase_order_id\" = EXCLUDED.\"purchase_order_id\",\"code\" = EXCLUDED.\"code\",\"code_norm\" = EXCLUDED.\"code_norm\",\"note\" = EXCLUDED.\"note\",\"lines\" = EXCLUDED.\"lines\",\"created_at\" = EXCLUDED.\"created_at\",\"updated_at\" = EXCLUDED.\"updated_at\",\"cancelled_at\" = EXCLUDED.\"cancelled_at\",\"confirmed_at\" = EXCLUDED.\"confirmed_at\",\"created_by\" = EXCLUDED.\"created_by\",\"updated_by\" = EXCLUDED.\"updated_by\",\"cancel_reason\" = EXCLUDED.\"cancel_reason\",\"status\" = EXCLUDED.\"status\",\"supplier_id\" = EXCLUDED.\"supplier_id\",\"total_amount\" = EXCLUDED.\"total_amount\",\"basket_value\" = EXCLUDED.\"basket_value\",\"rid\" = EXCLUDED.\"rid\""
const __sqlPurchaseRefund_Insert = "INSERT INTO \"purchase_refund\" (" + __sqlPurchaseRefund_ListCols + ") VALUES"
const __sqlPurchaseRefund_Select = "SELECT " + __sqlPurchaseRefund_ListCols + " FROM \"purchase_refund\""
const __sqlPurchaseRefund_Select_history = "SELECT " + __sqlPurchaseRefund_ListCols + " FROM history.\"purchase_refund\""
const __sqlPurchaseRefund_UpdateAll = "UPDATE \"purchase_refund\" SET (" + __sqlPurchaseRefund_ListCols + ")"
const __sqlPurchaseRefund_UpdateOnConflict = " ON CONFLICT ON CONSTRAINT purchase_refund_pkey DO UPDATE SET"

func (m *PurchaseRefund) SQLTableName() string  { return "purchase_refund" }
func (m *PurchaseRefunds) SQLTableName() string { return "purchase_refund" }
func (m *PurchaseRefund) SQLListCols() string   { return __sqlPurchaseRefund_ListCols }

func (m *PurchaseRefund) SQLVerifySchema(db *cmsql.Database) {
	query := "SELECT " + __sqlPurchaseRefund_ListCols + " FROM \"purchase_refund\" WHERE false"
	if _, err := db.SQL(query).Exec(); err != nil {
		db.RecordError(err)
	}
}

func (m *PurchaseRefund) Migration(db *cmsql.Database) {
	var mDBColumnNameAndType map[string]string
	if val, err := migration.GetColumnNamesAndTypes(db, "purchase_refund"); err != nil {
		db.RecordError(err)
		return
	} else {
		mDBColumnNameAndType = val
	}
	mModelColumnNameAndType := map[string]migration.ColumnDef{
		"id": {
			ColumnName:       "id",
			ColumnType:       "dot.ID",
			ColumnDBType:     "int64",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"shop_id": {
			ColumnName:       "shop_id",
			ColumnType:       "dot.ID",
			ColumnDBType:     "int64",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"purchase_order_id": {
			ColumnName:       "purchase_order_id",
			ColumnType:       "dot.ID",
			ColumnDBType:     "int64",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"code": {
			ColumnName:       "code",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"code_norm": {
			ColumnName:       "code_norm",
			ColumnType:       "int",
			ColumnDBType:     "int",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"note": {
			ColumnName:       "note",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"lines": {
			ColumnName:       "lines",
			ColumnType:       "[]*PurchaseRefundLine",
			ColumnDBType:     "[]*struct",
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
		"cancelled_at": {
			ColumnName:       "cancelled_at",
			ColumnType:       "time.Time",
			ColumnDBType:     "struct",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"confirmed_at": {
			ColumnName:       "confirmed_at",
			ColumnType:       "time.Time",
			ColumnDBType:     "struct",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"created_by": {
			ColumnName:       "created_by",
			ColumnType:       "dot.ID",
			ColumnDBType:     "int64",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"updated_by": {
			ColumnName:       "updated_by",
			ColumnType:       "dot.ID",
			ColumnDBType:     "int64",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"cancel_reason": {
			ColumnName:       "cancel_reason",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"status": {
			ColumnName:       "status",
			ColumnType:       "status3.Status",
			ColumnDBType:     "enum",
			ColumnTag:        "int4",
			ColumnEnumValues: []string{"Z", "P", "N"},
		},
		"supplier_id": {
			ColumnName:       "supplier_id",
			ColumnType:       "dot.ID",
			ColumnDBType:     "int64",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"total_amount": {
			ColumnName:       "total_amount",
			ColumnType:       "int",
			ColumnDBType:     "int",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"basket_value": {
			ColumnName:       "basket_value",
			ColumnType:       "int",
			ColumnDBType:     "int",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"rid": {
			ColumnName:       "rid",
			ColumnType:       "dot.ID",
			ColumnDBType:     "int64",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
	}
	if err := migration.Compare(db, "purchase_refund", mModelColumnNameAndType, mDBColumnNameAndType); err != nil {
		db.RecordError(err)
	}
}

func init() {
	__sqlModels = append(__sqlModels, (*PurchaseRefund)(nil))
}

func (m *PurchaseRefund) SQLArgs(opts core.Opts, create bool) []interface{} {
	return []interface{}{
		m.ID,
		m.ShopID,
		m.PurchaseOrderID,
		core.String(m.Code),
		core.Int(m.CodeNorm),
		core.String(m.Note),
		core.JSON{m.Lines},
		core.Time(m.CreatedAt),
		core.Time(m.UpdatedAt),
		core.Time(m.CancelledAt),
		core.Time(m.ConfirmedAt),
		m.CreatedBy,
		m.UpdatedBy,
		core.String(m.CancelReason),
		m.Status,
		m.SupplierID,
		core.Int(m.TotalAmount),
		core.Int(m.BasketValue),
		m.Rid,
	}
}

func (m *PurchaseRefund) SQLScanArgs(opts core.Opts) []interface{} {
	return []interface{}{
		&m.ID,
		&m.ShopID,
		&m.PurchaseOrderID,
		(*core.String)(&m.Code),
		(*core.Int)(&m.CodeNorm),
		(*core.String)(&m.Note),
		core.JSON{&m.Lines},
		(*core.Time)(&m.CreatedAt),
		(*core.Time)(&m.UpdatedAt),
		(*core.Time)(&m.CancelledAt),
		(*core.Time)(&m.ConfirmedAt),
		&m.CreatedBy,
		&m.UpdatedBy,
		(*core.String)(&m.CancelReason),
		&m.Status,
		&m.SupplierID,
		(*core.Int)(&m.TotalAmount),
		(*core.Int)(&m.BasketValue),
		&m.Rid,
	}
}

func (m *PurchaseRefund) SQLScan(opts core.Opts, row *sql.Row) error {
	return row.Scan(m.SQLScanArgs(opts)...)
}

func (ms *PurchaseRefunds) SQLScan(opts core.Opts, rows *sql.Rows) error {
	res := make(PurchaseRefunds, 0, 128)
	for rows.Next() {
		m := new(PurchaseRefund)
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

func (_ *PurchaseRefund) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlPurchaseRefund_Select)
	return nil
}

func (_ *PurchaseRefunds) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlPurchaseRefund_Select)
	return nil
}

func (m *PurchaseRefund) SQLInsert(w SQLWriter) error {
	w.WriteQueryString(__sqlPurchaseRefund_Insert)
	w.WriteRawString(" (")
	w.WriteMarkers(19)
	w.WriteByte(')')
	w.WriteArgs(m.SQLArgs(w.Opts(), true))
	return nil
}

func (ms PurchaseRefunds) SQLInsert(w SQLWriter) error {
	w.WriteQueryString(__sqlPurchaseRefund_Insert)
	w.WriteRawString(" (")
	for i := 0; i < len(ms); i++ {
		w.WriteMarkers(19)
		w.WriteArgs(ms[i].SQLArgs(w.Opts(), true))
		w.WriteRawString("),(")
	}
	w.TrimLast(2)
	return nil
}

func (m *PurchaseRefund) SQLUpsert(w SQLWriter) error {
	m.SQLInsert(w)
	w.WriteQueryString(__sqlPurchaseRefund_UpdateOnConflict)
	w.WriteQueryString(" ")
	w.WriteQueryString(__sqlPurchaseRefund_ListColsOnConflict)
	return nil
}

func (ms PurchaseRefunds) SQLUpsert(w SQLWriter) error {
	ms.SQLInsert(w)
	w.WriteQueryString(__sqlPurchaseRefund_UpdateOnConflict)
	w.WriteQueryString(" ")
	w.WriteQueryString(__sqlPurchaseRefund_ListColsOnConflict)
	return nil
}

func (m *PurchaseRefund) SQLUpdate(w SQLWriter) error {
	now, opts := time.Now(), w.Opts()
	_, _ = now, opts // suppress unuse error
	var flag bool
	w.WriteRawString("UPDATE ")
	w.WriteName("purchase_refund")
	w.WriteRawString(" SET ")
	if m.ID != 0 {
		flag = true
		w.WriteName("id")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.ID)
	}
	if m.ShopID != 0 {
		flag = true
		w.WriteName("shop_id")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.ShopID)
	}
	if m.PurchaseOrderID != 0 {
		flag = true
		w.WriteName("purchase_order_id")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.PurchaseOrderID)
	}
	if m.Code != "" {
		flag = true
		w.WriteName("code")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Code)
	}
	if m.CodeNorm != 0 {
		flag = true
		w.WriteName("code_norm")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.CodeNorm)
	}
	if m.Note != "" {
		flag = true
		w.WriteName("note")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Note)
	}
	if m.Lines != nil {
		flag = true
		w.WriteName("lines")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(core.JSON{m.Lines})
	}
	if !m.CreatedAt.IsZero() {
		flag = true
		w.WriteName("created_at")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.CreatedAt)
	}
	if !m.UpdatedAt.IsZero() {
		flag = true
		w.WriteName("updated_at")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.UpdatedAt)
	}
	if !m.CancelledAt.IsZero() {
		flag = true
		w.WriteName("cancelled_at")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.CancelledAt)
	}
	if !m.ConfirmedAt.IsZero() {
		flag = true
		w.WriteName("confirmed_at")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.ConfirmedAt)
	}
	if m.CreatedBy != 0 {
		flag = true
		w.WriteName("created_by")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.CreatedBy)
	}
	if m.UpdatedBy != 0 {
		flag = true
		w.WriteName("updated_by")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.UpdatedBy)
	}
	if m.CancelReason != "" {
		flag = true
		w.WriteName("cancel_reason")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.CancelReason)
	}
	if m.Status != 0 {
		flag = true
		w.WriteName("status")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Status)
	}
	if m.SupplierID != 0 {
		flag = true
		w.WriteName("supplier_id")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.SupplierID)
	}
	if m.TotalAmount != 0 {
		flag = true
		w.WriteName("total_amount")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.TotalAmount)
	}
	if m.BasketValue != 0 {
		flag = true
		w.WriteName("basket_value")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.BasketValue)
	}
	if m.Rid != 0 {
		flag = true
		w.WriteName("rid")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Rid)
	}
	if !flag {
		return core.ErrNoColumn
	}
	w.TrimLast(1)
	return nil
}

func (m *PurchaseRefund) SQLUpdateAll(w SQLWriter) error {
	w.WriteQueryString(__sqlPurchaseRefund_UpdateAll)
	w.WriteRawString(" = (")
	w.WriteMarkers(19)
	w.WriteByte(')')
	w.WriteArgs(m.SQLArgs(w.Opts(), false))
	return nil
}

type PurchaseRefundHistory map[string]interface{}
type PurchaseRefundHistories []map[string]interface{}

func (m *PurchaseRefundHistory) SQLTableName() string  { return "history.\"purchase_refund\"" }
func (m PurchaseRefundHistories) SQLTableName() string { return "history.\"purchase_refund\"" }

func (m *PurchaseRefundHistory) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlPurchaseRefund_Select_history)
	return nil
}

func (m PurchaseRefundHistories) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlPurchaseRefund_Select_history)
	return nil
}

func (m PurchaseRefundHistory) ID() core.Interface     { return core.Interface{m["id"]} }
func (m PurchaseRefundHistory) ShopID() core.Interface { return core.Interface{m["shop_id"]} }
func (m PurchaseRefundHistory) PurchaseOrderID() core.Interface {
	return core.Interface{m["purchase_order_id"]}
}
func (m PurchaseRefundHistory) Code() core.Interface        { return core.Interface{m["code"]} }
func (m PurchaseRefundHistory) CodeNorm() core.Interface    { return core.Interface{m["code_norm"]} }
func (m PurchaseRefundHistory) Note() core.Interface        { return core.Interface{m["note"]} }
func (m PurchaseRefundHistory) Lines() core.Interface       { return core.Interface{m["lines"]} }
func (m PurchaseRefundHistory) CreatedAt() core.Interface   { return core.Interface{m["created_at"]} }
func (m PurchaseRefundHistory) UpdatedAt() core.Interface   { return core.Interface{m["updated_at"]} }
func (m PurchaseRefundHistory) CancelledAt() core.Interface { return core.Interface{m["cancelled_at"]} }
func (m PurchaseRefundHistory) ConfirmedAt() core.Interface { return core.Interface{m["confirmed_at"]} }
func (m PurchaseRefundHistory) CreatedBy() core.Interface   { return core.Interface{m["created_by"]} }
func (m PurchaseRefundHistory) UpdatedBy() core.Interface   { return core.Interface{m["updated_by"]} }
func (m PurchaseRefundHistory) CancelReason() core.Interface {
	return core.Interface{m["cancel_reason"]}
}
func (m PurchaseRefundHistory) Status() core.Interface      { return core.Interface{m["status"]} }
func (m PurchaseRefundHistory) SupplierID() core.Interface  { return core.Interface{m["supplier_id"]} }
func (m PurchaseRefundHistory) TotalAmount() core.Interface { return core.Interface{m["total_amount"]} }
func (m PurchaseRefundHistory) BasketValue() core.Interface { return core.Interface{m["basket_value"]} }
func (m PurchaseRefundHistory) Rid() core.Interface         { return core.Interface{m["rid"]} }

func (m *PurchaseRefundHistory) SQLScan(opts core.Opts, row *sql.Row) error {
	data := make([]interface{}, 19)
	args := make([]interface{}, 19)
	for i := 0; i < 19; i++ {
		args[i] = &data[i]
	}
	if err := row.Scan(args...); err != nil {
		return err
	}
	res := make(PurchaseRefundHistory, 19)
	res["id"] = data[0]
	res["shop_id"] = data[1]
	res["purchase_order_id"] = data[2]
	res["code"] = data[3]
	res["code_norm"] = data[4]
	res["note"] = data[5]
	res["lines"] = data[6]
	res["created_at"] = data[7]
	res["updated_at"] = data[8]
	res["cancelled_at"] = data[9]
	res["confirmed_at"] = data[10]
	res["created_by"] = data[11]
	res["updated_by"] = data[12]
	res["cancel_reason"] = data[13]
	res["status"] = data[14]
	res["supplier_id"] = data[15]
	res["total_amount"] = data[16]
	res["basket_value"] = data[17]
	res["rid"] = data[18]
	*m = res
	return nil
}

func (ms *PurchaseRefundHistories) SQLScan(opts core.Opts, rows *sql.Rows) error {
	data := make([]interface{}, 19)
	args := make([]interface{}, 19)
	for i := 0; i < 19; i++ {
		args[i] = &data[i]
	}
	res := make(PurchaseRefundHistories, 0, 128)
	for rows.Next() {
		if err := rows.Scan(args...); err != nil {
			return err
		}
		m := make(PurchaseRefundHistory)
		m["id"] = data[0]
		m["shop_id"] = data[1]
		m["purchase_order_id"] = data[2]
		m["code"] = data[3]
		m["code_norm"] = data[4]
		m["note"] = data[5]
		m["lines"] = data[6]
		m["created_at"] = data[7]
		m["updated_at"] = data[8]
		m["cancelled_at"] = data[9]
		m["confirmed_at"] = data[10]
		m["created_by"] = data[11]
		m["updated_by"] = data[12]
		m["cancel_reason"] = data[13]
		m["status"] = data[14]
		m["supplier_id"] = data[15]
		m["total_amount"] = data[16]
		m["basket_value"] = data[17]
		m["rid"] = data[18]
		res = append(res, m)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	*ms = res
	return nil
}
