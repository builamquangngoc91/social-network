package etelecom

import (
	"context"
	"strconv"
	"time"

	"o.o/api/etelecom"
	"o.o/api/etelecom/summary"
	"o.o/api/main/authorization"
	"o.o/api/main/connectioning"
	"o.o/api/main/identity"
	etelecomapi "o.o/api/top/int/etelecom"
	etelecomtypes "o.o/api/top/int/etelecom/types"
	pbcm "o.o/api/top/types/common"
	"o.o/api/top/types/etc/status3"
	cm "o.o/backend/pkg/common"
	"o.o/backend/pkg/common/apifw/cmapi"
	"o.o/backend/pkg/common/validate"
	"o.o/backend/pkg/etop/api/convertpb"
	"o.o/backend/pkg/etop/authorize/session"
	"o.o/capi/dot"
)

type EtelecomService struct {
	session.Session

	EtelecomAggr    etelecom.CommandBus
	EtelecomQuery   etelecom.QueryBus
	SummaryQuery    summary.QueryBus
	IdentityAggr    identity.CommandBus
	IdentityQuery   identity.QueryBus
	ConnectionQuery connectioning.QueryBus
}

func (s *EtelecomService) Clone() etelecomapi.EtelecomService {
	res := *s
	return &res
}

func (s *EtelecomService) AssignUserToExtension(ctx context.Context, r *etelecomapi.AssignUserToExtensionRequest) (*pbcm.UpdatedResponse, error) {
	cmd := &etelecom.AssignUserToExtensionCommand{
		AccountID:   s.SS.Shop().ID,
		UserID:      r.UserID,
		ExtensionID: r.ExtensionID,
	}
	if err := s.EtelecomAggr.Dispatch(ctx, cmd); err != nil {
		return nil, err
	}

	return &pbcm.UpdatedResponse{
		Updated: 1,
	}, nil
}

func (s *EtelecomService) RemoveUserOfExtension(ctx context.Context, r *etelecomapi.RemoveUserOfExtensionRequest) (*pbcm.UpdatedResponse, error) {
	cmd := &etelecom.RemoveUserOfExtensionCommand{
		AccountID:   s.SS.Shop().ID,
		ExtensionID: r.ExtensionID,
		UserID:      r.UserID,
	}
	if err := s.EtelecomAggr.Dispatch(ctx, cmd); err != nil {
		return nil, err
	}
	return &pbcm.UpdatedResponse{
		Updated: cmd.Result,
	}, nil
}

func (s *EtelecomService) GetExtensions(ctx context.Context, r *etelecomtypes.GetExtensionsRequest) (*etelecomtypes.GetExtensionsResponse, error) {
	query := &etelecom.ListExtensionsQuery{
		AccountIDs: []dot.ID{s.SS.Shop().ID},
	}
	if r.Filter != nil {
		if r.Filter.HotlineID != 0 {
			query.HotlineIDs = []dot.ID{r.Filter.HotlineID}
		}
	}

	if err := s.EtelecomQuery.Dispatch(ctx, query); err != nil {
		return nil, err
	}
	res := Convert_etelecom_Extensions_etelecomtypes_Extensions(query.Result)

	// censor extension password
	for _, ext := range res {
		if ext.UserID != s.SS.User().ID ||
			(!ext.ExpiresAt.IsZero() && ext.ExpiresAt.Sub(time.Now()) <= 0) {
			ext.ExtensionPassword = ""
		}
	}
	return &etelecomtypes.GetExtensionsResponse{Extensions: res}, nil
}

func (s *EtelecomService) CreateExtension(ctx context.Context, r *etelecomtypes.CreateExtensionRequest) (*etelecomtypes.Extension, error) {
	extNumber := ""
	if r.ExtensionNumber != 0 {
		extNumber = strconv.Itoa(r.ExtensionNumber)
	}
	cmd := &etelecom.CreateExtensionCommand{
		UserID:          r.UserID,
		AccountID:       s.SS.Shop().ID,
		HotlineID:       r.HotlineID,
		OwnerID:         s.SS.Shop().OwnerID,
		ExtensionNumber: extNumber,
	}
	if err := s.EtelecomAggr.Dispatch(ctx, cmd); err != nil {
		return nil, err
	}
	var res etelecomtypes.Extension
	Convert_etelecom_Extension_etelecomtypes_Extension(cmd.Result, &res)
	return &res, nil
}

