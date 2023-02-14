// +build !generator

// Code generated by generator api. DO NOT EDIT.

package gateway

import (
	context "context"

	capi "o.o/capi"
)

type CommandBus struct{ bus capi.Bus }

func NewCommandBus(bus capi.Bus) CommandBus { return CommandBus{bus} }

func (b CommandBus) Dispatch(ctx context.Context, msg interface{ command() }) error {
	return b.bus.Dispatch(ctx, msg)
}

type GetResultCommand struct {
	BillCode        string `json:"billcode"`
	CustMsisdn      string `json:"cust_msisdn"`
	ErrorCode       string `json:"error_code"`
	MerchantCode    string `json:"merchant_code"`
	OrderID         string `json:"order_id"`
	PaymentStatus   string `json:"payment_status"`
	TransAmount     int    `json:"trans_amount"`
	VtTransactionID string `json:"vt_transaction_id"`
	CheckSum        string `json:"check_sum"`

	Result *GetResultResult `json:"-"`
}

func (h AggregateHandler) HandleGetResult(ctx context.Context, msg *GetResultCommand) (err error) {
	msg.Result, err = h.inner.GetResult(msg.GetArgs(ctx))
	return err
}

type ValidateTransactionCommand struct {
	BillCode     string `json:"billcode"`
	MerchantCode string `json:"merchant_code"`
	OrderID      string `json:"order_id"`
	CheckSum     string `json:"check_sum"`

	Result *ValidateTransactionResult `json:"-"`
}

func (h AggregateHandler) HandleValidateTransaction(ctx context.Context, msg *ValidateTransactionCommand) (err error) {
	msg.Result, err = h.inner.ValidateTransaction(msg.GetArgs(ctx))
	return err
}

// implement interfaces

func (q *GetResultCommand) command()           {}
func (q *ValidateTransactionCommand) command() {}

// implement conversion

func (q *GetResultCommand) GetArgs(ctx context.Context) (_ context.Context, _ *GetResultArgs) {
	return ctx,
		&GetResultArgs{
			BillCode:        q.BillCode,
			CustMsisdn:      q.CustMsisdn,
			ErrorCode:       q.ErrorCode,
			MerchantCode:    q.MerchantCode,
			OrderID:         q.OrderID,
			PaymentStatus:   q.PaymentStatus,
			TransAmount:     q.TransAmount,
			VtTransactionID: q.VtTransactionID,
			CheckSum:        q.CheckSum,
		}
}

func (q *GetResultCommand) SetGetResultArgs(args *GetResultArgs) {
	q.BillCode = args.BillCode
	q.CustMsisdn = args.CustMsisdn
	q.ErrorCode = args.ErrorCode
	q.MerchantCode = args.MerchantCode
	q.OrderID = args.OrderID
	q.PaymentStatus = args.PaymentStatus
	q.TransAmount = args.TransAmount
	q.VtTransactionID = args.VtTransactionID
	q.CheckSum = args.CheckSum
}

func (q *ValidateTransactionCommand) GetArgs(ctx context.Context) (_ context.Context, _ *ValidateTransactionArgs) {
	return ctx,
		&ValidateTransactionArgs{
			BillCode:     q.BillCode,
			MerchantCode: q.MerchantCode,
			OrderID:      q.OrderID,
			CheckSum:     q.CheckSum,
		}
}

func (q *ValidateTransactionCommand) SetValidateTransactionArgs(args *ValidateTransactionArgs) {
	q.BillCode = args.BillCode
	q.MerchantCode = args.MerchantCode
	q.OrderID = args.OrderID
	q.CheckSum = args.CheckSum
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
	b.AddHandler(h.HandleGetResult)
	b.AddHandler(h.HandleValidateTransaction)
	return CommandBus{b}
}
