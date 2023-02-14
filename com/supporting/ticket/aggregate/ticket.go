package aggregate

import (
	"context"
	"fmt"
	"time"

	"o.o/api/main/connectioning"
	"o.o/api/main/contact"
	"o.o/api/main/identity"
	"o.o/api/main/moneytx"
	"o.o/api/main/ordering"
	"o.o/api/main/shipping"
	"o.o/api/supporting/ticket"
	pbcm "o.o/api/top/types/common"
	"o.o/api/top/types/etc/status5"
	"o.o/api/top/types/etc/ticket/ticket_ref_type"
	"o.o/api/top/types/etc/ticket/ticket_state"
	"o.o/api/top/types/etc/ticket/ticket_type"
	com "o.o/backend/com/main"
	"o.o/backend/com/supporting/ticket/model"
	"o.o/backend/com/supporting/ticket/provider"
	"o.o/backend/com/supporting/ticket/sqlstore"
	cm "o.o/backend/pkg/common"
	"o.o/backend/pkg/common/apifw/whitelabel/wl"
	"o.o/backend/pkg/common/bus"
	"o.o/backend/pkg/common/redis"
	"o.o/backend/pkg/common/sql/cmsql"
	"o.o/capi"
	"o.o/capi/dot"
)

const (
	ticketLabelsVersion = "v1.6"
)

var _ ticket.Aggregate = &TicketAggregate{}

type TicketAggregate struct {
	TicketStore                         sqlstore.TicketStoreFactory
	TicketCommentStore                  sqlstore.TicketCommentStoreFactory
	TicketLabelStore                    sqlstore.TicketLabelStoreFactory
	TicketLabelExternalStore            sqlstore.TicketLabelExternalStoreFactory
	TicketLabelTicketLabelExternalStore sqlstore.TicketLabelTicketLabelExternalsStoreFactory
	EventBus                            capi.EventBus
	db                                  *cmsql.Database
	MoneyTxQuery                        moneytx.QueryBus
	ShippingQuery                       shipping.QueryBus
	OrderQuery                          ordering.QueryBus
	IdentityQuery                       identity.QueryBus
	TicketManager                       *provider.TicketManager
	ConnectionQuery                     connectioning.QueryBus
	ContactQuery                        contact.QueryBus
	RedisStore                          redis.Store
}

func NewTicketAggregate(
	eventBus capi.EventBus,
	db com.MainDB,
	moneyTxQ moneytx.QueryBus,
	shippingQ shipping.QueryBus,
	orderQ ordering.QueryBus,
	identityQ identity.QueryBus,
	ticketManager *provider.TicketManager,
	connectionQ connectioning.QueryBus,
	contactQ contact.QueryBus,
	redisStore redis.Store,
) *TicketAggregate {
	return &TicketAggregate{
		TicketStore:                         sqlstore.NewTicketStore(db),
		TicketCommentStore:                  sqlstore.NewTicketCommentStore(db),
		TicketLabelStore:                    sqlstore.NewTicketLabelStore(db),
		TicketLabelExternalStore:            sqlstore.NewTicketLabelExternalStore(db),
		TicketLabelTicketLabelExternalStore: sqlstore.NewTicketLabelTicketLabelExternalStore(db),
		MoneyTxQuery:                        moneyTxQ,
		EventBus:                            eventBus,
		ShippingQuery:                       shippingQ,
		db:                                  db,
		OrderQuery:                          orderQ,
		IdentityQuery:                       identityQ,
		ConnectionQuery:                     connectionQ,
		TicketManager:                       ticketManager,
		ContactQuery:                        contactQ,
		RedisStore:                          redisStore,
	}
}

func TicketAggregateMessageBus(q *TicketAggregate) ticket.CommandBus {
	b := bus.New()
	return ticket.NewAggregateHandler(q).RegisterHandlers(b)
}

