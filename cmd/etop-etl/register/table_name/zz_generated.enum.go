// +build !generator

// Code generated by generator enum. DO NOT EDIT.

package table_name

import (
	driver "database/sql/driver"
	fmt "fmt"

	dot "o.o/capi/dot"
	mix "o.o/capi/mix"
)

var __jsonNull = []byte("null")

var enumTableNameName = map[int]string{
	0:  "user",
	1:  "account",
	2:  "shop_customer",
	3:  "order",
	4:  "shop",
	5:  "fulfillment",
	6:  "shop_brand",
	7:  "shop_product",
	8:  "account_user",
	9:  "address",
	10: "inventory_variant",
	11: "inventory_voucher",
	12: "invitation",
	13: "money_transaction_shipping",
	14: "product_shop_collection",
	15: "purchase_order",
	16: "purchase_refund",
	17: "receipt",
	18: "refund",
	19: "shipnow_fulfillment",
	20: "shop_carrier",
	21: "shop_category",
	22: "shop_collection",
	23: "shop_customer_group",
	24: "shop_customer_group_customer",
	25: "shop_ledger",
	26: "shop_product_collection",
	27: "shop_stocktake",
	28: "shop_supplier",
	29: "shop_trader",
	30: "shop_trader_address",
	31: "shop_variant",
	32: "shop_variant_supplier",
}

var enumTableNameValue = map[string]int{
	"user":                         0,
	"account":                      1,
	"shop_customer":                2,
	"order":                        3,
	"shop":                         4,
	"fulfillment":                  5,
	"shop_brand":                   6,
	"shop_product":                 7,
	"account_user":                 8,
	"address":                      9,
	"inventory_variant":            10,
	"inventory_voucher":            11,
	"invitation":                   12,
	"money_transaction_shipping":   13,
	"product_shop_collection":      14,
	"purchase_order":               15,
	"purchase_refund":              16,
	"receipt":                      17,
	"refund":                       18,
	"shipnow_fulfillment":          19,
	"shop_carrier":                 20,
	"shop_category":                21,
	"shop_collection":              22,
	"shop_customer_group":          23,
	"shop_customer_group_customer": 24,
	"shop_ledger":                  25,
	"shop_product_collection":      26,
	"shop_stocktake":               27,
	"shop_supplier":                28,
	"shop_trader":                  29,
	"shop_trader_address":          30,
	"shop_variant":                 31,
	"shop_variant_supplier":        32,
}

func ParseTableName(s string) (TableName, bool) {
	val, ok := enumTableNameValue[s]
	return TableName(val), ok
}

func ParseTableNameWithDefault(s string, d TableName) TableName {
	val, ok := enumTableNameValue[s]
	if !ok {
		return d
	}
	return TableName(val)
}

func (e TableName) Apply(d TableName) TableName {
	if e == 0 {
		return d
	}
	return e
}

func (e TableName) Enum() int {
	return int(e)
}

func (e TableName) Name() string {
	return enumTableNameName[e.Enum()]
}

func (e TableName) String() string {
	s, ok := enumTableNameName[e.Enum()]
	if ok {
		return s
	}
	return fmt.Sprintf("TableName(%v)", e.Enum())
}

func (e TableName) MarshalJSON() ([]byte, error) {
	return []byte("\"" + enumTableNameName[e.Enum()] + "\""), nil
}

func (e *TableName) UnmarshalJSON(data []byte) error {
	value, err := mix.UnmarshalJSONEnumInt(enumTableNameValue, data, "TableName")
	if err != nil {
		return err
	}
	*e = TableName(value)
	return nil
}

func (e TableName) Value() (driver.Value, error) {
	if e == 0 {
		return nil, nil
	}
	return e.String(), nil
}

func (e *TableName) Scan(src interface{}) error {
	value, err := mix.ScanEnumInt(enumTableNameValue, src, "TableName")
	*e = (TableName)(value)
	return err
}

func (e TableName) Wrap() NullTableName {
	return WrapTableName(e)
}

func ParseTableNameWithNull(s dot.NullString, d TableName) NullTableName {
	if !s.Valid {
		return NullTableName{}
	}
	val, ok := enumTableNameValue[s.String]
	if !ok {
		return d.Wrap()
	}
	return TableName(val).Wrap()
}

func WrapTableName(enum TableName) NullTableName {
	return NullTableName{Enum: enum, Valid: true}
}

func (n NullTableName) Apply(s TableName) TableName {
	if n.Valid {
		return n.Enum
	}
	return s
}

func (n NullTableName) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Enum.Value()
}

func (n *NullTableName) Scan(src interface{}) error {
	if src == nil {
		n.Enum, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return n.Enum.Scan(src)
}

func (n NullTableName) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return n.Enum.MarshalJSON()
	}
	return __jsonNull, nil
}

func (n *NullTableName) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Enum, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return n.Enum.UnmarshalJSON(data)
}
