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

type VhtCallHistories []*VhtCallHistory

const __sqlVhtCallHistory_Table = "vht_call_history"
const __sqlVhtCallHistory_ListCols = "\"cdr_id\",\"call_id\",\"sip_call_id\",\"sdk_call_id\",\"cause\",\"q850_cause\",\"from_extension\",\"to_extension\",\"from_number\",\"to_number\",\"duration\",\"direction\",\"time_started\",\"time_connected\",\"time_ended\",\"created_at\",\"updated_at\",\"recording_path\",\"recording_url\",\"record_file_size\",\"etop_account_id\",\"vtiger_account_id\",\"sync_status\",\"o_data\",\"search_norm\""
const __sqlVhtCallHistory_ListColsOnConflict = "\"cdr_id\" = EXCLUDED.\"cdr_id\",\"call_id\" = EXCLUDED.\"call_id\",\"sip_call_id\" = EXCLUDED.\"sip_call_id\",\"sdk_call_id\" = EXCLUDED.\"sdk_call_id\",\"cause\" = EXCLUDED.\"cause\",\"q850_cause\" = EXCLUDED.\"q850_cause\",\"from_extension\" = EXCLUDED.\"from_extension\",\"to_extension\" = EXCLUDED.\"to_extension\",\"from_number\" = EXCLUDED.\"from_number\",\"to_number\" = EXCLUDED.\"to_number\",\"duration\" = EXCLUDED.\"duration\",\"direction\" = EXCLUDED.\"direction\",\"time_started\" = EXCLUDED.\"time_started\",\"time_connected\" = EXCLUDED.\"time_connected\",\"time_ended\" = EXCLUDED.\"time_ended\",\"created_at\" = EXCLUDED.\"created_at\",\"updated_at\" = EXCLUDED.\"updated_at\",\"recording_path\" = EXCLUDED.\"recording_path\",\"recording_url\" = EXCLUDED.\"recording_url\",\"record_file_size\" = EXCLUDED.\"record_file_size\",\"etop_account_id\" = EXCLUDED.\"etop_account_id\",\"vtiger_account_id\" = EXCLUDED.\"vtiger_account_id\",\"sync_status\" = EXCLUDED.\"sync_status\",\"o_data\" = EXCLUDED.\"o_data\",\"search_norm\" = EXCLUDED.\"search_norm\""
const __sqlVhtCallHistory_Insert = "INSERT INTO \"vht_call_history\" (" + __sqlVhtCallHistory_ListCols + ") VALUES"
const __sqlVhtCallHistory_Select = "SELECT " + __sqlVhtCallHistory_ListCols + " FROM \"vht_call_history\""
const __sqlVhtCallHistory_Select_history = "SELECT " + __sqlVhtCallHistory_ListCols + " FROM history.\"vht_call_history\""
const __sqlVhtCallHistory_UpdateAll = "UPDATE \"vht_call_history\" SET (" + __sqlVhtCallHistory_ListCols + ")"
const __sqlVhtCallHistory_UpdateOnConflict = " ON CONFLICT ON CONSTRAINT vht_call_history_pkey DO UPDATE SET"

func (m *VhtCallHistory) SQLTableName() string   { return "vht_call_history" }
func (m *VhtCallHistories) SQLTableName() string { return "vht_call_history" }
func (m *VhtCallHistory) SQLListCols() string    { return __sqlVhtCallHistory_ListCols }

func (m *VhtCallHistory) SQLVerifySchema(db *cmsql.Database) {
	query := "SELECT " + __sqlVhtCallHistory_ListCols + " FROM \"vht_call_history\" WHERE false"
	if _, err := db.SQL(query).Exec(); err != nil {
		db.RecordError(err)
	}
}

