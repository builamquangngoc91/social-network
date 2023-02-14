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

type ShopTraderAddresses []*ShopTraderAddress

const __sqlShopTraderAddress_Table = "shop_trader_address"
const __sqlShopTraderAddress_ListCols = "\"id\",\"shop_id\",\"trader_id\",\"full_name\",\"phone\",\"email\",\"company\",\"address1\",\"address2\",\"district_code\",\"ward_code\",\"position\",\"is_default\",\"coordinates\",\"created_at\",\"updated_at\",\"status\",\"rid\""
const __sqlShopTraderAddress_ListColsOnConflict = "\"id\" = EXCLUDED.\"id\",\"shop_id\" = EXCLUDED.\"shop_id\",\"trader_id\" = EXCLUDED.\"trader_id\",\"full_name\" = EXCLUDED.\"full_name\",\"phone\" = EXCLUDED.\"phone\",\"email\" = EXCLUDED.\"email\",\"company\" = EXCLUDED.\"company\",\"address1\" = EXCLUDED.\"address1\",\"address2\" = EXCLUDED.\"address2\",\"district_code\" = EXCLUDED.\"district_code\",\"ward_code\" = EXCLUDED.\"ward_code\",\"position\" = EXCLUDED.\"position\",\"is_default\" = EXCLUDED.\"is_default\",\"coordinates\" = EXCLUDED.\"coordinates\",\"created_at\" = EXCLUDED.\"created_at\",\"updated_at\" = EXCLUDED.\"updated_at\",\"status\" = EXCLUDED.\"status\",\"rid\" = EXCLUDED.\"rid\""
const __sqlShopTraderAddress_Insert = "INSERT INTO \"shop_trader_address\" (" + __sqlShopTraderAddress_ListCols + ") VALUES"
const __sqlShopTraderAddress_Select = "SELECT " + __sqlShopTraderAddress_ListCols + " FROM \"shop_trader_address\""
const __sqlShopTraderAddress_Select_history = "SELECT " + __sqlShopTraderAddress_ListCols + " FROM history.\"shop_trader_address\""
const __sqlShopTraderAddress_UpdateAll = "UPDATE \"shop_trader_address\" SET (" + __sqlShopTraderAddress_ListCols + ")"
const __sqlShopTraderAddress_UpdateOnConflict = " ON CONFLICT ON CONSTRAINT shop_trader_address_pkey DO UPDATE SET"

func (m *ShopTraderAddress) SQLTableName() string   { return "shop_trader_address" }
func (m *ShopTraderAddresses) SQLTableName() string { return "shop_trader_address" }
func (m *ShopTraderAddress) SQLListCols() string    { return __sqlShopTraderAddress_ListCols }

func (m *ShopTraderAddress) SQLVerifySchema(db *cmsql.Database) {
	query := "SELECT " + __sqlShopTraderAddress_ListCols + " FROM \"shop_trader_address\" WHERE false"
	if _, err := db.SQL(query).Exec(); err != nil {
		db.RecordError(err)
	}
}

