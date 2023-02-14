// +build !generator

// Code generated by generator api. DO NOT EDIT.

package location

import (
	context "context"
	time "time"

	meta "o.o/api/meta"
	capi "o.o/capi"
	dot "o.o/capi/dot"
)

type CommandBus struct{ bus capi.Bus }
type QueryBus struct{ bus capi.Bus }

func NewCommandBus(bus capi.Bus) CommandBus { return CommandBus{bus} }
func NewQueryBus(bus capi.Bus) QueryBus     { return QueryBus{bus} }

func (b CommandBus) Dispatch(ctx context.Context, msg interface{ command() }) error {
	return b.bus.Dispatch(ctx, msg)
}
func (b QueryBus) Dispatch(ctx context.Context, msg interface{ query() }) error {
	return b.bus.Dispatch(ctx, msg)
}

type CreateCustomRegionCommand struct {
	Name          string
	Description   string
	ProvinceCodes []string

	Result *CustomRegion `json:"-"`
}

func (h AggregateHandler) HandleCreateCustomRegion(ctx context.Context, msg *CreateCustomRegionCommand) (err error) {
	msg.Result, err = h.inner.CreateCustomRegion(msg.GetArgs(ctx))
	return err
}

type DeleteCustomRegionCommand struct {
	ID dot.ID

	Result struct {
	} `json:"-"`
}

func (h AggregateHandler) HandleDeleteCustomRegion(ctx context.Context, msg *DeleteCustomRegionCommand) (err error) {
	return h.inner.DeleteCustomRegion(msg.GetArgs(ctx))
}

type UpdateCustomRegionCommand struct {
	ID            dot.ID
	Name          string
	ProvinceCodes []string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
	WLPartnerID   dot.ID

	Result struct {
	} `json:"-"`
}

func (h AggregateHandler) HandleUpdateCustomRegion(ctx context.Context, msg *UpdateCustomRegionCommand) (err error) {
	return h.inner.UpdateCustomRegion(msg.GetArgs(ctx))
}

type FindLocationQuery struct {
	Province string
	District string
	Ward     string

	Result *LocationQueryResult `json:"-"`
}

func (h QueryServiceHandler) HandleFindLocation(ctx context.Context, msg *FindLocationQuery) (err error) {
	msg.Result, err = h.inner.FindLocation(msg.GetArgs(ctx))
	return err
}

type FindOrGetLocationQuery struct {
	Province     string
	District     string
	Ward         string
	ProvinceCode string
	DistrictCode string
	WardCode     string

	Result *LocationQueryResult `json:"-"`
}

func (h QueryServiceHandler) HandleFindOrGetLocation(ctx context.Context, msg *FindOrGetLocationQuery) (err error) {
	msg.Result, err = h.inner.FindOrGetLocation(msg.GetArgs(ctx))
	return err
}

type GetAllLocationsQuery struct {
	All          bool
	ProvinceCode string
	DistrictCode string

	Result *GetAllLocationsQueryResult `json:"-"`
}

func (h QueryServiceHandler) HandleGetAllLocations(ctx context.Context, msg *GetAllLocationsQuery) (err error) {
	msg.Result, err = h.inner.GetAllLocations(msg.GetArgs(ctx))
	return err
}

type GetCustomRegionQuery struct {
	ID dot.ID

	Result *CustomRegion `json:"-"`
}

func (h QueryServiceHandler) HandleGetCustomRegion(ctx context.Context, msg *GetCustomRegionQuery) (err error) {
	msg.Result, err = h.inner.GetCustomRegion(msg.GetArgs(ctx))
	return err
}

type GetLocationQuery struct {
	ProvinceCode     string
	DistrictCode     string
	WardCode         string
	LocationCodeType LocationCodeType

	Result *LocationQueryResult `json:"-"`
}