func (m *VhtCallHistory) Migration(db *cmsql.Database) {
	var mDBColumnNameAndType map[string]string
	if val, err := migration.GetColumnNamesAndTypes(db, "vht_call_history"); err != nil {
		db.RecordError(err)
		return
	} else {
		mDBColumnNameAndType = val
	}
	mModelColumnNameAndType := map[string]migration.ColumnDef{
		"cdr_id": {
			ColumnName:       "cdr_id",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"call_id": {
			ColumnName:       "call_id",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"sip_call_id": {
			ColumnName:       "sip_call_id",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"sdk_call_id": {
			ColumnName:       "sdk_call_id",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"cause": {
			ColumnName:       "cause",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"q850_cause": {
			ColumnName:       "q850_cause",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"from_extension": {
			ColumnName:       "from_extension",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"to_extension": {
			ColumnName:       "to_extension",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"from_number": {
			ColumnName:       "from_number",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"to_number": {
			ColumnName:       "to_number",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"duration": {
			ColumnName:       "duration",
			ColumnType:       "int",
			ColumnDBType:     "int",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"direction": {
			ColumnName:       "direction",
			ColumnType:       "int",
			ColumnDBType:     "int",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"time_started": {
			ColumnName:       "time_started",
			ColumnType:       "time.Time",
			ColumnDBType:     "struct",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"time_connected": {
			ColumnName:       "time_connected",
			ColumnType:       "time.Time",
			ColumnDBType:     "struct",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"time_ended": {
			ColumnName:       "time_ended",
			ColumnType:       "time.Time",
			ColumnDBType:     "struct",
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
		"recording_path": {
			ColumnName:       "recording_path",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"recording_url": {
			ColumnName:       "recording_url",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"record_file_size": {
			ColumnName:       "record_file_size",
			ColumnType:       "int",
			ColumnDBType:     "int",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"etop_account_id": {
			ColumnName:       "etop_account_id",
			ColumnType:       "dot.ID",
			ColumnDBType:     "int64",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"vtiger_account_id": {
			ColumnName:       "vtiger_account_id",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"sync_status": {
			ColumnName:       "sync_status",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"o_data": {
			ColumnName:       "o_data",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"search_norm": {
			ColumnName:       "search_norm",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
	}
	if err := migration.Compare(db, "vht_call_history", mModelColumnNameAndType, mDBColumnNameAndType); err != nil {
		db.RecordError(err)
	}
}

func init() {
	__sqlModels = append(__sqlModels, (*VhtCallHistory)(nil))
}

func (m *VhtCallHistory) SQLArgs(opts core.Opts, create bool) []interface{} {
	now := time.Now()
	return []interface{}{
		core.String(m.CdrID),
		core.String(m.CallID),
		core.String(m.SipCallID),
		core.String(m.SdkCallID),
		core.String(m.Cause),
		core.String(m.Q850Cause),
		core.String(m.FromExtension),
		core.String(m.ToExtension),
		core.String(m.FromNumber),
		core.String(m.ToNumber),
		core.Int(m.Duration),
		core.Int(m.Direction),
		core.Time(m.TimeStarted),
		core.Time(m.TimeConnected),
		core.Time(m.TimeEnded),
		core.Now(m.CreatedAt, now, create),
		core.Now(m.UpdatedAt, now, true),
		core.String(m.RecordingPath),
		core.String(m.RecordingURL),
		core.Int(m.RecordFileSize),
		m.EtopAccountID,
		core.String(m.VtigerAccountID),
		core.String(m.SyncStatus),
		core.String(m.OData),
		core.String(m.SearchNorm),
	}
}

func (m *VhtCallHistory) SQLScanArgs(opts core.Opts) []interface{} {
	return []interface{}{
		(*core.String)(&m.CdrID),
		(*core.String)(&m.CallID),
		(*core.String)(&m.SipCallID),
		(*core.String)(&m.SdkCallID),
		(*core.String)(&m.Cause),
		(*core.String)(&m.Q850Cause),
		(*core.String)(&m.FromExtension),
		(*core.String)(&m.ToExtension),
		(*core.String)(&m.FromNumber),
		(*core.String)(&m.ToNumber),
		(*core.Int)(&m.Duration),
		(*core.Int)(&m.Direction),
		(*core.Time)(&m.TimeStarted),
		(*core.Time)(&m.TimeConnected),
		(*core.Time)(&m.TimeEnded),
		(*core.Time)(&m.CreatedAt),
		(*core.Time)(&m.UpdatedAt),
		(*core.String)(&m.RecordingPath),
		(*core.String)(&m.RecordingURL),
		(*core.Int)(&m.RecordFileSize),
		&m.EtopAccountID,
		(*core.String)(&m.VtigerAccountID),
		(*core.String)(&m.SyncStatus),
		(*core.String)(&m.OData),
		(*core.String)(&m.SearchNorm),
	}
}

func (m *VhtCallHistory) SQLScan(opts core.Opts, row *sql.Row) error {
	return row.Scan(m.SQLScanArgs(opts)...)
}

func (ms *VhtCallHistories) SQLScan(opts core.Opts, rows *sql.Rows) error {
	res := make(VhtCallHistories, 0, 128)
	for rows.Next() {
		m := new(VhtCallHistory)
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

func (_ *VhtCallHistory) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlVhtCallHistory_Select)
	return nil
}

func (_ *VhtCallHistories) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlVhtCallHistory_Select)
	return nil
}

func (m *VhtCallHistory) SQLInsert(w SQLWriter) error {
	w.WriteQueryString(__sqlVhtCallHistory_Insert)
	w.WriteRawString(" (")
	w.WriteMarkers(25)
	w.WriteByte(')')
	w.WriteArgs(m.SQLArgs(w.Opts(), true))
	return nil
}

func (ms VhtCallHistories) SQLInsert(w SQLWriter) error {
	w.WriteQueryString(__sqlVhtCallHistory_Insert)
	w.WriteRawString(" (")
	for i := 0; i < len(ms); i++ {
		w.WriteMarkers(25)
		w.WriteArgs(ms[i].SQLArgs(w.Opts(), true))
		w.WriteRawString("),(")
	}
	w.TrimLast(2)
	return nil
}

func (m *VhtCallHistory) SQLUpsert(w SQLWriter) error {
	m.SQLInsert(w)
	w.WriteQueryString(__sqlVhtCallHistory_UpdateOnConflict)
	w.WriteQueryString(" ")
	w.WriteQueryString(__sqlVhtCallHistory_ListColsOnConflict)
	return nil
}

func (ms VhtCallHistories) SQLUpsert(w SQLWriter) error {
	ms.SQLInsert(w)
	w.WriteQueryString(__sqlVhtCallHistory_UpdateOnConflict)
	w.WriteQueryString(" ")
	w.WriteQueryString(__sqlVhtCallHistory_ListColsOnConflict)
	return nil
}

func (m *VhtCallHistory) SQLUpdate(w SQLWriter) error {
	now, opts := time.Now(), w.Opts()
	_, _ = now, opts // suppress unuse error
	var flag bool
	w.WriteRawString("UPDATE ")
	w.WriteName("vht_call_history")
	w.WriteRawString(" SET ")
	if m.CdrID != "" {
		flag = true
		w.WriteName("cdr_id")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.CdrID)
	}
	if m.CallID != "" {
		flag = true
		w.WriteName("call_id")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.CallID)
	}
	if m.SipCallID != "" {
		flag = true
		w.WriteName("sip_call_id")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.SipCallID)
	}
	if m.SdkCallID != "" {
		flag = true
		w.WriteName("sdk_call_id")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.SdkCallID)
	}
	if m.Cause != "" {
		flag = true
		w.WriteName("cause")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Cause)
	}
	if m.Q850Cause != "" {
		flag = true
		w.WriteName("q850_cause")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Q850Cause)
	}
	if m.FromExtension != "" {
		flag = true
		w.WriteName("from_extension")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.FromExtension)
	}
	if m.ToExtension != "" {
		flag = true
		w.WriteName("to_extension")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.ToExtension)
	}
	if m.FromNumber != "" {
		flag = true
		w.WriteName("from_number")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.FromNumber)
	}
	if m.ToNumber != "" {
		flag = true
		w.WriteName("to_number")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.ToNumber)
	}
	if m.Duration != 0 {
		flag = true
		w.WriteName("duration")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Duration)
	}
	if m.Direction != 0 {
		flag = true
		w.WriteName("direction")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Direction)
	}
	if !m.TimeStarted.IsZero() {
		flag = true
		w.WriteName("time_started")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.TimeStarted)
	}
	if !m.TimeConnected.IsZero() {
		flag = true
		w.WriteName("time_connected")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.TimeConnected)
	}
	if !m.TimeEnded.IsZero() {
		flag = true
		w.WriteName("time_ended")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.TimeEnded)
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
	if m.RecordingPath != "" {
		flag = true
		w.WriteName("recording_path")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.RecordingPath)
	}
	if m.RecordingURL != "" {
		flag = true
		w.WriteName("recording_url")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.RecordingURL)
	}
	if m.RecordFileSize != 0 {
		flag = true
		w.WriteName("record_file_size")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.RecordFileSize)
	}
	if m.EtopAccountID != 0 {
		flag = true
		w.WriteName("etop_account_id")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.EtopAccountID)
	}
	if m.VtigerAccountID != "" {
		flag = true
		w.WriteName("vtiger_account_id")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.VtigerAccountID)
	}
	if m.SyncStatus != "" {
		flag = true
		w.WriteName("sync_status")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.SyncStatus)
	}
	if m.OData != "" {
		flag = true
		w.WriteName("o_data")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.OData)
	}
	if m.SearchNorm != "" {
		flag = true
		w.WriteName("search_norm")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.SearchNorm)
	}
	if !flag {
		return core.ErrNoColumn
	}
	w.TrimLast(1)
	return nil
}

