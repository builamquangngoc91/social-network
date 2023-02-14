package aggregate

import (
	"context"
	"time"

	"o.o/api/etelecom"
	"o.o/api/etelecom/call_direction"
	"o.o/api/etelecom/call_state"
	contacting "o.o/api/main/contact"
	cm "o.o/backend/pkg/common"
	"o.o/capi/dot"
	"o.o/common/l"
)

func (a *EtelecomAggregate) CreateCallLog(ctx context.Context, args *etelecom.CreateCallLogArgs) (*etelecom.CallLog, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}
	extension, err := a.extensionStore(ctx).ID(args.ExtensionID).GetExtension()
	if err != nil {
		return nil, err
	}

	callState := args.CallState
	if callState == 0 {
		callState = call_state.Answered
	}
	callLog := &etelecom.CallLog{
		AccountID:         args.AccountID,
		Caller:            args.Caller,
		Callee:            args.Callee,
		Direction:         args.Direction,
		ExtensionID:       args.ExtensionID,
		HotlineID:         extension.HotlineID,
		ContactID:         args.ContactID,
		CallState:         callState,
		CallStatus:        callState.ToStatus5(),
		ExternalSessionID: args.ExternalSessionID,
		StartedAt:         args.StartedAt,
		EndedAt:           args.EndedAt,
		UserID:            extension.UserID,
		OwnerID:           args.OwnerID,
	}
	if callLog.StartedAt.IsZero() {
		// workaround
		callLog.StartedAt = time.Now()
	}

	cLog, err := a.callLogStore(ctx).ExternalSessionID(args.ExternalSessionID).GetCallLog()
	switch cm.ErrorCode(err) {
	case cm.NoError:
		// update call log
		if cLog.CallState == call_state.Answered {
			// do not update. Keep the call log
			return cLog, nil
		}
		err = a.callLogStore(ctx).ID(cLog.ID).UpdateCallLog(callLog)
		if err != nil {
			return nil, err
		}
		return a.callLogStore(ctx).ID(cLog.ID).GetCallLog()
	case cm.NotFound:
		// create call log
		callLog.ID = cm.NewID()
		return a.callLogStore(ctx).CreateCallLog(callLog)
	default:
		return nil, err
	}
}

func (a *EtelecomAggregate) CreateOrUpdateCallLogFromCDR(
	ctx context.Context, args *etelecom.CreateOrUpdateCallLogFromCDRArgs,
) (*etelecom.CallLog, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}
	var callLog etelecom.CallLog
	if err := scheme.Convert(args, &callLog); err != nil {
		return nil, err
	}
	callLog.CallStatus = callLog.CallState.ToStatus5()

	if args.ExtensionID != 0 {
		extension, err := a.extensionStore(ctx).ID(args.ExtensionID).GetExtension()
		if err != nil {
			return nil, err
		}

		var phoneNumber string
		if args.Direction == call_direction.In {
			phoneNumber = args.Caller
		} else {
			phoneNumber = args.Callee
		}
		getContactsQuery := &contacting.GetContactsByPhoneQuery{
			ShopID: extension.AccountID,
			Phone:  phoneNumber,
		}
		if err = a.contactQuery.Dispatch(ctx, getContactsQuery); err != nil {
			return nil, err
		}
		contacts := getContactsQuery.Result
		if len(contacts) > 0 {
			callLog.ContactID = contacts[0].ID
		}
		callLog.AccountID = extension.AccountID
		callLog.UserID = extension.UserID
	}

	oldCallLog, err := a.callLogStore(ctx).ExternalSessionID(args.ExternalSessionID).GetCallLog()
	if err != nil && cm.ErrorCode(err) != cm.NotFound {
		return nil, err
	}

	var result *etelecom.CallLog
	if err != nil && cm.ErrorCode(err) == cm.NotFound {
		// create new call log
		callLog.ID = cm.NewID()
		result, err = a.callLogStore(ctx).CreateCallLog(&callLog)
		if err != nil {
			return nil, err
		}
	} else {
		// update call log
		// update via cdr info
		if err = a.callLogStore(ctx).ID(oldCallLog.ID).UpdateCallLog(&callLog); err != nil {
			return nil, err
		}
		result, err = a.callLogStore(ctx).ID(oldCallLog.ID).GetCallLog()
		if err != nil {
			return nil, err
		}
	}

	event := &etelecom.CallLogCalcPostageEvent{
		ID:         result.ID,
		Direction:  result.Direction,
		Callee:     result.Callee,
		Duration:   result.Duration,
		CallStatus: result.CallStatus,
		HotlineID:  result.HotlineID,
	}
	if _err := a.eventBus.Publish(ctx, event); _err != nil {
		ll.Error("Error event call log created :: ", l.Error(_err))
	}
	return result, nil
}

func (a *EtelecomAggregate) getHotlineIDs(ctx context.Context, connectionID, ownerID dot.ID) (hotlineIDs []dot.ID, _ error) {
	hotlines, err := a.hotlineStore(ctx).OptionalOwnerID(ownerID).ConnectionID(connectionID).ListHotlines()
	if err != nil {
		return nil, err
	}
	for _, hl := range hotlines {
		hotlineIDs = append(hotlineIDs, hl.ID)
	}
	return hotlineIDs, nil
}

func (a *EtelecomAggregate) UpdateCallLogPostage(ctx context.Context, args *etelecom.UpdateCallLogPostageArgs) error {
	if args.ID == 0 {
		return cm.Errorf(cm.InvalidArgument, nil, "Missing call log ID")
	}
	update := &etelecom.CallLog{
		DurationPostage: args.DurationForPostage,
		Postage:         args.Postage,
	}
	return a.callLogStore(ctx).ID(args.ID).UpdateCallLog(update)
}