func (a *TicketAggregate) CreateTicket(ctx context.Context, args *ticket.CreateTicketArgs) (_ *ticket.Ticket, err error) {
	if args.AccountID == 0 {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Missing AccountID")
	}
	if args.Type != ticket_type.System && args.Type != ticket_type.Internal {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Unsupported type %v", args.Type)
	}

	ticketCore := &ticket.Ticket{}
	if err := scheme.Convert(args, ticketCore); err != nil {
		return nil, err
	}

	if len(args.LabelIDs) != 0 {
		var labels []*ticket.TicketLabel
		labels, err = a.listTicketLabels(ctx, listTicketLabelsArgs{
			Type:   args.Type,
			ShopID: args.AccountID,
		})
		if err != nil {
			return nil, err
		}

		// get all father label_ids of all labels
		for _, labelID := range args.LabelIDs {
			listLabelIDs := getListLabelFatherID(labelID, labels)
			if len(listLabelIDs) == 0 {
				return nil, cm.Errorf(cm.InvalidArgument, nil, "Label không tồn tại (id = %d)", labelID)
			}
			for _, _labelID := range listLabelIDs {
				if !cm.IDsContain(ticketCore.LabelIDs, _labelID) {
					ticketCore.LabelIDs = append(ticketCore.LabelIDs, _labelID)
				}
			}
		}
	}

	// get reference ticket
	if ticketCore.RefTicketID.Valid {
		if _, err := a.TicketStore(ctx).ID(ticketCore.RefTicketID.ID).GetTicket(); err != nil {
			return nil, err
		}
	}

	// check reference code
	var refCode = ""
	var connectionID dot.ID
	if args.RefID != 0 {
		switch args.RefType {
		case ticket_ref_type.FFM:
			getFfmQuery := &shipping.GetFulfillmentByIDOrShippingCodeQuery{
				ID: args.RefID,
			}
			if err := a.ShippingQuery.Dispatch(ctx, getFfmQuery); err != nil {
				if cm.ErrorCode(err) == cm.NotFound {
					return nil, cm.Errorf(cm.NotFound, err, "Không tìm thấy đơn giao hàng")
				}
				return nil, err
			}
			refCode = getFfmQuery.Result.ShippingCode
			connectionID = getFfmQuery.Result.ConnectionID
		case ticket_ref_type.MoneyTransaction:
			// type system
			getMoneyTxQuery := &moneytx.GetMoneyTxShippingByIDQuery{
				MoneyTxShippingID: args.RefID,
				ShopID:            args.AccountID,
			}
			if err := a.MoneyTxQuery.Dispatch(ctx, getMoneyTxQuery); err != nil {
				if cm.ErrorCode(err) == cm.NotFound {
					return nil, cm.Errorf(cm.NotFound, err, "Không tìm thấy phiên chuyển tiền")
				}
				return nil, err
			}
			refCode = getMoneyTxQuery.Result.Code
		case ticket_ref_type.OrderTrading:
			// type system
			getOrderQuery := &ordering.GetOrderByIDQuery{
				ID: args.RefID,
			}
			err := a.OrderQuery.Dispatch(ctx, getOrderQuery)
			if err != nil {
				if cm.ErrorCode(err) == cm.NotFound {
					return nil, cm.Errorf(cm.NotFound, err, "Không tìm thấy đơn hàng")
				}
				return nil, err
			}
			refCode = getOrderQuery.Result.Code
		case ticket_ref_type.Contact:
			getContactQuery := &contact.GetContactByIDQuery{
				ID:     args.RefID,
				ShopID: args.AccountID,
			}
			if err := a.ContactQuery.Dispatch(ctx, getContactQuery); err != nil {
				if cm.ErrorCode(err) == cm.NotFound {
					return nil, cm.Errorf(cm.NotFound, err, "Không tìm thấy liên lạc")
				}
				return nil, err
			}

		default:
			//no-op(other)
		}

		// check ref_code
		if args.RefCode != "" && args.RefCode != refCode {
			return nil, cm.Errorf(cm.NotFound, nil, "ref_code không đúng")
		}
		ticketCore.RefCode = refCode
		ticketCore.ConnectionID = connectionID
	}

	if err = a.TicketStore(ctx).Create(ticketCore); err != nil {
		return nil, err
	}
	return a.TicketStore(ctx).ID(ticketCore.ID).GetTicket()
}

func (a *TicketAggregate) UpdateTicketInfo(ctx context.Context, args *ticket.UpdateTicketInfoArgs) (*ticket.Ticket, error) {
	//TODO maybe not use
	panic("implement me")
}

func (a *TicketAggregate) ConfirmTicket(ctx context.Context, args *ticket.ConfirmTicketArgs) (*ticket.Ticket, error) {
	if args.ID == 0 {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Missing ID")
	}
	if args.ConfirmBy == 0 {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Missing confirm_by")
	}

	ticketCore, err := a.TicketStore(ctx).ID(args.ID).GetTicket()
	if err != nil {
		return nil, err
	}
	if ticketCore.Status != status5.Z && ticketCore.Status != status5.S {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Ticket đã đóng")
	}

	// system: leader có thể confirm mọi ticket
	// những người được assign vào ticker có thể confirm ticket
	if ticketCore.Type == ticket_type.System && !args.IsLeader {
		isPermission := false
		for _, v := range ticketCore.AssignedUserIDs {
			if v == args.ConfirmBy {
				isPermission = true
				break
			}
		}
		if !isPermission {
			return nil, cm.Errorf(cm.PermissionDenied, nil, "Ticket không thuộc sự quản lí của bạn")
		}
	}
	var ticketModel = &model.Ticket{
		ConfirmedAt: time.Now(),
		ConfirmedBy: args.ConfirmBy,
		Note:        args.Note,
		UpdatedBy:   args.ConfirmBy,
		Status:      status5.S,
		UpdatedAt:   time.Now(),
		State:       ticket_state.Processing,
	}
	if err = a.TicketStore(ctx).ID(args.ID).UpdateTicketDB(ticketModel); err != nil {
		return nil, err
	}
	return a.TicketStore(ctx).ID(args.ID).GetTicket()
}