func (s *EtelecomService) CreateExtensionBySubscription(ctx context.Context, r *etelecomtypes.CreateExtensionBySubscriptionRequest) (*etelecomtypes.Extension, error) {
	extNumber := ""
	if r.ExtensionNumber != 0 {
		extNumber = strconv.Itoa(r.ExtensionNumber)
	}
	cmd := &etelecom.CreateExtensionBySubscriptionCommand{
		SubscriptionID:     r.SubscriptionID,
		SubscriptionPlanID: r.SubscriptionPlanID,
		PaymentMethod:      r.PaymentMethod,
		AccountID:          s.SS.Shop().ID,
		UserID:             r.UserID,
		HotlineID:          r.HotlineID,
		OwnerID:            s.SS.Shop().OwnerID,
		ExtensionNumber:    extNumber,
	}
	if err := s.EtelecomAggr.Dispatch(ctx, cmd); err != nil {
		return nil, err
	}
	var res etelecomtypes.Extension
	Convert_etelecom_Extension_etelecomtypes_Extension(cmd.Result, &res)
	return &res, nil
}

func (s *EtelecomService) ExtendExtension(ctx context.Context, r *etelecomtypes.ExtendExtensionRequest) (*etelecomtypes.Extension, error) {
	cmd := &etelecom.ExtendExtensionCommand{
		ExtensionID:        r.ExtensionID,
		UserID:             r.UserID,
		AccountID:          s.SS.Shop().ID,
		SubscriptionID:     r.SubscriptionID,
		SubscriptionPlanID: r.SubscriptionPlanID,
		PaymentMethod:      r.PaymentMethod,
	}
	if err := s.EtelecomAggr.Dispatch(ctx, cmd); err != nil {
		return nil, err
	}
	var res etelecomtypes.Extension
	Convert_etelecom_Extension_etelecomtypes_Extension(cmd.Result, &res)
	return &res, nil
}

func (s *EtelecomService) GetHotlines(ctx context.Context, _ *pbcm.Empty) (*etelecomtypes.GetHotLinesResponse, error) {
	// list all hotline builtin
	queryBuiltinHotlines := &etelecom.ListBuiltinHotlinesQuery{}
	if err := s.EtelecomQuery.Dispatch(ctx, queryBuiltinHotlines); err != nil {
		return nil, err
	}
	builtinHotlines := queryBuiltinHotlines.Result

	query := &etelecom.ListHotlinesQuery{
		OwnerID: s.SS.Shop().OwnerID,
	}
	if err := s.EtelecomQuery.Dispatch(ctx, query); err != nil {
		return nil, err
	}
	hotlines := append(builtinHotlines, query.Result...)

	res := Convert_etelecom_Hotlines_etelecomtypes_Hotlines(hotlines)
	return &etelecomtypes.GetHotLinesResponse{Hotlines: res}, nil
}

func (s *EtelecomService) GetCallLogs(ctx context.Context, r *etelecomtypes.GetCallLogsRequest) (*etelecomtypes.GetCallLogsResponse, error) {
	paging, err := cmapi.CMCursorPaging(r.Paging)
	if err != nil {
		return nil, err
	}
	query := &etelecom.ListCallLogsQuery{
		AccountID: s.SS.Shop().ID,
		Paging:    *paging,
		OwnerID:   s.SS.Shop().OwnerID,
	}

	if r.Filter != nil {
		query.CallerOrCallee = r.Filter.CallNumber
		query.HotlineIDs = r.Filter.HotlineIDs
		query.ExtensionIDs = r.Filter.ExtensionIDs
		query.UserID = r.Filter.UserID
	}

	if err = s.EtelecomQuery.Dispatch(ctx, query); err != nil {
		return nil, err
	}
	res := Convert_etelecom_CallLogs_etelecomtypes_CallLogs(query.Result.CallLogs)
	return &etelecomtypes.GetCallLogsResponse{
		CallLogs: res,
		Paging:   cmapi.PbCursorPageInfo(paging, &query.Result.Paging),
	}, nil
}

func (s *EtelecomService) CreateCallLog(ctx context.Context, r *etelecomapi.CreateCallLogRequest) (*etelecomtypes.CallLog, error) {
	cmd := &etelecom.CreateCallLogCommand{
		ExternalSessionID: r.ExternalSessionID,
		Direction:         r.Direction,
		Caller:            r.Caller,
		Callee:            r.Callee,
		ExtensionID:       r.ExtensionID,
		AccountID:         s.SS.Shop().ID,
		OwnerID:           s.SS.Shop().OwnerID,
		ContactID:         r.ContactID,
		CallState:         r.CallState,
		StartedAt:         r.StartedAt,
		EndedAt:           r.EndedAt,
	}
	if err := s.EtelecomAggr.Dispatch(ctx, cmd); err != nil {
		return nil, err
	}
	res := Convert_etelecom_CallLog_etelecomtypes_CallLog(cmd.Result, nil)
	return res, nil
}

func (s *EtelecomService) SummaryEtelecom(
	ctx context.Context, req *etelecomapi.SummaryEtelecomRequest,
) (*etelecomapi.SummaryEtelecomResponse, error) {
	dateFrom, dateTo, err := cm.ParseDateFromTo(req.DateFrom, req.DateTo)
	if err != nil {
		return nil, err
	}

	if dateTo.Before(dateFrom) {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "date_to must be after date_from")
	}

	query := &summary.SummaryQuery{
		ShopID:   s.SS.Shop().ID,
		DateFrom: dateFrom,
		DateTo:   dateTo,
	}
	if err = s.SummaryQuery.Dispatch(ctx, query); err != nil {
		return nil, err
	}

	return &etelecomapi.SummaryEtelecomResponse{
		Tables: convertpb.PbSummaryTablesNew(query.Result.ListTable),
	}, nil
}

