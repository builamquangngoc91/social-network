// +build !generator

// Code generated by generator apix. DO NOT EDIT.

package whitelabel

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
	case func() ImportService:
		return NewImportServiceServer(builder, hooks...), true
	default:
		return nil, false
	}
}

type ImportServiceServer struct {
	hooks   httprpc.HooksBuilder
	builder func() ImportService
}

func NewImportServiceServer(builder func() ImportService, hooks ...httprpc.HooksBuilder) httprpc.Server {
	return &ImportServiceServer{
		hooks:   httprpc.ChainHooks(hooks...),
		builder: builder,
	}
}

const ImportServicePathPrefix = "/partner.Import/"

const Path_Import_Brands = "/partner.Import/Brands"
const Path_Import_Categories = "/partner.Import/Categories"
const Path_Import_Collections = "/partner.Import/Collections"
const Path_Import_Customers = "/partner.Import/Customers"
const Path_Import_ProductCollections = "/partner.Import/ProductCollections"
const Path_Import_Products = "/partner.Import/Products"
const Path_Import_Variants = "/partner.Import/Variants"

func (s *ImportServiceServer) PathPrefix() string {
	return ImportServicePathPrefix
}

func (s *ImportServiceServer) WithHooks(hooks httprpc.HooksBuilder) httprpc.Server {
	result := *s
	result.hooks = httprpc.ChainHooks(s.hooks, hooks)
	return &result
}

func (s *ImportServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
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

func (s *ImportServiceServer) parseRoute(path string, hooks httprpc.Hooks, info *httprpc.HookInfo) (reqMsg capi.Message, _ httprpc.ExecFunc, _ error) {
	switch path {
	case "/partner.Import/Brands":
		msg := &ImportBrandsRequest{}
		fn := func(ctx context.Context) (newCtx context.Context, resp capi.Message, err error) {
			inner := s.builder()
			info.Request, info.Inner = msg, inner
			newCtx, err = hooks.RequestRouted(ctx, *info)
			if err != nil {
				return
			}
			resp, err = inner.Brands(newCtx, msg)
			return
		}
		return msg, fn, nil
	case "/partner.Import/Categories":
		msg := &ImportCategoriesRequest{}
		fn := func(ctx context.Context) (newCtx context.Context, resp capi.Message, err error) {
			inner := s.builder()
			info.Request, info.Inner = msg, inner
			newCtx, err = hooks.RequestRouted(ctx, *info)
			if err != nil {
				return
			}
			resp, err = inner.Categories(newCtx, msg)
			return
		}
		return msg, fn, nil
	case "/partner.Import/Collections":
		msg := &ImportCollectionsRequest{}
		fn := func(ctx context.Context) (newCtx context.Context, resp capi.Message, err error) {
			inner := s.builder()
			info.Request, info.Inner = msg, inner
			newCtx, err = hooks.RequestRouted(ctx, *info)
			if err != nil {
				return
			}
			resp, err = inner.Collections(newCtx, msg)
			return
		}
		return msg, fn, nil
	case "/partner.Import/Customers":
		msg := &ImportCustomersRequest{}
		fn := func(ctx context.Context) (newCtx context.Context, resp capi.Message, err error) {
			inner := s.builder()
			info.Request, info.Inner = msg, inner
			newCtx, err = hooks.RequestRouted(ctx, *info)
			if err != nil {
				return
			}
			resp, err = inner.Customers(newCtx, msg)
			return
		}
		return msg, fn, nil
	case "/partner.Import/ProductCollections":
		msg := &ImportProductCollectionsRequest{}
		fn := func(ctx context.Context) (newCtx context.Context, resp capi.Message, err error) {
			inner := s.builder()
			info.Request, info.Inner = msg, inner
			newCtx, err = hooks.RequestRouted(ctx, *info)
			if err != nil {
				return
			}
			resp, err = inner.ProductCollections(newCtx, msg)
			return
		}
		return msg, fn, nil
	case "/partner.Import/Products":
		msg := &ImportProductsRequest{}
		fn := func(ctx context.Context) (newCtx context.Context, resp capi.Message, err error) {
			inner := s.builder()
			info.Request, info.Inner = msg, inner
			newCtx, err = hooks.RequestRouted(ctx, *info)
			if err != nil {
				return
			}
			resp, err = inner.Products(newCtx, msg)
			return
		}
		return msg, fn, nil
	case "/partner.Import/Variants":
		msg := &ImportShopVariantsRequest{}
		fn := func(ctx context.Context) (newCtx context.Context, resp capi.Message, err error) {
			inner := s.builder()
			info.Request, info.Inner = msg, inner
			newCtx, err = hooks.RequestRouted(ctx, *info)
			if err != nil {
				return
			}
			resp, err = inner.Variants(newCtx, msg)
			return
		}
		return msg, fn, nil
	default:
		msg := fmt.Sprintf("no handler for path %q", path)
		return nil, nil, httprpc.BadRouteError(msg, "POST", path)
	}
}