func (a *TicketAggregate) CloseTicket(ctx context.Context, args *ticket.CloseTicketArgs) (*ticket.Ticket, error) {
	if args.ID == 0 {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Missing ID")
	}
	if args.ClosedBy == 0 {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Missing closed_by")
	}

	ticketCore, err := a.TicketStore(ctx).ID(args.ID).GetTicket()
	if err != nil {
		return nil, err
	}

	// system: chỉ leader hoặc người confirm mới được close ticket
	// internal: không phân quyền chỗ này
	if ticketCore.Type == ticket_type.System && !args.IsLeader && args.ClosedBy != ticketCore.ConfirmedBy {
		return nil, cm.Errorf(cm.PermissionDenied, nil, "Ticket không thuộc sự quản lí của bạn")
	}

	if ticketCore.Status != status5.Z && ticketCore.Status != status5.S {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Ticket đã đóng")
	}

	var ticketModel = &model.Ticket{
		ClosedAt:  time.Now(),
		ClosedBy:  args.ClosedBy,
		Note:      args.Note,
		UpdatedBy: args.ClosedBy,
		UpdatedAt: time.Now(),
		State:     args.State,
	}

	switch args.State {
	case ticket_state.Success, ticket_state.Fail,
		ticket_state.Ignore, ticket_state.Cancel:
		// no-op
	default:
		return nil, cm.Errorf(cm.InvalidArgument, nil, "state đóng ticket không hợp lệ")
	}
	ticketModel.Status = ticketModel.State.ToStatus5()
	if err = a.TicketStore(ctx).ID(args.ID).UpdateTicketDB(ticketModel); err != nil {
		return nil, err
	}
	return a.TicketStore(ctx).ID(args.ID).GetTicket()
}

func (a *TicketAggregate) ReopenTicket(ctx context.Context, args *ticket.ReopenTicketArgs) (*ticket.Ticket, error) {
	if args.ID == 0 {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Missing ID")
	}

	ticketCore, err := a.TicketStore(ctx).ID(args.ID).GetTicket()
	if err != nil {
		return nil, err
	}

	// khi reopen thì trạng thái ticket sẽ là new
	// nếu có người được assign vào ticket thì trạng thái sẽ là received
	var state = ticket_state.New
	if len(ticketCore.AssignedUserIDs) > 0 {
		state = ticket_state.Received
	}

	switch ticketCore.Status {
	case status5.N, status5.NS, status5.P:
		// no-op
	default:
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Ticket chưa được close không thể mở lại.")
	}

	var ticketModel = &model.Ticket{
		Note:  args.Note,
		State: state,
	}
	if err = a.TicketStore(ctx).ID(args.ID).UpdateTicketDB(ticketModel); err != nil {
		return nil, err
	}

	if err = a.TicketStore(ctx).ID(args.ID).UpdateTicketStatus(state.ToStatus5()); err != nil {
		return nil, err
	}

	return a.TicketStore(ctx).ID(args.ID).GetTicket()
}

func (a *TicketAggregate) AssignTicket(ctx context.Context, args *ticket.AssignedTicketArgs) (*ticket.Ticket, error) {
	if args.ID == 0 {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Missing ID")
	}

	ticketCore, err := a.TicketStore(ctx).ID(args.ID).GetTicket()
	if err != nil {
		return nil, err
	}

	assignedUserIDs := ticketCore.AssignedUserIDs

	// system: if not leader user will add themselves
	if ticketCore.Type == ticket_type.System && !args.IsLeader {
		for _, v := range assignedUserIDs {
			if v == args.UpdatedBy {
				return nil, cm.Errorf(cm.InvalidArgument, nil, "Bạn đã được thêm vào ticket này rồi.")
			}
		}
		assignedUserIDs = append(assignedUserIDs, args.UpdatedBy)
	} else {
		assignedUserIDs = args.AssignedUserIDs
	}

	if ticketCore.Status != status5.Z && ticketCore.Status != status5.S {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Ticket đã đóng")
	}

	ticketModel := &model.Ticket{
		UpdatedBy:       args.UpdatedBy,
		UpdatedAt:       time.Now(),
		AssignedUserIDs: assignedUserIDs,
	}

	// Khi assign ticket mới tạo cho 1 người: chuyển trạng thái từ new -> received
	// Còn lại thì giữ nguyên trạng thái cũ.
	if ticketCore.State == ticket_state.New {
		ticketModel.State = ticket_state.Received
		ticketModel.Status = ticket_state.Received.ToStatus5()
	}

	if err = a.TicketStore(ctx).ID(args.ID).UpdateTicketDB(ticketModel); err != nil {
		return nil, err
	}

	return a.TicketStore(ctx).ID(args.ID).GetTicket()
}