func (m *VhtCallHistory) SQLUpdateAll(w SQLWriter) error {
	w.WriteQueryString(__sqlVhtCallHistory_UpdateAll)
	w.WriteRawString(" = (")
	w.WriteMarkers(25)
	w.WriteByte(')')
	w.WriteArgs(m.SQLArgs(w.Opts(), false))
	return nil
}

type VhtCallHistoryHistory map[string]interface{}
type VhtCallHistoryHistories []map[string]interface{}

func (m *VhtCallHistoryHistory) SQLTableName() string  { return "history.\"vht_call_history\"" }
func (m VhtCallHistoryHistories) SQLTableName() string { return "history.\"vht_call_history\"" }

func (m *VhtCallHistoryHistory) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlVhtCallHistory_Select_history)
	return nil
}

func (m VhtCallHistoryHistories) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlVhtCallHistory_Select_history)
	return nil
}

func (m VhtCallHistoryHistory) CdrID() core.Interface     { return core.Interface{m["cdr_id"]} }
func (m VhtCallHistoryHistory) CallID() core.Interface    { return core.Interface{m["call_id"]} }
func (m VhtCallHistoryHistory) SipCallID() core.Interface { return core.Interface{m["sip_call_id"]} }
func (m VhtCallHistoryHistory) SdkCallID() core.Interface { return core.Interface{m["sdk_call_id"]} }
func (m VhtCallHistoryHistory) Cause() core.Interface     { return core.Interface{m["cause"]} }
func (m VhtCallHistoryHistory) Q850Cause() core.Interface { return core.Interface{m["q850_cause"]} }
func (m VhtCallHistoryHistory) FromExtension() core.Interface {
	return core.Interface{m["from_extension"]}
}
func (m VhtCallHistoryHistory) ToExtension() core.Interface { return core.Interface{m["to_extension"]} }
func (m VhtCallHistoryHistory) FromNumber() core.Interface  { return core.Interface{m["from_number"]} }
func (m VhtCallHistoryHistory) ToNumber() core.Interface    { return core.Interface{m["to_number"]} }
func (m VhtCallHistoryHistory) Duration() core.Interface    { return core.Interface{m["duration"]} }
func (m VhtCallHistoryHistory) Direction() core.Interface   { return core.Interface{m["direction"]} }
func (m VhtCallHistoryHistory) TimeStarted() core.Interface { return core.Interface{m["time_started"]} }
func (m VhtCallHistoryHistory) TimeConnected() core.Interface {
	return core.Interface{m["time_connected"]}
}
func (m VhtCallHistoryHistory) TimeEnded() core.Interface { return core.Interface{m["time_ended"]} }
func (m VhtCallHistoryHistory) CreatedAt() core.Interface { return core.Interface{m["created_at"]} }
func (m VhtCallHistoryHistory) UpdatedAt() core.Interface { return core.Interface{m["updated_at"]} }
func (m VhtCallHistoryHistory) RecordingPath() core.Interface {
	return core.Interface{m["recording_path"]}
}
func (m VhtCallHistoryHistory) RecordingURL() core.Interface {
	return core.Interface{m["recording_url"]}
}
func (m VhtCallHistoryHistory) RecordFileSize() core.Interface {
	return core.Interface{m["record_file_size"]}
}
func (m VhtCallHistoryHistory) EtopAccountID() core.Interface {
	return core.Interface{m["etop_account_id"]}
}
func (m VhtCallHistoryHistory) VtigerAccountID() core.Interface {
	return core.Interface{m["vtiger_account_id"]}
}
func (m VhtCallHistoryHistory) SyncStatus() core.Interface { return core.Interface{m["sync_status"]} }
func (m VhtCallHistoryHistory) OData() core.Interface      { return core.Interface{m["o_data"]} }
func (m VhtCallHistoryHistory) SearchNorm() core.Interface { return core.Interface{m["search_norm"]} }