func (h QueryServiceHandler) HandleGetLocation(ctx context.Context, msg *GetLocationQuery) (err error) {
	msg.Result, err = h.inner.GetLocation(msg.GetArgs(ctx))
	return err
}

type ListCustomRegionsQuery struct {
	Result []*CustomRegion `json:"-"`
}

func (h QueryServiceHandler) HandleListCustomRegions(ctx context.Context, msg *ListCustomRegionsQuery) (err error) {
	msg.Result, err = h.inner.ListCustomRegions(msg.GetArgs(ctx))
	return err
}

type ListCustomRegionsByCodeQuery struct {
	ProvinceCode string
	DistrictCode string

	Result []*CustomRegion `json:"-"`
}

func (h QueryServiceHandler) HandleListCustomRegionsByCode(ctx context.Context, msg *ListCustomRegionsByCodeQuery) (err error) {
	msg.Result, err = h.inner.ListCustomRegionsByCode(msg.GetArgs(ctx))
	return err
}

// implement interfaces

func (q *CreateCustomRegionCommand) command() {}
func (q *DeleteCustomRegionCommand) command() {}
func (q *UpdateCustomRegionCommand) command() {}

func (q *FindLocationQuery) query()            {}
func (q *FindOrGetLocationQuery) query()       {}
func (q *GetAllLocationsQuery) query()         {}
func (q *GetCustomRegionQuery) query()         {}
func (q *GetLocationQuery) query()             {}
func (q *ListCustomRegionsQuery) query()       {}
func (q *ListCustomRegionsByCodeQuery) query() {}

// implement conversion

func (q *CreateCustomRegionCommand) GetArgs(ctx context.Context) (_ context.Context, _ *CreateCustomRegionArgs) {
	return ctx,
		&CreateCustomRegionArgs{
			Name:          q.Name,
			Description:   q.Description,
			ProvinceCodes: q.ProvinceCodes,
		}
}

func (q *CreateCustomRegionCommand) SetCreateCustomRegionArgs(args *CreateCustomRegionArgs) {
	q.Name = args.Name
	q.Description = args.Description
	q.ProvinceCodes = args.ProvinceCodes
}

func (q *DeleteCustomRegionCommand) GetArgs(ctx context.Context) (_ context.Context, ID dot.ID) {
	return ctx,
		q.ID
}

func (q *UpdateCustomRegionCommand) GetArgs(ctx context.Context) (_ context.Context, _ *CustomRegion) {
	return ctx,
		&CustomRegion{
			ID:            q.ID,
			Name:          q.Name,
			ProvinceCodes: q.ProvinceCodes,
			Description:   q.Description,
			CreatedAt:     q.CreatedAt,
			UpdatedAt:     q.UpdatedAt,
			DeletedAt:     q.DeletedAt,
			WLPartnerID:   q.WLPartnerID,
		}
}

func (q *UpdateCustomRegionCommand) SetCustomRegion(args *CustomRegion) {
	q.ID = args.ID
	q.Name = args.Name
	q.ProvinceCodes = args.ProvinceCodes
	q.Description = args.Description
	q.CreatedAt = args.CreatedAt
	q.UpdatedAt = args.UpdatedAt
	q.DeletedAt = args.DeletedAt
	q.WLPartnerID = args.WLPartnerID
}

func (q *FindLocationQuery) GetArgs(ctx context.Context) (_ context.Context, _ *FindLocationQueryArgs) {
	return ctx,
		&FindLocationQueryArgs{
			Province: q.Province,
			District: q.District,
			Ward:     q.Ward,
		}
}

func (q *FindLocationQuery) SetFindLocationQueryArgs(args *FindLocationQueryArgs) {
	q.Province = args.Province
	q.District = args.District
	q.Ward = args.Ward
}

