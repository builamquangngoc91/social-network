// +build !generator

// Code generated by generator api. DO NOT EDIT.

package etelecom

import (
	context "context"
	time "time"

	call_direction "o.o/api/etelecom/call_direction"
	call_state "o.o/api/etelecom/call_state"
	mobile_network "o.o/api/etelecom/mobile_network"
	meta "o.o/api/meta"
	common "o.o/api/top/types/common"
	connection_type "o.o/api/top/types/etc/connection_type"
	payment_method "o.o/api/top/types/etc/payment_method"
	status3 "o.o/api/top/types/etc/status3"
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

type ActivateTenantCommand struct {
	OwnerID            dot.ID
	TenantID           dot.ID
	HotlineID          dot.ID
	ConnectionID       dot.ID
	ConnectionProvider connection_type.ConnectionProvider

	Result *Tenant `json:"-"`
}

func (h AggregateHandler) HandleActivateTenant(ctx context.Context, msg *ActivateTenantCommand) (err error) {
	msg.Result, err = h.inner.ActivateTenant(msg.GetArgs(ctx))
	return err
}

type AssignUserToExtensionCommand struct {
	AccountID   dot.ID
	UserID      dot.ID
	ExtensionID dot.ID

	Result struct {
	} `json:"-"`
}

func (h AggregateHandler) HandleAssignUserToExtension(ctx context.Context, msg *AssignUserToExtensionCommand) (err error) {
	return h.inner.AssignUserToExtension(msg.GetArgs(ctx))
}

type CreateCallLogCommand struct {
	ExternalSessionID string
	Direction         call_direction.CallDirection
	Caller            string
	Callee            string
	ExtensionID       dot.ID
	AccountID         dot.ID
	OwnerID           dot.ID
	ContactID         dot.ID
	CallState         call_state.CallState
	StartedAt         time.Time
	EndedAt           time.Time

	Result *CallLog `json:"-"`
}

func (h AggregateHandler) HandleCreateCallLog(ctx context.Context, msg *CreateCallLogCommand) (err error) {
	msg.Result, err = h.inner.CreateCallLog(msg.GetArgs(ctx))
	return err
}

type CreateExtensionCommand struct {
	ExtensionNumber   string
	UserID            dot.ID
	AccountID         dot.ID
	ExtensionPassword string
	HotlineID         dot.ID
	OwnerID           dot.ID
	SubscriptionID    dot.ID
	ExpiresAt         time.Time

	Result *Extension `json:"-"`
}

func (h AggregateHandler) HandleCreateExtension(ctx context.Context, msg *CreateExtensionCommand) (err error) {
	msg.Result, err = h.inner.CreateExtension(msg.GetArgs(ctx))
	return err
}

type CreateExtensionBySubscriptionCommand struct {
	SubscriptionID     dot.ID
	SubscriptionPlanID dot.ID
	PaymentMethod      payment_method.PaymentMethod
	AccountID          dot.ID
	ExtensionNumber    string
	UserID             dot.ID
	HotlineID          dot.ID
	OwnerID            dot.ID

	Result *Extension `json:"-"`
}

func (h AggregateHandler) HandleCreateExtensionBySubscription(ctx context.Context, msg *CreateExtensionBySubscriptionCommand) (err error) {
	msg.Result, err = h.inner.CreateExtensionBySubscription(msg.GetArgs(ctx))
	return err
}

type CreateHotlineCommand struct {
	OwnerID      dot.ID
	Name         string
	Hotline      string
	Network      mobile_network.MobileNetwork
	ConnectionID dot.ID
	Status       status3.Status
	Description  string
	IsFreeCharge dot.NullBool

	Result *Hotline `json:"-"`
}

func (h AggregateHandler) HandleCreateHotline(ctx context.Context, msg *CreateHotlineCommand) (err error) {
	msg.Result, err = h.inner.CreateHotline(msg.GetArgs(ctx))
	return err
}

type CreateOrUpdateCallLogFromCDRCommand struct {
	ExternalID         string
	StartedAt          time.Time
	EndedAt            time.Time
	Duration           int
	Caller             string
	Callee             string
	AudioURLs          []string
	ExternalDirection  string
	Direction          call_direction.CallDirection
	ExternalCallStatus string
	CallState          call_state.CallState
	ExternalSessionID  string
	ExtensionID        dot.ID
	HotlineID          dot.ID
	OwnerID            dot.ID
	UserID             dot.ID
	ConnectionID       dot.ID

	Result *CallLog `json:"-"`
}

func (h AggregateHandler) HandleCreateOrUpdateCallLogFromCDR(ctx context.Context, msg *CreateOrUpdateCallLogFromCDRCommand) (err error) {
	msg.Result, err = h.inner.CreateOrUpdateCallLogFromCDR(msg.GetArgs(ctx))
	return err
}

type CreateTenantCommand struct {
	OwnerID      dot.ID
	ConnectionID dot.ID

	Result *Tenant `json:"-"`
}

func (h AggregateHandler) HandleCreateTenant(ctx context.Context, msg *CreateTenantCommand) (err error) {
	msg.Result, err = h.inner.CreateTenant(msg.GetArgs(ctx))
	return err
}

type DeleteExtensionCommand struct {
	Id dot.ID

	Result struct {
	} `json:"-"`
}

func (h AggregateHandler) HandleDeleteExtension(ctx context.Context, msg *DeleteExtensionCommand) (err error) {
	return h.inner.DeleteExtension(msg.GetArgs(ctx))
}

type DeleteHotlineCommand struct {
	Id dot.ID

	Result struct {
	} `json:"-"`
}

func (h AggregateHandler) HandleDeleteHotline(ctx context.Context, msg *DeleteHotlineCommand) (err error) {
	return h.inner.DeleteHotline(msg.GetArgs(ctx))
}

type DeleteTenantCommand struct {
	Id dot.ID

	Result struct {
	} `json:"-"`
}

func (h AggregateHandler) HandleDeleteTenant(ctx context.Context, msg *DeleteTenantCommand) (err error) {
	return h.inner.DeleteTenant(msg.GetArgs(ctx))
}

type ExtendExtensionCommand struct {
	ExtensionID        dot.ID
	UserID             dot.ID
	AccountID          dot.ID
	SubscriptionID     dot.ID
	SubscriptionPlanID dot.ID
	PaymentMethod      payment_method.PaymentMethod

	Result *Extension `json:"-"`
}

func (h AggregateHandler) HandleExtendExtension(ctx context.Context, msg *ExtendExtensionCommand) (err error) {
	msg.Result, err = h.inner.ExtendExtension(msg.GetArgs(ctx))
	return err
}

type ImportExtensionsCommand struct {
	TenantID   dot.ID
	OwnerID    dot.ID
	AccountID  dot.ID
	Extensions []*ImportExtensionInfo

	Result struct {
	} `json:"-"`
}

func (h AggregateHandler) HandleImportExtensions(ctx context.Context, msg *ImportExtensionsCommand) (err error) {
	return h.inner.ImportExtensions(msg.GetArgs(ctx))
}

type RemoveHotlineOutOfTenantCommand struct {
	HotlineID dot.ID
	OwnerID   dot.ID

	Result struct {
	} `json:"-"`
}

func (h AggregateHandler) HandleRemoveHotlineOutOfTenant(ctx context.Context, msg *RemoveHotlineOutOfTenantCommand) (err error) {
	return h.inner.RemoveHotlineOutOfTenant(msg.GetArgs(ctx))
}

type RemoveUserOfExtensionCommand struct {
	AccountID   dot.ID
	ExtensionID dot.ID
	UserID      dot.ID

	Result int `json:"-"`
}

func (h AggregateHandler) HandleRemoveUserOfExtension(ctx context.Context, msg *RemoveUserOfExtensionCommand) (err error) {
	msg.Result, err = h.inner.RemoveUserOfExtension(msg.GetArgs(ctx))
	return err
}

type UpdateCallLogPostageCommand struct {
	ID                 dot.ID
	DurationForPostage int
	Postage            int

	Result struct {
	} `json:"-"`
}

func (h AggregateHandler) HandleUpdateCallLogPostage(ctx context.Context, msg *UpdateCallLogPostageCommand) (err error) {
	return h.inner.UpdateCallLogPostage(msg.GetArgs(ctx))
}

type UpdateExternalExtensionInfoCommand struct {
	ID                dot.ID
	HotlineID         dot.ID
	ExternalID        string
	ExtensionNumber   string
	ExtensionPassword string
	TenantDomain      string
	TenantID          dot.ID

	Result struct {
	} `json:"-"`
}

func (h AggregateHandler) HandleUpdateExternalExtensionInfo(ctx context.Context, msg *UpdateExternalExtensionInfoCommand) (err error) {
	return h.inner.UpdateExternalExtensionInfo(msg.GetArgs(ctx))
}

type UpdateHotlineInfoCommand struct {
	ID               dot.ID
	IsFreeCharge     dot.NullBool
	Name             string
	Description      string
	Status           status3.NullStatus
	TenantID         dot.ID
	ConnectionID     dot.ID
	ConnectionMethod connection_type.ConnectionMethod
	OwnerID          dot.ID
	Network          mobile_network.MobileNetwork

	Result struct {
	} `json:"-"`
}

func (h AggregateHandler) HandleUpdateHotlineInfo(ctx context.Context, msg *UpdateHotlineInfoCommand) (err error) {
	return h.inner.UpdateHotlineInfo(msg.GetArgs(ctx))
}

type GetCallLogQuery struct {
	ID dot.ID

	Result *CallLog `json:"-"`
}

func (h QueryServiceHandler) HandleGetCallLog(ctx context.Context, msg *GetCallLogQuery) (err error) {
	msg.Result, err = h.inner.GetCallLog(msg.GetArgs(ctx))
	return err
}

type GetCallLogByExternalIDQuery struct {
	ExternalID string

	Result *CallLog `json:"-"`
}

func (h QueryServiceHandler) HandleGetCallLogByExternalID(ctx context.Context, msg *GetCallLogByExternalIDQuery) (err error) {
	msg.Result, err = h.inner.GetCallLogByExternalID(msg.GetArgs(ctx))
	return err
}

type GetExtensionQuery struct {
	ID             dot.ID
	UserID         dot.ID
	AccountID      dot.ID
	HotlineID      dot.ID
	SubscriptionID dot.ID

	Result *Extension `json:"-"`
}

func (h QueryServiceHandler) HandleGetExtension(ctx context.Context, msg *GetExtensionQuery) (err error) {
	msg.Result, err = h.inner.GetExtension(msg.GetArgs(ctx))
	return err
}

type GetHotlineQuery struct {
	ID      dot.ID
	OwnerID dot.ID

	Result *Hotline `json:"-"`
}

func (h QueryServiceHandler) HandleGetHotline(ctx context.Context, msg *GetHotlineQuery) (err error) {
	msg.Result, err = h.inner.GetHotline(msg.GetArgs(ctx))
	return err
}

type GetHotlineByHotlineNumberQuery struct {
	Hotline string
	OwnerID dot.ID

	Result *Hotline `json:"-"`
}

func (h QueryServiceHandler) HandleGetHotlineByHotlineNumber(ctx context.Context, msg *GetHotlineByHotlineNumberQuery) (err error) {
	msg.Result, err = h.inner.GetHotlineByHotlineNumber(msg.GetArgs(ctx))
	return err
}

type GetPrivateExtensionNumberQuery struct {
	Result string `json:"-"`
}

func (h QueryServiceHandler) HandleGetPrivateExtensionNumber(ctx context.Context, msg *GetPrivateExtensionNumberQuery) (err error) {
	msg.Result, err = h.inner.GetPrivateExtensionNumber(msg.GetArgs(ctx))
	return err
}

type GetTenantByConnectionQuery struct {
	OwnerID      dot.ID
	ConnectionID dot.ID

	Result *Tenant `json:"-"`
}

func (h QueryServiceHandler) HandleGetTenantByConnection(ctx context.Context, msg *GetTenantByConnectionQuery) (err error) {
	msg.Result, err = h.inner.GetTenantByConnection(msg.GetArgs(ctx))
	return err
}

type GetTenantByIDQuery struct {
	ID dot.ID

	Result *Tenant `json:"-"`
}

func (h QueryServiceHandler) HandleGetTenantByID(ctx context.Context, msg *GetTenantByIDQuery) (err error) {
	msg.Result, err = h.inner.GetTenantByID(msg.GetArgs(ctx))
	return err
}

type ListBuiltinHotlinesQuery struct {
	Result []*Hotline `json:"-"`
}

func (h QueryServiceHandler) HandleListBuiltinHotlines(ctx context.Context, msg *ListBuiltinHotlinesQuery) (err error) {
	msg.Result, err = h.inner.ListBuiltinHotlines(msg.GetArgs(ctx))
	return err
}

type ListCallLogsQuery struct {
	HotlineIDs     []dot.ID
	ExtensionIDs   []dot.ID
	UserID         dot.ID
	OwnerID        dot.ID
	AccountID      dot.ID
	CallerOrCallee string
	Paging         meta.Paging

	Result *ListCallLogsResponse `json:"-"`
}

func (h QueryServiceHandler) HandleListCallLogs(ctx context.Context, msg *ListCallLogsQuery) (err error) {
	msg.Result, err = h.inner.ListCallLogs(msg.GetArgs(ctx))
	return err
}

type ListExtensionsQuery struct {
	AccountIDs       []dot.ID
	HotlineIDs       []dot.ID
	TenantID         dot.ID
	ExtensionNumbers []string

	Result []*Extension `json:"-"`
}

func (h QueryServiceHandler) HandleListExtensions(ctx context.Context, msg *ListExtensionsQuery) (err error) {
	msg.Result, err = h.inner.ListExtensions(msg.GetArgs(ctx))
	return err
}

type ListHotlinesQuery struct {
	OwnerID      dot.ID
	TenantID     dot.ID
	ConnectionID dot.ID

	Result []*Hotline `json:"-"`
}

func (h QueryServiceHandler) HandleListHotlines(ctx context.Context, msg *ListHotlinesQuery) (err error) {
	msg.Result, err = h.inner.ListHotlines(msg.GetArgs(ctx))
	return err
}

type ListTenantsQuery struct {
	OwnerID      dot.ID
	ConnectionID dot.ID
	Paging       meta.Paging

	Result *ListTenantsResponse `json:"-"`
}

func (h QueryServiceHandler) HandleListTenants(ctx context.Context, msg *ListTenantsQuery) (err error) {
	msg.Result, err = h.inner.ListTenants(msg.GetArgs(ctx))
	return err
}

// implement interfaces

func (q *ActivateTenantCommand) command()                {}
func (q *AssignUserToExtensionCommand) command()         {}
func (q *CreateCallLogCommand) command()                 {}
func (q *CreateExtensionCommand) command()               {}
func (q *CreateExtensionBySubscriptionCommand) command() {}
func (q *CreateHotlineCommand) command()                 {}
func (q *CreateOrUpdateCallLogFromCDRCommand) command()  {}
func (q *CreateTenantCommand) command()                  {}
func (q *DeleteExtensionCommand) command()               {}
func (q *DeleteHotlineCommand) command()                 {}
func (q *DeleteTenantCommand) command()                  {}
func (q *ExtendExtensionCommand) command()               {}
func (q *ImportExtensionsCommand) command()              {}
func (q *RemoveHotlineOutOfTenantCommand) command()      {}
func (q *RemoveUserOfExtensionCommand) command()         {}
func (q *UpdateCallLogPostageCommand) command()          {}
func (q *UpdateExternalExtensionInfoCommand) command()   {}
func (q *UpdateHotlineInfoCommand) command()             {}

func (q *GetCallLogQuery) query()                {}
func (q *GetCallLogByExternalIDQuery) query()    {}
func (q *GetExtensionQuery) query()              {}
func (q *GetHotlineQuery) query()                {}
func (q *GetHotlineByHotlineNumberQuery) query() {}
func (q *GetPrivateExtensionNumberQuery) query() {}
func (q *GetTenantByConnectionQuery) query()     {}
func (q *GetTenantByIDQuery) query()             {}
func (q *ListBuiltinHotlinesQuery) query()       {}
func (q *ListCallLogsQuery) query()              {}
func (q *ListExtensionsQuery) query()            {}
func (q *ListHotlinesQuery) query()              {}
func (q *ListTenantsQuery) query()               {}

// implement conversion

func (q *ActivateTenantCommand) GetArgs(ctx context.Context) (_ context.Context, _ *ActivateTenantArgs) {
	return ctx,
		&ActivateTenantArgs{
			OwnerID:            q.OwnerID,
			TenantID:           q.TenantID,
			HotlineID:          q.HotlineID,
			ConnectionID:       q.ConnectionID,
			ConnectionProvider: q.ConnectionProvider,
		}
}

func (q *ActivateTenantCommand) SetActivateTenantArgs(args *ActivateTenantArgs) {
	q.OwnerID = args.OwnerID
	q.TenantID = args.TenantID
	q.HotlineID = args.HotlineID
	q.ConnectionID = args.ConnectionID
	q.ConnectionProvider = args.ConnectionProvider
}

func (q *AssignUserToExtensionCommand) GetArgs(ctx context.Context) (_ context.Context, _ *AssignUserToExtensionArgs) {
	return ctx,
		&AssignUserToExtensionArgs{
			AccountID:   q.AccountID,
			UserID:      q.UserID,
			ExtensionID: q.ExtensionID,
		}
}

func (q *AssignUserToExtensionCommand) SetAssignUserToExtensionArgs(args *AssignUserToExtensionArgs) {
	q.AccountID = args.AccountID
	q.UserID = args.UserID
	q.ExtensionID = args.ExtensionID
}

func (q *CreateCallLogCommand) GetArgs(ctx context.Context) (_ context.Context, _ *CreateCallLogArgs) {
	return ctx,
		&CreateCallLogArgs{
			ExternalSessionID: q.ExternalSessionID,
			Direction:         q.Direction,
			Caller:            q.Caller,
			Callee:            q.Callee,
			ExtensionID:       q.ExtensionID,
			AccountID:         q.AccountID,
			OwnerID:           q.OwnerID,
			ContactID:         q.ContactID,
			CallState:         q.CallState,
			StartedAt:         q.StartedAt,
			EndedAt:           q.EndedAt,
		}
}

func (q *CreateCallLogCommand) SetCreateCallLogArgs(args *CreateCallLogArgs) {
	q.ExternalSessionID = args.ExternalSessionID
	q.Direction = args.Direction
	q.Caller = args.Caller
	q.Callee = args.Callee
	q.ExtensionID = args.ExtensionID
	q.AccountID = args.AccountID
	q.OwnerID = args.OwnerID
	q.ContactID = args.ContactID
	q.CallState = args.CallState
	q.StartedAt = args.StartedAt
	q.EndedAt = args.EndedAt
}

func (q *CreateExtensionCommand) GetArgs(ctx context.Context) (_ context.Context, _ *CreateExtensionArgs) {
	return ctx,
		&CreateExtensionArgs{
			ExtensionNumber:   q.ExtensionNumber,
			UserID:            q.UserID,
			AccountID:         q.AccountID,
			ExtensionPassword: q.ExtensionPassword,
			HotlineID:         q.HotlineID,
			OwnerID:           q.OwnerID,
			SubscriptionID:    q.SubscriptionID,
			ExpiresAt:         q.ExpiresAt,
		}
}

func (q *CreateExtensionCommand) SetCreateExtensionArgs(args *CreateExtensionArgs) {
	q.ExtensionNumber = args.ExtensionNumber
	q.UserID = args.UserID
	q.AccountID = args.AccountID
	q.ExtensionPassword = args.ExtensionPassword
	q.HotlineID = args.HotlineID
	q.OwnerID = args.OwnerID
	q.SubscriptionID = args.SubscriptionID
	q.ExpiresAt = args.ExpiresAt
}

func (q *CreateExtensionBySubscriptionCommand) GetArgs(ctx context.Context) (_ context.Context, _ *CreateExtenstionBySubscriptionArgs) {
	return ctx,
		&CreateExtenstionBySubscriptionArgs{
			SubscriptionID:     q.SubscriptionID,
			SubscriptionPlanID: q.SubscriptionPlanID,
			PaymentMethod:      q.PaymentMethod,
			AccountID:          q.AccountID,
			ExtensionNumber:    q.ExtensionNumber,
			UserID:             q.UserID,
			HotlineID:          q.HotlineID,
			OwnerID:            q.OwnerID,
		}
}

func (q *CreateExtensionBySubscriptionCommand) SetCreateExtenstionBySubscriptionArgs(args *CreateExtenstionBySubscriptionArgs) {
	q.SubscriptionID = args.SubscriptionID
	q.SubscriptionPlanID = args.SubscriptionPlanID
	q.PaymentMethod = args.PaymentMethod
	q.AccountID = args.AccountID
	q.ExtensionNumber = args.ExtensionNumber
	q.UserID = args.UserID
	q.HotlineID = args.HotlineID
	q.OwnerID = args.OwnerID
}

func (q *CreateHotlineCommand) GetArgs(ctx context.Context) (_ context.Context, _ *CreateHotlineArgs) {
	return ctx,
		&CreateHotlineArgs{
			OwnerID:      q.OwnerID,
			Name:         q.Name,
			Hotline:      q.Hotline,
			Network:      q.Network,
			ConnectionID: q.ConnectionID,
			Status:       q.Status,
			Description:  q.Description,
			IsFreeCharge: q.IsFreeCharge,
		}
}

func (q *CreateHotlineCommand) SetCreateHotlineArgs(args *CreateHotlineArgs) {
	q.OwnerID = args.OwnerID
	q.Name = args.Name
	q.Hotline = args.Hotline
	q.Network = args.Network
	q.ConnectionID = args.ConnectionID
	q.Status = args.Status
	q.Description = args.Description
	q.IsFreeCharge = args.IsFreeCharge
}

func (q *CreateOrUpdateCallLogFromCDRCommand) GetArgs(ctx context.Context) (_ context.Context, _ *CreateOrUpdateCallLogFromCDRArgs) {
	return ctx,
		&CreateOrUpdateCallLogFromCDRArgs{
			ExternalID:         q.ExternalID,
			StartedAt:          q.StartedAt,
			EndedAt:            q.EndedAt,
			Duration:           q.Duration,
			Caller:             q.Caller,
			Callee:             q.Callee,
			AudioURLs:          q.AudioURLs,
			ExternalDirection:  q.ExternalDirection,
			Direction:          q.Direction,
			ExternalCallStatus: q.ExternalCallStatus,
			CallState:          q.CallState,
			ExternalSessionID:  q.ExternalSessionID,
			ExtensionID:        q.ExtensionID,
			HotlineID:          q.HotlineID,
			OwnerID:            q.OwnerID,
			UserID:             q.UserID,
			ConnectionID:       q.ConnectionID,
		}
}

func (q *CreateOrUpdateCallLogFromCDRCommand) SetCreateOrUpdateCallLogFromCDRArgs(args *CreateOrUpdateCallLogFromCDRArgs) {
	q.ExternalID = args.ExternalID
	q.StartedAt = args.StartedAt
	q.EndedAt = args.EndedAt
	q.Duration = args.Duration
	q.Caller = args.Caller
	q.Callee = args.Callee
	q.AudioURLs = args.AudioURLs
	q.ExternalDirection = args.ExternalDirection
	q.Direction = args.Direction
	q.ExternalCallStatus = args.ExternalCallStatus
	q.CallState = args.CallState
	q.ExternalSessionID = args.ExternalSessionID
	q.ExtensionID = args.ExtensionID
	q.HotlineID = args.HotlineID
	q.OwnerID = args.OwnerID
	q.UserID = args.UserID
	q.ConnectionID = args.ConnectionID
}

func (q *CreateTenantCommand) GetArgs(ctx context.Context) (_ context.Context, _ *CreateTenantArgs) {
	return ctx,
		&CreateTenantArgs{
			OwnerID:      q.OwnerID,
			ConnectionID: q.ConnectionID,
		}
}

func (q *CreateTenantCommand) SetCreateTenantArgs(args *CreateTenantArgs) {
	q.OwnerID = args.OwnerID
	q.ConnectionID = args.ConnectionID
}

func (q *DeleteExtensionCommand) GetArgs(ctx context.Context) (_ context.Context, id dot.ID) {
	return ctx,
		q.Id
}

func (q *DeleteHotlineCommand) GetArgs(ctx context.Context) (_ context.Context, id dot.ID) {
	return ctx,
		q.Id
}

func (q *DeleteTenantCommand) GetArgs(ctx context.Context) (_ context.Context, id dot.ID) {
	return ctx,
		q.Id
}

func (q *ExtendExtensionCommand) GetArgs(ctx context.Context) (_ context.Context, _ *ExtendExtensionArgs) {
	return ctx,
		&ExtendExtensionArgs{
			ExtensionID:        q.ExtensionID,
			UserID:             q.UserID,
			AccountID:          q.AccountID,
			SubscriptionID:     q.SubscriptionID,
			SubscriptionPlanID: q.SubscriptionPlanID,
			PaymentMethod:      q.PaymentMethod,
		}
}

func (q *ExtendExtensionCommand) SetExtendExtensionArgs(args *ExtendExtensionArgs) {
	q.ExtensionID = args.ExtensionID
	q.UserID = args.UserID
	q.AccountID = args.AccountID
	q.SubscriptionID = args.SubscriptionID
	q.SubscriptionPlanID = args.SubscriptionPlanID
	q.PaymentMethod = args.PaymentMethod
}

func (q *ImportExtensionsCommand) GetArgs(ctx context.Context) (_ context.Context, _ *ImportExtensionsArgs) {
	return ctx,
		&ImportExtensionsArgs{
			TenantID:   q.TenantID,
			OwnerID:    q.OwnerID,
			AccountID:  q.AccountID,
			Extensions: q.Extensions,
		}
}

func (q *ImportExtensionsCommand) SetImportExtensionsArgs(args *ImportExtensionsArgs) {
	q.TenantID = args.TenantID
	q.OwnerID = args.OwnerID
	q.AccountID = args.AccountID
	q.Extensions = args.Extensions
}

func (q *RemoveHotlineOutOfTenantCommand) GetArgs(ctx context.Context) (_ context.Context, _ *RemoveHotlineOutOfTenantArgs) {
	return ctx,
		&RemoveHotlineOutOfTenantArgs{
			HotlineID: q.HotlineID,
			OwnerID:   q.OwnerID,
		}
}

func (q *RemoveHotlineOutOfTenantCommand) SetRemoveHotlineOutOfTenantArgs(args *RemoveHotlineOutOfTenantArgs) {
	q.HotlineID = args.HotlineID
	q.OwnerID = args.OwnerID
}

func (q *RemoveUserOfExtensionCommand) GetArgs(ctx context.Context) (_ context.Context, _ *RemoveUserOfExtensionArgs) {
	return ctx,
		&RemoveUserOfExtensionArgs{
			AccountID:   q.AccountID,
			ExtensionID: q.ExtensionID,
			UserID:      q.UserID,
		}
}

func (q *RemoveUserOfExtensionCommand) SetRemoveUserOfExtensionArgs(args *RemoveUserOfExtensionArgs) {
	q.AccountID = args.AccountID
	q.ExtensionID = args.ExtensionID
	q.UserID = args.UserID
}

func (q *UpdateCallLogPostageCommand) GetArgs(ctx context.Context) (_ context.Context, _ *UpdateCallLogPostageArgs) {
	return ctx,
		&UpdateCallLogPostageArgs{
			ID:                 q.ID,
			DurationForPostage: q.DurationForPostage,
			Postage:            q.Postage,
		}
}

func (q *UpdateCallLogPostageCommand) SetUpdateCallLogPostageArgs(args *UpdateCallLogPostageArgs) {
	q.ID = args.ID
	q.DurationForPostage = args.DurationForPostage
	q.Postage = args.Postage
}

func (q *UpdateExternalExtensionInfoCommand) GetArgs(ctx context.Context) (_ context.Context, _ *UpdateExternalExtensionInfoArgs) {
	return ctx,
		&UpdateExternalExtensionInfoArgs{
			ID:                q.ID,
			HotlineID:         q.HotlineID,
			ExternalID:        q.ExternalID,
			ExtensionNumber:   q.ExtensionNumber,
			ExtensionPassword: q.ExtensionPassword,
			TenantDomain:      q.TenantDomain,
			TenantID:          q.TenantID,
		}
}

func (q *UpdateExternalExtensionInfoCommand) SetUpdateExternalExtensionInfoArgs(args *UpdateExternalExtensionInfoArgs) {
	q.ID = args.ID
	q.HotlineID = args.HotlineID
	q.ExternalID = args.ExternalID
	q.ExtensionNumber = args.ExtensionNumber
	q.ExtensionPassword = args.ExtensionPassword
	q.TenantDomain = args.TenantDomain
	q.TenantID = args.TenantID
}

func (q *UpdateHotlineInfoCommand) GetArgs(ctx context.Context) (_ context.Context, _ *UpdateHotlineInfoArgs) {
	return ctx,
		&UpdateHotlineInfoArgs{
			ID:               q.ID,
			IsFreeCharge:     q.IsFreeCharge,
			Name:             q.Name,
			Description:      q.Description,
			Status:           q.Status,
			TenantID:         q.TenantID,
			ConnectionID:     q.ConnectionID,
			ConnectionMethod: q.ConnectionMethod,
			OwnerID:          q.OwnerID,
			Network:          q.Network,
		}
}

func (q *UpdateHotlineInfoCommand) SetUpdateHotlineInfoArgs(args *UpdateHotlineInfoArgs) {
	q.ID = args.ID
	q.IsFreeCharge = args.IsFreeCharge
	q.Name = args.Name
	q.Description = args.Description
	q.Status = args.Status
	q.TenantID = args.TenantID
	q.ConnectionID = args.ConnectionID
	q.ConnectionMethod = args.ConnectionMethod
	q.OwnerID = args.OwnerID
	q.Network = args.Network
}

func (q *GetCallLogQuery) GetArgs(ctx context.Context) (_ context.Context, ID dot.ID) {
	return ctx,
		q.ID
}

func (q *GetCallLogByExternalIDQuery) GetArgs(ctx context.Context) (_ context.Context, _ *GetCallLogByExternalIDArgs) {
	return ctx,
		&GetCallLogByExternalIDArgs{
			ExternalID: q.ExternalID,
		}
}

func (q *GetCallLogByExternalIDQuery) SetGetCallLogByExternalIDArgs(args *GetCallLogByExternalIDArgs) {
	q.ExternalID = args.ExternalID
}

func (q *GetExtensionQuery) GetArgs(ctx context.Context) (_ context.Context, _ *GetExtensionArgs) {
	return ctx,
		&GetExtensionArgs{
			ID:             q.ID,
			UserID:         q.UserID,
			AccountID:      q.AccountID,
			HotlineID:      q.HotlineID,
			SubscriptionID: q.SubscriptionID,
		}
}

func (q *GetExtensionQuery) SetGetExtensionArgs(args *GetExtensionArgs) {
	q.ID = args.ID
	q.UserID = args.UserID
	q.AccountID = args.AccountID
	q.HotlineID = args.HotlineID
	q.SubscriptionID = args.SubscriptionID
}

func (q *GetHotlineQuery) GetArgs(ctx context.Context) (_ context.Context, _ *GetHotlineArgs) {
	return ctx,
		&GetHotlineArgs{
			ID:      q.ID,
			OwnerID: q.OwnerID,
		}
}

func (q *GetHotlineQuery) SetGetHotlineArgs(args *GetHotlineArgs) {
	q.ID = args.ID
	q.OwnerID = args.OwnerID
}

func (q *GetHotlineByHotlineNumberQuery) GetArgs(ctx context.Context) (_ context.Context, _ *GetHotlineByHotlineNumberArgs) {
	return ctx,
		&GetHotlineByHotlineNumberArgs{
			Hotline: q.Hotline,
			OwnerID: q.OwnerID,
		}
}

func (q *GetHotlineByHotlineNumberQuery) SetGetHotlineByHotlineNumberArgs(args *GetHotlineByHotlineNumberArgs) {
	q.Hotline = args.Hotline
	q.OwnerID = args.OwnerID
}

func (q *GetPrivateExtensionNumberQuery) GetArgs(ctx context.Context) (_ context.Context, _ *common.Empty) {
	return ctx,
		&common.Empty{}
}

func (q *GetPrivateExtensionNumberQuery) SetEmpty(args *common.Empty) {
}

func (q *GetTenantByConnectionQuery) GetArgs(ctx context.Context) (_ context.Context, _ *GetTenantByConnectionArgs) {
	return ctx,
		&GetTenantByConnectionArgs{
			OwnerID:      q.OwnerID,
			ConnectionID: q.ConnectionID,
		}
}

func (q *GetTenantByConnectionQuery) SetGetTenantByConnectionArgs(args *GetTenantByConnectionArgs) {
	q.OwnerID = args.OwnerID
	q.ConnectionID = args.ConnectionID
}

func (q *GetTenantByIDQuery) GetArgs(ctx context.Context) (_ context.Context, ID dot.ID) {
	return ctx,
		q.ID
}

func (q *ListBuiltinHotlinesQuery) GetArgs(ctx context.Context) (_ context.Context, _ *common.Empty) {
	return ctx,
		&common.Empty{}
}

func (q *ListBuiltinHotlinesQuery) SetEmpty(args *common.Empty) {
}

func (q *ListCallLogsQuery) GetArgs(ctx context.Context) (_ context.Context, _ *ListCallLogsArgs) {
	return ctx,
		&ListCallLogsArgs{
			HotlineIDs:     q.HotlineIDs,
			ExtensionIDs:   q.ExtensionIDs,
			UserID:         q.UserID,
			OwnerID:        q.OwnerID,
			AccountID:      q.AccountID,
			CallerOrCallee: q.CallerOrCallee,
			Paging:         q.Paging,
		}
}

func (q *ListCallLogsQuery) SetListCallLogsArgs(args *ListCallLogsArgs) {
	q.HotlineIDs = args.HotlineIDs
	q.ExtensionIDs = args.ExtensionIDs
	q.UserID = args.UserID
	q.OwnerID = args.OwnerID
	q.AccountID = args.AccountID
	q.CallerOrCallee = args.CallerOrCallee
	q.Paging = args.Paging
}

func (q *ListExtensionsQuery) GetArgs(ctx context.Context) (_ context.Context, _ *ListExtensionsArgs) {
	return ctx,
		&ListExtensionsArgs{
			AccountIDs:       q.AccountIDs,
			HotlineIDs:       q.HotlineIDs,
			TenantID:         q.TenantID,
			ExtensionNumbers: q.ExtensionNumbers,
		}
}

func (q *ListExtensionsQuery) SetListExtensionsArgs(args *ListExtensionsArgs) {
	q.AccountIDs = args.AccountIDs
	q.HotlineIDs = args.HotlineIDs
	q.TenantID = args.TenantID
	q.ExtensionNumbers = args.ExtensionNumbers
}

func (q *ListHotlinesQuery) GetArgs(ctx context.Context) (_ context.Context, _ *ListHotlinesArgs) {
	return ctx,
		&ListHotlinesArgs{
			OwnerID:      q.OwnerID,
			TenantID:     q.TenantID,
			ConnectionID: q.ConnectionID,
		}
}

func (q *ListHotlinesQuery) SetListHotlinesArgs(args *ListHotlinesArgs) {
	q.OwnerID = args.OwnerID
	q.TenantID = args.TenantID
	q.ConnectionID = args.ConnectionID
}

func (q *ListTenantsQuery) GetArgs(ctx context.Context) (_ context.Context, _ *ListTenantsArgs) {
	return ctx,
		&ListTenantsArgs{
			OwnerID:      q.OwnerID,
			ConnectionID: q.ConnectionID,
			Paging:       q.Paging,
		}
}

func (q *ListTenantsQuery) SetListTenantsArgs(args *ListTenantsArgs) {
	q.OwnerID = args.OwnerID
	q.ConnectionID = args.ConnectionID
	q.Paging = args.Paging
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
	b.AddHandler(h.HandleActivateTenant)
	b.AddHandler(h.HandleAssignUserToExtension)
	b.AddHandler(h.HandleCreateCallLog)
	b.AddHandler(h.HandleCreateExtension)
	b.AddHandler(h.HandleCreateExtensionBySubscription)
	b.AddHandler(h.HandleCreateHotline)
	b.AddHandler(h.HandleCreateOrUpdateCallLogFromCDR)
	b.AddHandler(h.HandleCreateTenant)
	b.AddHandler(h.HandleDeleteExtension)
	b.AddHandler(h.HandleDeleteHotline)
	b.AddHandler(h.HandleDeleteTenant)
	b.AddHandler(h.HandleExtendExtension)
	b.AddHandler(h.HandleImportExtensions)
	b.AddHandler(h.HandleRemoveHotlineOutOfTenant)
	b.AddHandler(h.HandleRemoveUserOfExtension)
	b.AddHandler(h.HandleUpdateCallLogPostage)
	b.AddHandler(h.HandleUpdateExternalExtensionInfo)
	b.AddHandler(h.HandleUpdateHotlineInfo)
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
	b.AddHandler(h.HandleGetCallLog)
	b.AddHandler(h.HandleGetCallLogByExternalID)
	b.AddHandler(h.HandleGetExtension)
	b.AddHandler(h.HandleGetHotline)
	b.AddHandler(h.HandleGetHotlineByHotlineNumber)
	b.AddHandler(h.HandleGetPrivateExtensionNumber)
	b.AddHandler(h.HandleGetTenantByConnection)
	b.AddHandler(h.HandleGetTenantByID)
	b.AddHandler(h.HandleListBuiltinHotlines)
	b.AddHandler(h.HandleListCallLogs)
	b.AddHandler(h.HandleListExtensions)
	b.AddHandler(h.HandleListHotlines)
	b.AddHandler(h.HandleListTenants)
	return QueryBus{b}
}