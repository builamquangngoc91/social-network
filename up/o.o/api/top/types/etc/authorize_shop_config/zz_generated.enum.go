// +build !generator

// Code generated by generator enum. DO NOT EDIT.

package authorize_shop_config

import (
	driver "database/sql/driver"
	fmt "fmt"

	dot "o.o/capi/dot"
	mix "o.o/capi/mix"
)

var __jsonNull = []byte("null")

var enumAuthorizeShopConfigName = map[int]string{
	1: "shipment",
}

var enumAuthorizeShopConfigValue = map[string]int{
	"shipment": 1,
}

func ParseAuthorizeShopConfig(s string) (AuthorizeShopConfig, bool) {
	val, ok := enumAuthorizeShopConfigValue[s]
	return AuthorizeShopConfig(val), ok
}

func ParseAuthorizeShopConfigWithDefault(s string, d AuthorizeShopConfig) AuthorizeShopConfig {
	val, ok := enumAuthorizeShopConfigValue[s]
	if !ok {
		return d
	}
	return AuthorizeShopConfig(val)
}

func (e AuthorizeShopConfig) Enum() int {
	return int(e)
}

func (e AuthorizeShopConfig) Name() string {
	return enumAuthorizeShopConfigName[e.Enum()]
}

func (e AuthorizeShopConfig) String() string {
	s, ok := enumAuthorizeShopConfigName[e.Enum()]
	if ok {
		return s
	}
	return fmt.Sprintf("AuthorizeShopConfig(%v)", e.Enum())
}

func (e AuthorizeShopConfig) MarshalJSON() ([]byte, error) {
	return []byte("\"" + enumAuthorizeShopConfigName[e.Enum()] + "\""), nil
}

func (e *AuthorizeShopConfig) UnmarshalJSON(data []byte) error {
	value, err := mix.UnmarshalJSONEnumInt(enumAuthorizeShopConfigValue, data, "AuthorizeShopConfig")
	if err != nil {
		return err
	}
	*e = AuthorizeShopConfig(value)
	return nil
}

func (e AuthorizeShopConfig) Value() (driver.Value, error) {
	return e.String(), nil
}

func (e *AuthorizeShopConfig) Scan(src interface{}) error {
	value, err := mix.ScanEnumInt(enumAuthorizeShopConfigValue, src, "AuthorizeShopConfig")
	*e = (AuthorizeShopConfig)(value)
	return err
}

func (e AuthorizeShopConfig) Wrap() NullAuthorizeShopConfig {
	return WrapAuthorizeShopConfig(e)
}

func ParseAuthorizeShopConfigWithNull(s dot.NullString, d AuthorizeShopConfig) NullAuthorizeShopConfig {
	if !s.Valid {
		return NullAuthorizeShopConfig{}
	}
	val, ok := enumAuthorizeShopConfigValue[s.String]
	if !ok {
		return d.Wrap()
	}
	return AuthorizeShopConfig(val).Wrap()
}

func WrapAuthorizeShopConfig(enum AuthorizeShopConfig) NullAuthorizeShopConfig {
	return NullAuthorizeShopConfig{Enum: enum, Valid: true}
}

func (n NullAuthorizeShopConfig) Apply(s AuthorizeShopConfig) AuthorizeShopConfig {
	if n.Valid {
		return n.Enum
	}
	return s
}

func (n NullAuthorizeShopConfig) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Enum.Value()
}

func (n *NullAuthorizeShopConfig) Scan(src interface{}) error {
	if src == nil {
		n.Enum, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return n.Enum.Scan(src)
}

func (n NullAuthorizeShopConfig) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return n.Enum.MarshalJSON()
	}
	return __jsonNull, nil
}

func (n *NullAuthorizeShopConfig) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Enum, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return n.Enum.UnmarshalJSON(data)
}