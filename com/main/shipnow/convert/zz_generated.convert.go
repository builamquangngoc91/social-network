// +build !generator

// Code generated by generator convert. DO NOT EDIT.

package convert

import (
	shipnow "o.o/api/main/shipnow"
	shipnowtypes "o.o/api/main/shipnow/types"
	shippingtypes "o.o/api/main/shipping/types"
	orderingconvert "o.o/backend/com/main/ordering/convert"
	shipnowmodel "o.o/backend/com/main/shipnow/model"
	shippingconvert "o.o/backend/com/main/shipping/convert"
	conversion "o.o/backend/pkg/common/conversion"
)

/*
Custom conversions:
    DeliveryPoint           // not use, no conversions between params
    DeliveryPointToModel    // not use, no conversions between params
    OrderToDeliveryPoint    // not use, no conversions between params
    Shipnow                 // in use
    ShipnowToModel          // in use
    SyncStateToModel        // not use, no conversions between params

Ignored functions:
    DeliveryPoints               // params are not pointer to named types
    DeliveryPointsToModel        // params are not pointer to named types
    GetValueInfo                 // params are not pointer to named types
    GetWeightInfo                // params are not pointer to named types
    Shipnows                     // params are not pointer to named types
    normalizePhoneAndFullName    // not recognized
    normalizeSearch              // params are not pointer to named types
*/

func RegisterConversions(s *conversion.Scheme) {
	registerConversions(s)
}

func registerConversions(s *conversion.Scheme) {
	s.Register((*shipnowmodel.DeliveryPoint)(nil), (*shipnowtypes.DeliveryPoint)(nil), func(arg, out interface{}) error {
		Convert_shipnowmodel_DeliveryPoint_shipnowtypes_DeliveryPoint(arg.(*shipnowmodel.DeliveryPoint), out.(*shipnowtypes.DeliveryPoint))
		return nil
	})
	s.Register(([]*shipnowmodel.DeliveryPoint)(nil), (*[]*shipnowtypes.DeliveryPoint)(nil), func(arg, out interface{}) error {
		out0 := Convert_shipnowmodel_DeliveryPoints_shipnowtypes_DeliveryPoints(arg.([]*shipnowmodel.DeliveryPoint))
		*out.(*[]*shipnowtypes.DeliveryPoint) = out0
		return nil
	})
	s.Register((*shipnowtypes.DeliveryPoint)(nil), (*shipnowmodel.DeliveryPoint)(nil), func(arg, out interface{}) error {
		Convert_shipnowtypes_DeliveryPoint_shipnowmodel_DeliveryPoint(arg.(*shipnowtypes.DeliveryPoint), out.(*shipnowmodel.DeliveryPoint))
		return nil
	})
	s.Register(([]*shipnowtypes.DeliveryPoint)(nil), (*[]*shipnowmodel.DeliveryPoint)(nil), func(arg, out interface{}) error {
		out0 := Convert_shipnowtypes_DeliveryPoints_shipnowmodel_DeliveryPoints(arg.([]*shipnowtypes.DeliveryPoint))
		*out.(*[]*shipnowmodel.DeliveryPoint) = out0
		return nil
	})
	s.Register((*shipnowmodel.ShipnowFulfillment)(nil), (*shipnow.ShipnowFulfillment)(nil), func(arg, out interface{}) error {
		Convert_shipnowmodel_ShipnowFulfillment_shipnow_ShipnowFulfillment(arg.(*shipnowmodel.ShipnowFulfillment), out.(*shipnow.ShipnowFulfillment))
		return nil
	})
	s.Register(([]*shipnowmodel.ShipnowFulfillment)(nil), (*[]*shipnow.ShipnowFulfillment)(nil), func(arg, out interface{}) error {
		out0 := Convert_shipnowmodel_ShipnowFulfillments_shipnow_ShipnowFulfillments(arg.([]*shipnowmodel.ShipnowFulfillment))
		*out.(*[]*shipnow.ShipnowFulfillment) = out0
		return nil
	})
	s.Register((*shipnow.ShipnowFulfillment)(nil), (*shipnowmodel.ShipnowFulfillment)(nil), func(arg, out interface{}) error {
		Convert_shipnow_ShipnowFulfillment_shipnowmodel_ShipnowFulfillment(arg.(*shipnow.ShipnowFulfillment), out.(*shipnowmodel.ShipnowFulfillment))
		return nil
	})
	s.Register(([]*shipnow.ShipnowFulfillment)(nil), (*[]*shipnowmodel.ShipnowFulfillment)(nil), func(arg, out interface{}) error {
		out0 := Convert_shipnow_ShipnowFulfillments_shipnowmodel_ShipnowFulfillments(arg.([]*shipnow.ShipnowFulfillment))
		*out.(*[]*shipnowmodel.ShipnowFulfillment) = out0
		return nil
	})
}