func (m *VhtCallHistoryHistory) SQLScan(opts core.Opts, row *sql.Row) error {
	data := make([]interface{}, 25)
	args := make([]interface{}, 25)
	for i := 0; i < 25; i++ {
		args[i] = &data[i]
	}
	if err := row.Scan(args...); err != nil {
		return err
	}
	res := make(VhtCallHistoryHistory, 25)
	res["cdr_id"] = data[0]
	res["call_id"] = data[1]
	res["sip_call_id"] = data[2]
	res["sdk_call_id"] = data[3]
	res["cause"] = data[4]
	res["q850_cause"] = data[5]
	res["from_extension"] = data[6]
	res["to_extension"] = data[7]
	res["from_number"] = data[8]
	res["to_number"] = data[9]
	res["duration"] = data[10]
	res["direction"] = data[11]
	res["time_started"] = data[12]
	res["time_connected"] = data[13]
	res["time_ended"] = data[14]
	res["created_at"] = data[15]
	res["updated_at"] = data[16]
	res["recording_path"] = data[17]
	res["recording_url"] = data[18]
	res["record_file_size"] = data[19]
	res["etop_account_id"] = data[20]
	res["vtiger_account_id"] = data[21]
	res["sync_status"] = data[22]
	res["o_data"] = data[23]
	res["search_norm"] = data[24]
	*m = res
	return nil
}

