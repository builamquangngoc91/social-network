package partner_proto

import (
	"o.o/api/top/types/etc/authorize_shop_config"
	"o.o/capi/dot"
	"o.o/common/jsonx"
)

type AuthorizeShopRequest struct {
	ShopId         dot.ID `json:"shop_id"`
	ExternalShopID string `json:"external_shop_id"`
	ExternalUserID string `json:"external_user_id"`
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	RedirectUrl    string `json:"redirect_url"`
	ShopName       string `json:"shop_name"`
	ExtraToken     string `json:"extra_token"`

	Config []authorize_shop_config.AuthorizeShopConfig `json:"config"`
}

func (m *AuthorizeShopRequest) String() string { return jsonx.MustMarshalToString(m) }

type AuthorizeShopResponse struct {
	Code      string            `json:"code"`
	Msg       string            `json:"msg"`
	Type      string            `json:"type"`
	AuthToken string            `json:"auth_token"`
	ExpiresIn int               `json:"expires_in"`
	AuthUrl   string            `json:"auth_url"`
	Meta      map[string]string `json:"meta"`
}

func (m *AuthorizeShopResponse) String() string { return jsonx.MustMarshalToString(m) }