//-- convert o.o/api/main/shipnow.DeliveryPoint --//

func Convert_shipnowmodel_DeliveryPoint_shipnowtypes_DeliveryPoint(arg *shipnowmodel.DeliveryPoint, out *shipnowtypes.DeliveryPoint) *shipnowtypes.DeliveryPoint {
	return DeliveryPoint(arg)
}

func convert_shipnowmodel_DeliveryPoint_shipnowtypes_DeliveryPoint(arg *shipnowmodel.DeliveryPoint, out *shipnowtypes.DeliveryPoint) {
	out.ShippingAddress = orderingconvert.Convert_orderingmodel_OrderAddress_orderingtypes_Address(arg.ShippingAddress, nil)
	out.Lines = nil                             // zero value
	out.ShippingNote = arg.ShippingNote         // simple assign
	out.OrderId = 0                             // zero value
	out.OrderCode = arg.OrderCode               // simple assign
	out.WeightInfo = shippingtypes.WeightInfo{} // zero value
	out.ValueInfo = shippingtypes.ValueInfo{}   // zero value
	out.TryOn = arg.TryOn                       // simple assign
	out.ShippingState = arg.ShippingState       // simple assign
}

func Convert_shipnowmodel_DeliveryPoints_shipnowtypes_DeliveryPoints(args []*shipnowmodel.DeliveryPoint) (outs []*shipnowtypes.DeliveryPoint) {
	if args == nil {
		return nil
	}
	tmps := make([]shipnowtypes.DeliveryPoint, len(args))
	outs = make([]*shipnowtypes.DeliveryPoint, len(args))
	for i := range tmps {
		outs[i] = Convert_shipnowmodel_DeliveryPoint_shipnowtypes_DeliveryPoint(args[i], &tmps[i])
	}
	return outs
}

func Convert_shipnowtypes_DeliveryPoint_shipnowmodel_DeliveryPoint(arg *shipnowtypes.DeliveryPoint, out *shipnowmodel.DeliveryPoint) *shipnowmodel.DeliveryPoint {
	return DeliveryPointToModel(arg)
}

func convert_shipnowtypes_DeliveryPoint_shipnowmodel_DeliveryPoint(arg *shipnowtypes.DeliveryPoint, out *shipnowmodel.DeliveryPoint) {
	out.ShippingAddress = orderingconvert.Convert_orderingtypes_Address_orderingmodel_OrderAddress(arg.ShippingAddress, nil)
	out.Items = nil                       // zero value
	out.OrderID = 0                       // zero value
	out.OrderCode = arg.OrderCode         // simple assign
	out.GrossWeight = 0                   // zero value
	out.ChargeableWeight = 0              // zero value
	out.Length = 0                        // zero value
	out.Width = 0                         // zero value
	out.Height = 0                        // zero value
	out.BasketValue = 0                   // zero value
	out.CODAmount = 0                     // zero value
	out.TryOn = arg.TryOn                 // simple assign
	out.ShippingNote = arg.ShippingNote   // simple assign
	out.ShippingState = arg.ShippingState // simple assign
}

func Convert_shipnowtypes_DeliveryPoints_shipnowmodel_DeliveryPoints(args []*shipnowtypes.DeliveryPoint) (outs []*shipnowmodel.DeliveryPoint) {
	if args == nil {
		return nil
	}
	tmps := make([]shipnowmodel.DeliveryPoint, len(args))
	outs = make([]*shipnowmodel.DeliveryPoint, len(args))
	for i := range tmps {
		outs[i] = Convert_shipnowtypes_DeliveryPoint_shipnowmodel_DeliveryPoint(args[i], &tmps[i])
	}
	return outs
}

//-- convert o.o/api/main/shipnow.ShipnowFulfillment --//

func Convert_shipnowmodel_ShipnowFulfillment_shipnow_ShipnowFulfillment(arg *shipnowmodel.ShipnowFulfillment, out *shipnow.ShipnowFulfillment) *shipnow.ShipnowFulfillment {
	return Shipnow(arg)
}

