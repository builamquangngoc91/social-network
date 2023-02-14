// +build !generator

// Code generated by generator enum. DO NOT EDIT.

package additional_fee_base_value

import (
	driver "database/sql/driver"
	fmt "fmt"

	dot "o.o/capi/dot"
	mix "o.o/capi/mix"
)

var __jsonNull = []byte("null")

var enumBaseValueTypeName = map[int]string{
	0: "unknown",
	1: "cod_amount",
	2: "main_fee",
	3: "basket_value",
}

var enumBaseValueTypeValue = map[string]int{
	"unknown":      0,
	"cod_amount":   1,
	"main_fee":     2,
	"basket_value": 3,
}

func ParseBaseValueType(s string) (BaseValueType, bool) {
	val, ok := enumBaseValueTypeValue[s]
	return BaseValueType(val), ok
}

func ParseBaseValueTypeWithDefault(s string, d BaseValueType) BaseValueType {
	val, ok := enumBaseValueTypeValue[s]
	if !ok {
		return d
	}
	return BaseValueType(val)
}

func (e BaseValueType) Apply(d BaseValueType) BaseValueType {
	if e == 0 {
		return d
	}
	return e
}

func (e BaseValueType) Enum() int {
	return int(e)
}

func (e BaseValueType) Name() string {
	return enumBaseValueTypeName[e.Enum()]
}

func (e BaseValueType) String() string {
	s, ok := enumBaseValueTypeName[e.Enum()]
	if ok {
		return s
	}
	return fmt.Sprintf("BaseValueType(%v)", e.Enum())
}

func (e BaseValueType) MarshalJSON() ([]byte, error) {
	return []byte("\"" + enumBaseValueTypeName[e.Enum()] + "\""), nil
}

func (e *BaseValueType) UnmarshalJSON(data []byte) error {
	value, err := mix.UnmarshalJSONEnumInt(enumBaseValueTypeValue, data, "BaseValueType")
	if err != nil {
		return err
	}
	*e = BaseValueType(value)
	return nil
}

func (e BaseValueType) Value() (driver.Value, error) {
	if e == 0 {
		return nil, nil
	}
	return e.String(), nil
}

func (e *BaseValueType) Scan(src interface{}) error {
	value, err := mix.ScanEnumInt(enumBaseValueTypeValue, src, "BaseValueType")
	*e = (BaseValueType)(value)
	return err
}

func (e BaseValueType) Wrap() NullBaseValueType {
	return WrapBaseValueType(e)
}

func ParseBaseValueTypeWithNull(s dot.NullString, d BaseValueType) NullBaseValueType {
	if !s.Valid {
		return NullBaseValueType{}
	}
	val, ok := enumBaseValueTypeValue[s.String]
	if !ok {
		return d.Wrap()
	}
	return BaseValueType(val).Wrap()
}

func WrapBaseValueType(enum BaseValueType) NullBaseValueType {
	return NullBaseValueType{Enum: enum, Valid: true}
}

func (n NullBaseValueType) Apply(s BaseValueType) BaseValueType {
	if n.Valid {
		return n.Enum
	}
	return s
}

func (n NullBaseValueType) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Enum.Value()
}

func (n *NullBaseValueType) Scan(src interface{}) error {
	if src == nil {
		n.Enum, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return n.Enum.Scan(src)
}

func (n NullBaseValueType) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return n.Enum.MarshalJSON()
	}
	return __jsonNull, nil
}

func (n *NullBaseValueType) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Enum, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return n.Enum.UnmarshalJSON(data)
}