func (m *ShopTraderAddress) Migration(db *cmsql.Database) {
	var mDBColumnNameAndType map[string]string
	if val, err := migration.GetColumnNamesAndTypes(db, "shop_trader_address"); err != nil {
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
		"trader_id": {
			ColumnName:       "trader_id",
			ColumnType:       "dot.ID",
			ColumnDBType:     "int64",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"full_name": {
			ColumnName:       "full_name",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"phone": {
			ColumnName:       "phone",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"email": {
			ColumnName:       "email",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"company": {
			ColumnName:       "company",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"address1": {
			ColumnName:       "address1",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"address2": {
			ColumnName:       "address2",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"district_code": {
			ColumnName:       "district_code",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"ward_code": {
			ColumnName:       "ward_code",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"position": {
			ColumnName:       "position",
			ColumnType:       "string",
			ColumnDBType:     "string",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"is_default": {
			ColumnName:       "is_default",
			ColumnType:       "bool",
			ColumnDBType:     "bool",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
		"coordinates": {
			ColumnName:       "coordinates",
			ColumnType:       "*addressmodel.Coordinates",
			ColumnDBType:     "*struct",
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
		"status": {
			ColumnName:       "status",
			ColumnType:       "status3.Status",
			ColumnDBType:     "enum",
			ColumnTag:        "int2",
			ColumnEnumValues: []string{"Z", "P", "N"},
		},
		"rid": {
			ColumnName:       "rid",
			ColumnType:       "dot.ID",
			ColumnDBType:     "int64",
			ColumnTag:        "",
			ColumnEnumValues: []string{},
		},
	}
	if err := migration.Compare(db, "shop_trader_address", mModelColumnNameAndType, mDBColumnNameAndType); err != nil {
		db.RecordError(err)
	}
}

func init() {
	__sqlModels = append(__sqlModels, (*ShopTraderAddress)(nil))
}

func (m *ShopTraderAddress) SQLArgs(opts core.Opts, create bool) []interface{} {
	return []interface{}{
		m.ID,
		m.ShopID,
		m.TraderID,
		core.String(m.FullName),
		core.String(m.Phone),
		core.String(m.Email),
		core.String(m.Company),
		core.String(m.Address1),
		core.String(m.Address2),
		core.String(m.DistrictCode),
		core.String(m.WardCode),
		core.String(m.Position),
		core.Bool(m.IsDefault),
		core.JSON{m.Coordinates},
		core.Time(m.CreatedAt),
		core.Time(m.UpdatedAt),
		m.Status,
		m.Rid,
	}
}

func (m *ShopTraderAddress) SQLScanArgs(opts core.Opts) []interface{} {
	return []interface{}{
		&m.ID,
		&m.ShopID,
		&m.TraderID,
		(*core.String)(&m.FullName),
		(*core.String)(&m.Phone),
		(*core.String)(&m.Email),
		(*core.String)(&m.Company),
		(*core.String)(&m.Address1),
		(*core.String)(&m.Address2),
		(*core.String)(&m.DistrictCode),
		(*core.String)(&m.WardCode),
		(*core.String)(&m.Position),
		(*core.Bool)(&m.IsDefault),
		core.JSON{&m.Coordinates},
		(*core.Time)(&m.CreatedAt),
		(*core.Time)(&m.UpdatedAt),
		&m.Status,
		&m.Rid,
	}
}

func (m *ShopTraderAddress) SQLScan(opts core.Opts, row *sql.Row) error {
	return row.Scan(m.SQLScanArgs(opts)...)
}

func (ms *ShopTraderAddresses) SQLScan(opts core.Opts, rows *sql.Rows) error {
	res := make(ShopTraderAddresses, 0, 128)
	for rows.Next() {
		m := new(ShopTraderAddress)
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

func (_ *ShopTraderAddress) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlShopTraderAddress_Select)
	return nil
}

func (_ *ShopTraderAddresses) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlShopTraderAddress_Select)
	return nil
}

func (m *ShopTraderAddress) SQLInsert(w SQLWriter) error {
	w.WriteQueryString(__sqlShopTraderAddress_Insert)
	w.WriteRawString(" (")
	w.WriteMarkers(18)
	w.WriteByte(')')
	w.WriteArgs(m.SQLArgs(w.Opts(), true))
	return nil
}

func (ms ShopTraderAddresses) SQLInsert(w SQLWriter) error {
	w.WriteQueryString(__sqlShopTraderAddress_Insert)
	w.WriteRawString(" (")
	for i := 0; i < len(ms); i++ {
		w.WriteMarkers(18)
		w.WriteArgs(ms[i].SQLArgs(w.Opts(), true))
		w.WriteRawString("),(")
	}
	w.TrimLast(2)
	return nil
}

func (m *ShopTraderAddress) SQLUpsert(w SQLWriter) error {
	m.SQLInsert(w)
	w.WriteQueryString(__sqlShopTraderAddress_UpdateOnConflict)
	w.WriteQueryString(" ")
	w.WriteQueryString(__sqlShopTraderAddress_ListColsOnConflict)
	return nil
}

func (ms ShopTraderAddresses) SQLUpsert(w SQLWriter) error {
	ms.SQLInsert(w)
	w.WriteQueryString(__sqlShopTraderAddress_UpdateOnConflict)
	w.WriteQueryString(" ")
	w.WriteQueryString(__sqlShopTraderAddress_ListColsOnConflict)
	return nil
}

func (m *ShopTraderAddress) SQLUpdate(w SQLWriter) error {
	now, opts := time.Now(), w.Opts()
	_, _ = now, opts // suppress unuse error
	var flag bool
	w.WriteRawString("UPDATE ")
	w.WriteName("shop_trader_address")
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
	if m.TraderID != 0 {
		flag = true
		w.WriteName("trader_id")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.TraderID)
	}
	if m.FullName != "" {
		flag = true
		w.WriteName("full_name")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.FullName)
	}
	if m.Phone != "" {
		flag = true
		w.WriteName("phone")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Phone)
	}
	if m.Email != "" {
		flag = true
		w.WriteName("email")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Email)
	}
	if m.Company != "" {
		flag = true
		w.WriteName("company")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Company)
	}
	if m.Address1 != "" {
		flag = true
		w.WriteName("address1")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Address1)
	}
	if m.Address2 != "" {
		flag = true
		w.WriteName("address2")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Address2)
	}
	if m.DistrictCode != "" {
		flag = true
		w.WriteName("district_code")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.DistrictCode)
	}
	if m.WardCode != "" {
		flag = true
		w.WriteName("ward_code")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.WardCode)
	}
	if m.Position != "" {
		flag = true
		w.WriteName("position")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Position)
	}
	if m.IsDefault {
		flag = true
		w.WriteName("is_default")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.IsDefault)
	}
	if m.Coordinates != nil {
		flag = true
		w.WriteName("coordinates")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(core.JSON{m.Coordinates})
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
	if m.Status != 0 {
		flag = true
		w.WriteName("status")
		w.WriteByte('=')
		w.WriteMarker()
		w.WriteByte(',')
		w.WriteArg(m.Status)
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

func (m *ShopTraderAddress) SQLUpdateAll(w SQLWriter) error {
	w.WriteQueryString(__sqlShopTraderAddress_UpdateAll)
	w.WriteRawString(" = (")
	w.WriteMarkers(18)
	w.WriteByte(')')
	w.WriteArgs(m.SQLArgs(w.Opts(), false))
	return nil
}

type ShopTraderAddressHistory map[string]interface{}
type ShopTraderAddressHistories []map[string]interface{}

func (m *ShopTraderAddressHistory) SQLTableName() string  { return "history.\"shop_trader_address\"" }
func (m ShopTraderAddressHistories) SQLTableName() string { return "history.\"shop_trader_address\"" }

func (m *ShopTraderAddressHistory) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlShopTraderAddress_Select_history)
	return nil
}

