// +build !generator

// Code generated by generator apix. DO NOT EDIT.

package vht

import (
	context "context"
	fmt "fmt"
	http "net/http"

	capi "o.o/capi"
	httprpc "o.o/capi/httprpc"
)

func init() {
	httprpc.Register(NewServer)
}

func NewServer(builder interface{}, hooks ...httprpc.HooksBuilder) (httprpc.Server, bool) {
	switch builder := builder.(type) {
	case func() UserService:
		return NewUserServiceServer(builder, hooks...), true
	default:
		return nil, false
	}
}

type UserServiceServer struct {
	hooks   httprpc.HooksBuilder
	builder func() UserService
}

func NewUserServiceServer(builder func() UserService, hooks ...httprpc.HooksBuilder) httprpc.Server {
	return &UserServiceServer{
		hooks:   httprpc.ChainHooks(hooks...),
		builder: builder,
	}
}

const UserServicePathPrefix = "/vht.User/"

const Path_User_RegisterUser = "/vht.User/RegisterUser"

func (s *UserServiceServer) PathPrefix() string {
	return UserServicePathPrefix
}

func (s *UserServiceServer) WithHooks(hooks httprpc.HooksBuilder) httprpc.Server {
	result := *s
	result.hooks = httprpc.ChainHooks(s.hooks, hooks)
	return &result
}

func (s *UserServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	hooks := httprpc.WrapHooks(s.hooks)
	ctx, info := req.Context(), &httprpc.HookInfo{Route: req.URL.Path, HTTPRequest: req}
	ctx, err := hooks.RequestReceived(ctx, *info)
	if err != nil {
		httprpc.WriteError(ctx, resp, hooks, *info, err)
		return
	}
	serve, err := httprpc.ParseRequestHeader(req)
	if err != nil {
		httprpc.WriteError(ctx, resp, hooks, *info, err)
		return
	}
	reqMsg, exec, err := s.parseRoute(req.URL.Path, hooks, info)
	if err != nil {
		httprpc.WriteError(ctx, resp, hooks, *info, err)
		return
	}
	serve(ctx, resp, req, hooks, info, reqMsg, exec)
}

func (s *UserServiceServer) parseRoute(path string, hooks httprpc.Hooks, info *httprpc.HookInfo) (reqMsg capi.Message, _ httprpc.ExecFunc, _ error) {
	switch path {
	case "/vht.User/RegisterUser":
		msg := &VHTRegisterUser{}
		fn := func(ctx context.Context) (newCtx context.Context, resp capi.Message, err error) {
			inner := s.builder()
			info.Request, info.Inner = msg, inner
			newCtx, err = hooks.RequestRouted(ctx, *info)
			if err != nil {
				return
			}
			resp, err = inner.RegisterUser(newCtx, msg)
			return
		}
		return msg, fn, nil
	default:
		msg := fmt.Sprintf("no handler for path %q", path)
		return nil, nil, httprpc.BadRouteError(msg, "POST", path)
	}
}