func convert_shipnowmodel_ShipnowFulfillment_shipnow_ShipnowFulfillment(arg *shipnowmodel.ShipnowFulfillment, out *shipnow.ShipnowFulfillment) {
	out.ID = arg.ID               // simple assign
	out.ShopID = arg.ShopID       // simple assign
	out.PartnerID = arg.PartnerID // simple assign
	out.PickupAddress = orderingconvert.Convert_orderingmodel_OrderAddress_orderingtypes_Address(arg.PickupAddress, nil)
	out.DeliveryPoints = Convert_shipnowmodel_DeliveryPoints_shipnowtypes_DeliveryPoints(arg.DeliveryPoints)
	out.Carrier = arg.Carrier                                       // simple assign
	out.ShippingServiceCode = arg.ShippingServiceCode               // simple assign
	out.ShippingServiceFee = arg.ShippingServiceFee                 // simple assign
	out.ShippingServiceName = arg.ShippingServiceName               // simple assign
	out.ShippingServiceDescription = arg.ShippingServiceDescription // simple assign
	out.WeightInfo = shippingtypes.WeightInfo{}                     // zero value
	out.ValueInfo = shippingtypes.ValueInfo{}                       // zero value
	out.ShippingNote = arg.ShippingNote                             // simple assign
	out.RequestPickupAt = arg.RequestPickupAt                       // simple assign
	out.Status = arg.Status                                         // simple assign
	out.ShippingStatus = arg.ShippingStatus                         // simple assign
	out.ShippingCode = arg.ShippingCode                             // simple assign
	out.ShippingState = arg.ShippingState                           // simple assign
	out.ConfirmStatus = arg.ConfirmStatus                           // simple assign
	out.OrderIDs = arg.OrderIDs                                     // simple assign
	out.ShippingCreatedAt = arg.ShippingCreatedAt                   // simple assign
	out.CreatedAt = arg.CreatedAt                                   // simple assign
	out.UpdatedAt = arg.UpdatedAt                                   // simple assign
	out.EtopPaymentStatus = arg.EtopPaymentStatus                   // simple assign
	out.CODEtopTransferedAt = arg.CODEtopTransferedAt               // simple assign
	out.ShippingPickingAt = arg.ShippingPickingAt                   // simple assign
	out.ShippingDeliveringAt = arg.ShippingDeliveringAt             // simple assign
	out.ShippingDeliveredAt = arg.ShippingDeliveredAt               // simple assign
	out.ShippingCancelledAt = arg.ShippingCancelledAt               // simple assign
	out.ShippingSharedLink = arg.ShippingSharedLink                 // simple assign
	out.CancelReason = arg.CancelReason                             // simple assign
	out.ConnectionID = arg.ConnectionID                             // simple assign
	out.ConnectionMethod = arg.ConnectionMethod                     // simple assign
	out.FeeLines = shippingconvert.Convert_sharemodel_ShippingFeeLines_shippingtypes_ShippingFeeLines(arg.FeeLines)
	out.CarrierFeeLines = shippingconvert.Convert_sharemodel_ShippingFeeLines_shippingtypes_ShippingFeeLines(arg.CarrierFeeLines)
	out.ExternalID = arg.ExternalID   // simple assign
	out.TotalFee = arg.TotalFee       // simple assign
	out.Coupon = arg.Coupon           // simple assign
	out.DriverPhone = arg.DriverPhone // simple assign
	out.DriverName = arg.DriverName   // simple assign
}

func Convert_shipnowmodel_ShipnowFulfillments_shipnow_ShipnowFulfillments(args []*shipnowmodel.ShipnowFulfillment) (outs []*shipnow.ShipnowFulfillment) {
	if args == nil {
		return nil
	}
	tmps := make([]shipnow.ShipnowFulfillment, len(args))
	outs = make([]*shipnow.ShipnowFulfillment, len(args))
	for i := range tmps {
		outs[i] = Convert_shipnowmodel_ShipnowFulfillment_shipnow_ShipnowFulfillment(args[i], &tmps[i])
	}
	return outs
}

func Convert_shipnow_ShipnowFulfillment_shipnowmodel_ShipnowFulfillment(arg *shipnow.ShipnowFulfillment, out *shipnowmodel.ShipnowFulfillment) *shipnowmodel.ShipnowFulfillment {
	return ShipnowToModel(arg)
}