func (s *EtelecomService) CreateUserAndAssignExtension(ctx context.Context, r *etelecomapi.CreateUserAndAssignExtensionRequest) (*pbcm.MessageResponse, error) {
	phoneNorm, ok := validate.NormalizePhone(r.Phone)
	if !ok {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Số điện thoại không hợp lệ")
	}
	if r.FullName == "" {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Vui lòng điền họ tên")
	}
	if r.Password == "" {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Vui lòng điền password")
	}
	if r.HotlineID == 0 {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Vui lòng chọn hotline")
	}
	phone := phoneNorm.String()

	// register user if needed
	cmd := &identity.RegisterSimplifyCommand{
		Phone:    phone,
		FullName: r.FullName,
		Password: r.Password,
	}
	if err := s.IdentityAggr.Dispatch(ctx, cmd); err != nil {
		return nil, err
	}

	// get user
	query := &identity.GetUserByPhoneOrEmailQuery{
		Phone: phone,
	}
	if err := s.IdentityQuery.Dispatch(ctx, query); err != nil {
		return nil, err
	}
	if err := s.createEtelecomAccountUserAndAddRoleCS(ctx, query.Result.ID, s.SS.Shop().ID); err != nil {
		return nil, err
	}

	// create & assign extension for this user
	createExtRequest := &etelecomtypes.CreateExtensionRequest{
		UserID:    query.Result.ID,
		HotlineID: r.HotlineID,
	}
	_, err := s.CreateExtension(ctx, createExtRequest)
	if err != nil {
		return nil, err
	}
	return &pbcm.MessageResponse{
		Code: "OK",
		Msg:  "Tạo người dùng và gán extension thành công",
	}, nil
}

func (s *EtelecomService) createEtelecomAccountUserAndAddRoleCS(ctx context.Context, userID, accountID dot.ID) error {
	query := &identity.GetAccountUserQuery{
		UserID:    userID,
		AccountID: accountID,
	}
	err := s.IdentityQuery.Dispatch(ctx, query)
	switch cm.ErrorCode(err) {
	case cm.NotFound:
		// create new
		cmd := &identity.CreateAccountUserCommand{
			AccountID: accountID,
			UserID:    userID,
			Status:    status3.P,
			Permission: identity.Permission{
				Roles: []string{authorization.RoleTelecomCustomerService.String()},
			},
		}
		return s.IdentityAggr.Dispatch(ctx, cmd)

	case cm.NoError:
		// continue
		// check if it has role EtelecomCS => add if needed
		accountUser := query.Result

		if cm.StringsContain(accountUser.Roles, authorization.RoleShopOwner.String()) ||
			cm.StringsContain(accountUser.Roles, authorization.RoleTelecomCustomerService.String()) {
			return nil
		}
		cmd := &identity.UpdateAccountUserPermissionCommand{
			UserID:    userID,
			AccountID: accountID,
			Permission: identity.Permission{
				Roles: append(accountUser.Roles, authorization.RoleTelecomCustomerService.String()),
			},
		}
		return s.IdentityAggr.Dispatch(ctx, cmd)

	default:
		return err
	}
}

func (s *EtelecomService) CreateTenant(ctx context.Context, r *etelecomapi.CreateTenantRequest) (*etelecomtypes.Tenant, error) {
	cmd := &etelecom.CreateTenantCommand{
		OwnerID:      s.SS.Shop().OwnerID,
		ConnectionID: r.ConnectionID,
	}
	if cmd.ConnectionID == 0 {
		cmd.ConnectionID = connectioning.DefaultDirectPortsipConnectionID
	}
	if err := s.EtelecomAggr.Dispatch(ctx, cmd); err != nil {
		return nil, err
	}
	res := Convert_etelecom_Tenant_etelecomtypes_Tenant(cmd.Result, nil)
	return res, nil
}

func (s *EtelecomService) GetTenant(ctx context.Context, r *etelecomapi.GetTenantRequest) (*etelecomtypes.Tenant, error) {
	query := &etelecom.GetTenantByConnectionQuery{
		OwnerID:      s.SS.Shop().OwnerID,
		ConnectionID: r.ConnectionID,
	}
	if query.ConnectionID == 0 {
		query.ConnectionID = connectioning.DefaultDirectPortsipConnectionID
	}

	if err := s.EtelecomQuery.Dispatch(ctx, query); err != nil {
		return nil, err
	}
	res := Convert_etelecom_Tenant_etelecomtypes_Tenant(query.Result, nil)
	return res, nil
}