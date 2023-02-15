package up

import (
	"strings"

	"social-network/utils/xerror"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

func (r *RegisterRequest) Validate() error {
	if strings.TrimSpace(r.Username) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "username can't be null")
	}
	if strings.TrimSpace(r.Password) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "password can't be null")
	}
	if strings.TrimSpace(r.Fullname) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "fullname can't be null")
	}
	if strings.TrimSpace(r.Email) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "email can't be null")
	}

	return nil
}

type RegisterResponse struct{}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	if strings.TrimSpace(r.Username) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "username can't be null")
	}
	if strings.TrimSpace(r.Password) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "password can't be null")
	}
	return nil
}

type LoginResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type FollowAccountRequest struct {
	AccountID string `json:"account_id"`
}

func (r *FollowAccountRequest) Validate() error {
	if strings.TrimSpace(r.AccountID) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "account_id can't be null")
	}

	return nil
}

type FollowAccountResponse struct{}

type UnFollowAccountRequest struct {
	AccountID string `json:"account_id"`
}

func (r *UnFollowAccountRequest) Validate() error {
	if strings.TrimSpace(r.AccountID) == "" {
		return xerror.ErrorM(xerror.InvalidArgument, nil, "account_id can't be null")
	}

	return nil
}

type UnFollowAccountResponse struct{}
