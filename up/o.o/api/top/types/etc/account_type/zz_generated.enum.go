// +build !generator

// Code generated by generator enum. DO NOT EDIT.

package account_type

import (
	driver "database/sql/driver"
	fmt "fmt"

	dot "o.o/capi/dot"
	mix "o.o/capi/mix"
)

var __jsonNull = []byte("null")

var enumAccountTypeName = map[int]string{
	0:   "unknown",
	21:  "partner",
	33:  "shop",
	35:  "affiliate",
	37:  "carrier",
	101: "etop",
}

var enumAccountTypeValue = map[string]int{
	"unknown":   0,
	"partner":   21,
	"shop":      33,
	"affiliate": 35,
	"carrier":   37,
	"etop":      101,
}

func ParseAccountType(s string) (AccountType, bool) {
	val, ok := enumAccountTypeValue[s]
	return AccountType(val), ok
}

func ParseAccountTypeWithDefault(s string, d AccountType) AccountType {
	val, ok := enumAccountTypeValue[s]
	if !ok {
		return d
	}
	return AccountType(val)
}

func (e AccountType) Apply(d AccountType) AccountType {
	if e == 0 {
		return d
	}
	return e
}

func (e AccountType) Enum() int {
	return int(e)
}

func (e AccountType) Name() string {
	return enumAccountTypeName[e.Enum()]
}

func (e AccountType) String() string {
	s, ok := enumAccountTypeName[e.Enum()]
	if ok {
		return s
	}
	return fmt.Sprintf("AccountType(%v)", e.Enum())
}

func (e AccountType) MarshalJSON() ([]byte, error) {
	return []byte("\"" + enumAccountTypeName[e.Enum()] + "\""), nil
}

func (e *AccountType) UnmarshalJSON(data []byte) error {
	value, err := mix.UnmarshalJSONEnumInt(enumAccountTypeValue, data, "AccountType")
	if err != nil {
		return err
	}
	*e = AccountType(value)
	return nil
}

func (e AccountType) Value() (driver.Value, error) {
	if e == 0 {
		return nil, nil
	}
	return e.String(), nil
}

func (e *AccountType) Scan(src interface{}) error {
	value, err := mix.ScanEnumInt(enumAccountTypeValue, src, "AccountType")
	*e = (AccountType)(value)
	return err
}

func (e AccountType) Wrap() NullAccountType {
	return WrapAccountType(e)
}

func ParseAccountTypeWithNull(s dot.NullString, d AccountType) NullAccountType {
	if !s.Valid {
		return NullAccountType{}
	}
	val, ok := enumAccountTypeValue[s.String]
	if !ok {
		return d.Wrap()
	}
	return AccountType(val).Wrap()
}

func WrapAccountType(enum AccountType) NullAccountType {
	return NullAccountType{Enum: enum, Valid: true}
}

func (n NullAccountType) Apply(s AccountType) AccountType {
	if n.Valid {
		return n.Enum
	}
	return s
}

func (n NullAccountType) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Enum.Value()
}

func (n *NullAccountType) Scan(src interface{}) error {
	if src == nil {
		n.Enum, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return n.Enum.Scan(src)
}

func (n NullAccountType) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return n.Enum.MarshalJSON()
	}
	return __jsonNull, nil
}

func (n *NullAccountType) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Enum, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return n.Enum.UnmarshalJSON(data)
}