func (m ShopTraderAddressHistories) SQLSelect(w SQLWriter) error {
	w.WriteQueryString(__sqlShopTraderAddress_Select_history)
	return nil
}

func (m ShopTraderAddressHistory) ID() core.Interface       { return core.Interface{m["id"]} }
func (m ShopTraderAddressHistory) ShopID() core.Interface   { return core.Interface{m["shop_id"]} }
func (m ShopTraderAddressHistory) TraderID() core.Interface { return core.Interface{m["trader_id"]} }
func (m ShopTraderAddressHistory) FullName() core.Interface { return core.Interface{m["full_name"]} }
func (m ShopTraderAddressHistory) Phone() core.Interface    { return core.Interface{m["phone"]} }
func (m ShopTraderAddressHistory) Email() core.Interface    { return core.Interface{m["email"]} }
func (m ShopTraderAddressHistory) Company() core.Interface  { return core.Interface{m["company"]} }
func (m ShopTraderAddressHistory) Address1() core.Interface { return core.Interface{m["address1"]} }
func (m ShopTraderAddressHistory) Address2() core.Interface { return core.Interface{m["address2"]} }
func (m ShopTraderAddressHistory) DistrictCode() core.Interface {
	return core.Interface{m["district_code"]}
}
func (m ShopTraderAddressHistory) WardCode() core.Interface  { return core.Interface{m["ward_code"]} }
func (m ShopTraderAddressHistory) Position() core.Interface  { return core.Interface{m["position"]} }
func (m ShopTraderAddressHistory) IsDefault() core.Interface { return core.Interface{m["is_default"]} }
func (m ShopTraderAddressHistory) Coordinates() core.Interface {
	return core.Interface{m["coordinates"]}
}
func (m ShopTraderAddressHistory) CreatedAt() core.Interface { return core.Interface{m["created_at"]} }
func (m ShopTraderAddressHistory) UpdatedAt() core.Interface { return core.Interface{m["updated_at"]} }
func (m ShopTraderAddressHistory) Status() core.Interface    { return core.Interface{m["status"]} }
func (m ShopTraderAddressHistory) Rid() core.Interface       { return core.Interface{m["rid"]} }

func (m *ShopTraderAddressHistory) SQLScan(opts core.Opts, row *sql.Row) error {
	data := make([]interface{}, 18)
	args := make([]interface{}, 18)
	for i := 0; i < 18; i++ {
		args[i] = &data[i]
	}
	if err := row.Scan(args...); err != nil {
		return err
	}
	res := make(ShopTraderAddressHistory, 18)
	res["id"] = data[0]
	res["shop_id"] = data[1]
	res["trader_id"] = data[2]
	res["full_name"] = data[3]
	res["phone"] = data[4]
	res["email"] = data[5]
	res["company"] = data[6]
	res["address1"] = data[7]
	res["address2"] = data[8]
	res["district_code"] = data[9]
	res["ward_code"] = data[10]
	res["position"] = data[11]
	res["is_default"] = data[12]
	res["coordinates"] = data[13]
	res["created_at"] = data[14]
	res["updated_at"] = data[15]
	res["status"] = data[16]
	res["rid"] = data[17]
	*m = res
	return nil
}

func (ms *ShopTraderAddressHistories) SQLScan(opts core.Opts, rows *sql.Rows) error {
	data := make([]interface{}, 18)
	args := make([]interface{}, 18)
	for i := 0; i < 18; i++ {
		args[i] = &data[i]
	}
	res := make(ShopTraderAddressHistories, 0, 128)
	for rows.Next() {
		if err := rows.Scan(args...); err != nil {
			return err
		}
		m := make(ShopTraderAddressHistory)
		m["id"] = data[0]
		m["shop_id"] = data[1]
		m["trader_id"] = data[2]
		m["full_name"] = data[3]
		m["phone"] = data[4]
		m["email"] = data[5]
		m["company"] = data[6]
		m["address1"] = data[7]
		m["address2"] = data[8]
		m["district_code"] = data[9]
		m["ward_code"] = data[10]
		m["position"] = data[11]
		m["is_default"] = data[12]
		m["coordinates"] = data[13]
		m["created_at"] = data[14]
		m["updated_at"] = data[15]
		m["status"] = data[16]
		m["rid"] = data[17]
		res = append(res, m)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	*ms = res
	return nil
}