func convert_shipnow_ShipnowFulfillment_shipnowmodel_ShipnowFulfillment(arg *shipnow.ShipnowFulfillment, out *shipnowmodel.ShipnowFulfillment) {
	out.ID = arg.ID               // simple assign
	out.ShopID = arg.ShopID       // simple assign
	out.PartnerID = arg.PartnerID // simple assign
	out.OrderIDs = arg.OrderIDs   // simple assign
	out.PickupAddress = orderingconvert.Convert_orderingtypes_Address_orderingmodel_OrderAddress(arg.PickupAddress, nil)
	out.Carrier = arg.Carrier                                       // simple assign
	out.ShippingServiceCode = arg.ShippingServiceCode               // simple assign
	out.ShippingServiceFee = arg.ShippingServiceFee                 // simple assign
	out.ShippingServiceName = arg.ShippingServiceName               // simple assign
	out.ShippingServiceDescription = arg.ShippingServiceDescription // simple assign
	out.ChargeableWeight = 0                                        // zero value
	out.GrossWeight = 0                                             // zero value
	out.BasketValue = 0                                             // zero value
	out.CODAmount = 0                                               // zero value
	out.ShippingNote = arg.ShippingNote                             // simple assign
	out.RequestPickupAt = arg.RequestPickupAt                       // simple assign
	out.DeliveryPoints = Convert_shipnowtypes_DeliveryPoints_shipnowmodel_DeliveryPoints(arg.DeliveryPoints)
	out.CancelReason = arg.CancelReason           // simple assign
	out.Status = arg.Status                       // simple assign
	out.ConfirmStatus = arg.ConfirmStatus         // simple assign
	out.ShippingStatus = arg.ShippingStatus       // simple assign
	out.EtopPaymentStatus = arg.EtopPaymentStatus // simple assign
	out.ShippingState = arg.ShippingState         // simple assign
	out.ShippingCode = arg.ShippingCode           // simple assign
	out.FeeLines = shippingconvert.Convert_shippingtypes_ShippingFeeLines_sharemodel_ShippingFeeLines(arg.FeeLines)
	out.CarrierFeeLines = shippingconvert.Convert_shippingtypes_ShippingFeeLines_sharemodel_ShippingFeeLines(arg.CarrierFeeLines)
	out.TotalFee = arg.TotalFee                         // simple assign
	out.ShippingCreatedAt = arg.ShippingCreatedAt       // simple assign
	out.ShippingPickingAt = arg.ShippingPickingAt       // simple assign
	out.ShippingDeliveringAt = arg.ShippingDeliveringAt // simple assign
	out.ShippingDeliveredAt = arg.ShippingDeliveredAt   // simple assign
	out.ShippingCancelledAt = arg.ShippingCancelledAt   // simple assign
	out.SyncStatus = 0                                  // zero value
	out.SyncStates = nil                                // zero value
	out.CreatedAt = arg.CreatedAt                       // simple assign
	out.UpdatedAt = arg.UpdatedAt                       // simple assign
	out.CODEtopTransferedAt = arg.CODEtopTransferedAt   // simple assign
	out.ShippingSharedLink = arg.ShippingSharedLink     // simple assign
	out.AddressToProvinceCode = ""                      // zero value
	out.AddressToDistrictCode = ""                      // zero value
	out.AddressToPhone = ""                             // zero value
	out.AddressToFullNameNorm = ""                      // zero value
	out.ConnectionID = arg.ConnectionID                 // simple assign
	out.ConnectionMethod = arg.ConnectionMethod         // simple assign
	out.ExternalID = arg.ExternalID                     // simple assign
	out.Coupon = arg.Coupon                             // simple assign
	out.DriverPhone = arg.DriverPhone                   // simple assign
	out.DriverName = arg.DriverName                     // simple assign
	out.Rid = 0                                         // zero value
}

func Convert_shipnow_ShipnowFulfillments_shipnowmodel_ShipnowFulfillments(args []*shipnow.ShipnowFulfillment) (outs []*shipnowmodel.ShipnowFulfillment) {
	if args == nil {
		return nil
	}
	tmps := make([]shipnowmodel.ShipnowFulfillment, len(args))
	outs = make([]*shipnowmodel.ShipnowFulfillment, len(args))
	for i := range tmps {
		outs[i] = Convert_shipnow_ShipnowFulfillment_shipnowmodel_ShipnowFulfillment(args[i], &tmps[i])
	}
	return outs
}