func (a *TicketAggregate) UnassignTicket(ctx context.Context, args *ticket.UnassignTicketArgs) (*ticket.Ticket, error) {
	if args.ID == 0 {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Missing ID")
	}

	ticketCore, err := a.TicketStore(ctx).ID(args.ID).GetTicket()
	if err != nil {
		return nil, err
	}

	if ticketCore.Status != status5.Z {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Bạn không thể bỏ chỉ định trên ticket khác trạng thái mới.")
	}

	var assignedUserIDs []dot.ID
	isExisted := false
	for _, v := range ticketCore.AssignedUserIDs {
		if v == args.UpdatedBy {
			isExisted = true
			continue
		}
		assignedUserIDs = append(assignedUserIDs, v)
	}
	if !isExisted {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Bạn chưa được thêm vào ticket này.")
	}

	var state = ticket_state.New
	if len(assignedUserIDs) > 0 {
		state = ticket_state.Received
	}

	ticketModel := &model.Ticket{
		State:           state,
		UpdatedBy:       args.UpdatedBy,
		UpdatedAt:       time.Now(),
		AssignedUserIDs: assignedUserIDs,
		Status:          state.ToStatus5(),
	}
	if err = a.TicketStore(ctx).ID(args.ID).UpdateTicketDB(ticketModel); err != nil {
		return nil, err
	}

	return a.TicketStore(ctx).ID(args.ID).GetTicket()
}

type listTicketLabelsArgs struct {
	Type   ticket_type.TicketType
	ShopID dot.ID
}

func (a *TicketAggregate) listTicketLabels(ctx context.Context, args listTicketLabelsArgs) ([]*ticket.TicketLabel, error) {
	var labels []*ticket.TicketLabel
	redisKey := generateTicketLabelKey(ctx, args.ShopID)
	if args.Type == ticket_type.System {
		redisKey = generateTicketLabelKey(ctx, 0)
	}
	err := a.RedisStore.Get(redisKey, &labels)
	switch err {
	case redis.ErrNil:
		// no-op
	case nil:
		return labels, nil
	default:
		return nil, err
	}
	query := a.TicketLabelStore(ctx).Type(args.Type)

	if args.Type == ticket_type.Internal {
		query = query.ShopID(args.ShopID)
	}
	labels, err = query.ListTicketLabels()
	if err != nil {
		return nil, err
	}

	if err := a.SetTicketLabels(ctx, args.ShopID, args.Type, &labels); err != nil {
		return nil, err
	}
	return labels, nil
}

func (a *TicketAggregate) SetTicketLabels(
	ctx context.Context, shopID dot.ID,
	typ ticket_type.TicketType, labels *[]*ticket.TicketLabel) error {
	*labels = MakeTreeLabel(*labels)
	if typ == ticket_type.System {
		return a.RedisStore.SetWithTTL(generateTicketLabelKey(ctx, 0), labels, 1*24*60*60)
	}
	return a.RedisStore.SetWithTTL(generateTicketLabelKey(ctx, shopID), labels, 1*24*60*60)
}

func (a *TicketAggregate) UpdateTicketRefTicketID(ctx context.Context, args *ticket.UpdateTicketRefTicketIDArgs) (*pbcm.UpdatedResponse, error) {
	if args.ID == 0 {
		return nil, cm.Errorf(cm.InvalidArgument, nil, "Thiếu thông tin ticket ID")
	}
	if !args.RefTicketID.Valid {
		return &pbcm.UpdatedResponse{Updated: 0}, nil
	}
	update := &model.Ticket{
		RefTicketID: args.RefTicketID,
	}
	if err := a.TicketStore(ctx).ID(args.ID).UpdateTicketDB(update); err != nil {
		return nil, err
	}
	return &pbcm.UpdatedResponse{Updated: 1}, nil
}

func generateTicketLabelKey(ctx context.Context, shopID dot.ID) string {
	return fmt.Sprintf("ticket_labels:%s:wl%s:sh%s", ticketLabelsVersion, wl.GetWLPartnerID(ctx), shopID.String())
}