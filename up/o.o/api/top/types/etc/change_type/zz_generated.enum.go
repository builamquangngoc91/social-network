// +build !generator

// Code generated by generator enum. DO NOT EDIT.

package change_type

import (
	driver "database/sql/driver"
	fmt "fmt"

	dot "o.o/capi/dot"
	mix "o.o/capi/mix"
)

var __jsonNull = []byte("null")

var enumChangeTypeName = map[int]string{
	0: "unknown",
	1: "update",
	2: "create",
	3: "delete",
}

var enumChangeTypeValue = map[string]int{
	"unknown": 0,
	"update":  1,
	"create":  2,
	"delete":  3,
}

func ParseChangeType(s string) (ChangeType, bool) {
	val, ok := enumChangeTypeValue[s]
	return ChangeType(val), ok
}

func ParseChangeTypeWithDefault(s string, d ChangeType) ChangeType {
	val, ok := enumChangeTypeValue[s]
	if !ok {
		return d
	}
	return ChangeType(val)
}

func (e ChangeType) Apply(d ChangeType) ChangeType {
	if e == 0 {
		return d
	}
	return e
}

func (e ChangeType) Enum() int {
	return int(e)
}

func (e ChangeType) Name() string {
	return enumChangeTypeName[e.Enum()]
}

func (e ChangeType) String() string {
	s, ok := enumChangeTypeName[e.Enum()]
	if ok {
		return s
	}
	return fmt.Sprintf("ChangeType(%v)", e.Enum())
}

func (e ChangeType) MarshalJSON() ([]byte, error) {
	return []byte("\"" + enumChangeTypeName[e.Enum()] + "\""), nil
}

func (e *ChangeType) UnmarshalJSON(data []byte) error {
	value, err := mix.UnmarshalJSONEnumInt(enumChangeTypeValue, data, "ChangeType")
	if err != nil {
		return err
	}
	*e = ChangeType(value)
	return nil
}

func (e ChangeType) Value() (driver.Value, error) {
	if e == 0 {
		return nil, nil
	}
	return e.String(), nil
}

func (e *ChangeType) Scan(src interface{}) error {
	value, err := mix.ScanEnumInt(enumChangeTypeValue, src, "ChangeType")
	*e = (ChangeType)(value)
	return err
}

func (e ChangeType) Wrap() NullChangeType {
	return WrapChangeType(e)
}

func ParseChangeTypeWithNull(s dot.NullString, d ChangeType) NullChangeType {
	if !s.Valid {
		return NullChangeType{}
	}
	val, ok := enumChangeTypeValue[s.String]
	if !ok {
		return d.Wrap()
	}
	return ChangeType(val).Wrap()
}

func WrapChangeType(enum ChangeType) NullChangeType {
	return NullChangeType{Enum: enum, Valid: true}
}

func (n NullChangeType) Apply(s ChangeType) ChangeType {
	if n.Valid {
		return n.Enum
	}
	return s
}

func (n NullChangeType) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Enum.Value()
}

func (n *NullChangeType) Scan(src interface{}) error {
	if src == nil {
		n.Enum, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return n.Enum.Scan(src)
}

func (n NullChangeType) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return n.Enum.MarshalJSON()
	}
	return __jsonNull, nil
}

func (n *NullChangeType) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Enum, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return n.Enum.UnmarshalJSON(data)
}
