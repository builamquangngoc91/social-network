// +build !generator

// Code generated by generator enum. DO NOT EDIT.

package fb_internal_source

import (
	driver "database/sql/driver"
	fmt "fmt"

	dot "o.o/capi/dot"
	mix "o.o/capi/mix"
)

var __jsonNull = []byte("null")

var enumFbInternalSourceName = map[int]string{
	0:   "unknown",
	101: "fabo",
	256: "facebook",
}

var enumFbInternalSourceValue = map[string]int{
	"unknown":  0,
	"fabo":     101,
	"facebook": 256,
}

func ParseFbInternalSource(s string) (FbInternalSource, bool) {
	val, ok := enumFbInternalSourceValue[s]
	return FbInternalSource(val), ok
}

func ParseFbInternalSourceWithDefault(s string, d FbInternalSource) FbInternalSource {
	val, ok := enumFbInternalSourceValue[s]
	if !ok {
		return d
	}
	return FbInternalSource(val)
}

func (e FbInternalSource) Apply(d FbInternalSource) FbInternalSource {
	if e == 0 {
		return d
	}
	return e
}

func (e FbInternalSource) Enum() int {
	return int(e)
}

func (e FbInternalSource) Name() string {
	return enumFbInternalSourceName[e.Enum()]
}

func (e FbInternalSource) String() string {
	s, ok := enumFbInternalSourceName[e.Enum()]
	if ok {
		return s
	}
	return fmt.Sprintf("FbInternalSource(%v)", e.Enum())
}

func (e FbInternalSource) MarshalJSON() ([]byte, error) {
	return []byte("\"" + enumFbInternalSourceName[e.Enum()] + "\""), nil
}

func (e *FbInternalSource) UnmarshalJSON(data []byte) error {
	value, err := mix.UnmarshalJSONEnumInt(enumFbInternalSourceValue, data, "FbInternalSource")
	if err != nil {
		return err
	}
	*e = FbInternalSource(value)
	return nil
}

func (e FbInternalSource) Value() (driver.Value, error) {
	if e == 0 {
		return nil, nil
	}
	return int64(e), nil
}

func (e *FbInternalSource) Scan(src interface{}) error {
	value, err := mix.ScanEnumInt(enumFbInternalSourceValue, src, "FbInternalSource")
	*e = (FbInternalSource)(value)
	return err
}

func (e FbInternalSource) Wrap() NullFbInternalSource {
	return WrapFbInternalSource(e)
}

func ParseFbInternalSourceWithNull(s dot.NullString, d FbInternalSource) NullFbInternalSource {
	if !s.Valid {
		return NullFbInternalSource{}
	}
	val, ok := enumFbInternalSourceValue[s.String]
	if !ok {
		return d.Wrap()
	}
	return FbInternalSource(val).Wrap()
}

func WrapFbInternalSource(enum FbInternalSource) NullFbInternalSource {
	return NullFbInternalSource{Enum: enum, Valid: true}
}

func (n NullFbInternalSource) Apply(s FbInternalSource) FbInternalSource {
	if n.Valid {
		return n.Enum
	}
	return s
}

func (n NullFbInternalSource) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Enum.Value()
}

func (n *NullFbInternalSource) Scan(src interface{}) error {
	if src == nil {
		n.Enum, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return n.Enum.Scan(src)
}

func (n NullFbInternalSource) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return n.Enum.MarshalJSON()
	}
	return __jsonNull, nil
}

func (n *NullFbInternalSource) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Enum, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return n.Enum.UnmarshalJSON(data)
}