func (q *FindOrGetLocationQuery) GetArgs(ctx context.Context) (_ context.Context, _ *FindOrGetLocationQueryArgs) {
	return ctx,
		&FindOrGetLocationQueryArgs{
			Province:     q.Province,
			District:     q.District,
			Ward:         q.Ward,
			ProvinceCode: q.ProvinceCode,
			DistrictCode: q.DistrictCode,
			WardCode:     q.WardCode,
		}
}

func (q *FindOrGetLocationQuery) SetFindOrGetLocationQueryArgs(args *FindOrGetLocationQueryArgs) {
	q.Province = args.Province
	q.District = args.District
	q.Ward = args.Ward
	q.ProvinceCode = args.ProvinceCode
	q.DistrictCode = args.DistrictCode
	q.WardCode = args.WardCode
}

func (q *GetAllLocationsQuery) GetArgs(ctx context.Context) (_ context.Context, _ *GetAllLocationsQueryArgs) {
	return ctx,
		&GetAllLocationsQueryArgs{
			All:          q.All,
			ProvinceCode: q.ProvinceCode,
			DistrictCode: q.DistrictCode,
		}
}

func (q *GetAllLocationsQuery) SetGetAllLocationsQueryArgs(args *GetAllLocationsQueryArgs) {
	q.All = args.All
	q.ProvinceCode = args.ProvinceCode
	q.DistrictCode = args.DistrictCode
}

func (q *GetCustomRegionQuery) GetArgs(ctx context.Context) (_ context.Context, ID dot.ID) {
	return ctx,
		q.ID
}

func (q *GetLocationQuery) GetArgs(ctx context.Context) (_ context.Context, _ *GetLocationQueryArgs) {
	return ctx,
		&GetLocationQueryArgs{
			ProvinceCode:     q.ProvinceCode,
			DistrictCode:     q.DistrictCode,
			WardCode:         q.WardCode,
			LocationCodeType: q.LocationCodeType,
		}
}

func (q *GetLocationQuery) SetGetLocationQueryArgs(args *GetLocationQueryArgs) {
	q.ProvinceCode = args.ProvinceCode
	q.DistrictCode = args.DistrictCode
	q.WardCode = args.WardCode
	q.LocationCodeType = args.LocationCodeType
}

func (q *ListCustomRegionsQuery) GetArgs(ctx context.Context) (_ context.Context, _ *meta.Empty) {
	return ctx,
		&meta.Empty{}
}

func (q *ListCustomRegionsQuery) SetEmpty(args *meta.Empty) {
}

func (q *ListCustomRegionsByCodeQuery) GetArgs(ctx context.Context) (_ context.Context, ProvinceCode string, DistrictCode string) {
	return ctx,
		q.ProvinceCode,
		q.DistrictCode
}

// implement dispatching

type AggregateHandler struct {
	inner Aggregate
}

func NewAggregateHandler(service Aggregate) AggregateHandler { return AggregateHandler{service} }

func (h AggregateHandler) RegisterHandlers(b interface {
	capi.Bus
	AddHandler(handler interface{})
}) CommandBus {
	b.AddHandler(h.HandleCreateCustomRegion)
	b.AddHandler(h.HandleDeleteCustomRegion)
	b.AddHandler(h.HandleUpdateCustomRegion)
	return CommandBus{b}
}

type QueryServiceHandler struct {
	inner QueryService
}

func NewQueryServiceHandler(service QueryService) QueryServiceHandler {
	return QueryServiceHandler{service}
}

func (h QueryServiceHandler) RegisterHandlers(b interface {
	capi.Bus
	AddHandler(handler interface{})
}) QueryBus {
	b.AddHandler(h.HandleFindLocation)
	b.AddHandler(h.HandleFindOrGetLocation)
	b.AddHandler(h.HandleGetAllLocations)
	b.AddHandler(h.HandleGetCustomRegion)
	b.AddHandler(h.HandleGetLocation)
	b.AddHandler(h.HandleListCustomRegions)
	b.AddHandler(h.HandleListCustomRegionsByCode)
	return QueryBus{b}
}