func (ms *VhtCallHistoryHistories) SQLScan(opts core.Opts, rows *sql.Rows) error {
	data := make([]interface{}, 25)
	args := make([]interface{}, 25)
	for i := 0; i < 25; i++ {
		args[i] = &data[i]
	}
	res := make(VhtCallHistoryHistories, 0, 128)
	for rows.Next() {
		if err := rows.Scan(args...); err != nil {
			return err
		}
		m := make(VhtCallHistoryHistory)
		m["cdr_id"] = data[0]
		m["call_id"] = data[1]
		m["sip_call_id"] = data[2]
		m["sdk_call_id"] = data[3]
		m["cause"] = data[4]
		m["q850_cause"] = data[5]
		m["from_extension"] = data[6]
		m["to_extension"] = data[7]
		m["from_number"] = data[8]
		m["to_number"] = data[9]
		m["duration"] = data[10]
		m["direction"] = data[11]
		m["time_started"] = data[12]
		m["time_connected"] = data[13]
		m["time_ended"] = data[14]
		m["created_at"] = data[15]
		m["updated_at"] = data[16]
		m["recording_path"] = data[17]
		m["recording_url"] = data[18]
		m["record_file_size"] = data[19]
		m["etop_account_id"] = data[20]
		m["vtiger_account_id"] = data[21]
		m["sync_status"] = data[22]
		m["o_data"] = data[23]
		m["search_norm"] = data[24]
		res = append(res, m)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	*ms = res
	return nil
}
