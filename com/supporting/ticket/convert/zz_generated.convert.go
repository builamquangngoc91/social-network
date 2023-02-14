// +build !generator

// Code generated by generator convert. DO NOT EDIT.

package convert

import (
	time "time"

	ticket "o.o/api/supporting/ticket"
	status5 "o.o/api/top/types/etc/status5"
	ticketmodel "o.o/backend/com/supporting/ticket/model"
	conversion "o.o/backend/pkg/common/conversion"
)

/*
Custom conversions:
    createComment    // in use
    createLabel      // in use
    createTicket     // in use

Ignored functions: (none)
*/

func RegisterConversions(s *conversion.Scheme) {
	registerConversions(s)
}

func registerConversions(s *conversion.Scheme) {
	s.Register((*ticketmodel.Ticket)(nil), (*ticket.Ticket)(nil), func(arg, out interface{}) error {
		Convert_ticketmodel_Ticket_ticket_Ticket(arg.(*ticketmodel.Ticket), out.(*ticket.Ticket))
		return nil
	})
	s.Register(([]*ticketmodel.Ticket)(nil), (*[]*ticket.Ticket)(nil), func(arg, out interface{}) error {
		out0 := Convert_ticketmodel_Tickets_ticket_Tickets(arg.([]*ticketmodel.Ticket))
		*out.(*[]*ticket.Ticket) = out0
		return nil
	})
	s.Register((*ticket.Ticket)(nil), (*ticketmodel.Ticket)(nil), func(arg, out interface{}) error {
		Convert_ticket_Ticket_ticketmodel_Ticket(arg.(*ticket.Ticket), out.(*ticketmodel.Ticket))
		return nil
	})
	s.Register(([]*ticket.Ticket)(nil), (*[]*ticketmodel.Ticket)(nil), func(arg, out interface{}) error {
		out0 := Convert_ticket_Tickets_ticketmodel_Tickets(arg.([]*ticket.Ticket))
		*out.(*[]*ticketmodel.Ticket) = out0
		return nil
	})
	s.Register((*ticket.CreateTicketArgs)(nil), (*ticket.Ticket)(nil), func(arg, out interface{}) error {
		Apply_ticket_CreateTicketArgs_ticket_Ticket(arg.(*ticket.CreateTicketArgs), out.(*ticket.Ticket))
		return nil
	})
	s.Register((*ticket.UpdateTicketInfoArgs)(nil), (*ticket.Ticket)(nil), func(arg, out interface{}) error {
		Apply_ticket_UpdateTicketInfoArgs_ticket_Ticket(arg.(*ticket.UpdateTicketInfoArgs), out.(*ticket.Ticket))
		return nil
	})
	s.Register((*ticketmodel.TicketComment)(nil), (*ticket.TicketComment)(nil), func(arg, out interface{}) error {
		Convert_ticketmodel_TicketComment_ticket_TicketComment(arg.(*ticketmodel.TicketComment), out.(*ticket.TicketComment))
		return nil
	})
	s.Register(([]*ticketmodel.TicketComment)(nil), (*[]*ticket.TicketComment)(nil), func(arg, out interface{}) error {
		out0 := Convert_ticketmodel_TicketComments_ticket_TicketComments(arg.([]*ticketmodel.TicketComment))
		*out.(*[]*ticket.TicketComment) = out0
		return nil
	})
	s.Register((*ticket.TicketComment)(nil), (*ticketmodel.TicketComment)(nil), func(arg, out interface{}) error {
		Convert_ticket_TicketComment_ticketmodel_TicketComment(arg.(*ticket.TicketComment), out.(*ticketmodel.TicketComment))
		return nil
	})
	s.Register(([]*ticket.TicketComment)(nil), (*[]*ticketmodel.TicketComment)(nil), func(arg, out interface{}) error {
		out0 := Convert_ticket_TicketComments_ticketmodel_TicketComments(arg.([]*ticket.TicketComment))
		*out.(*[]*ticketmodel.TicketComment) = out0
		return nil
	})
	s.Register((*ticket.CreateTicketCommentArgs)(nil), (*ticket.TicketComment)(nil), func(arg, out interface{}) error {
		Apply_ticket_CreateTicketCommentArgs_ticket_TicketComment(arg.(*ticket.CreateTicketCommentArgs), out.(*ticket.TicketComment))
		return nil
	})
	s.Register((*ticket.CreateTicketCommentWebhookArgs)(nil), (*ticket.TicketComment)(nil), func(arg, out interface{}) error {
		Apply_ticket_CreateTicketCommentWebhookArgs_ticket_TicketComment(arg.(*ticket.CreateTicketCommentWebhookArgs), out.(*ticket.TicketComment))
		return nil
	})
	s.Register((*ticket.UpdateTicketCommentArgs)(nil), (*ticket.TicketComment)(nil), func(arg, out interface{}) error {
		Apply_ticket_UpdateTicketCommentArgs_ticket_TicketComment(arg.(*ticket.UpdateTicketCommentArgs), out.(*ticket.TicketComment))
		return nil
	})
	s.Register((*ticketmodel.TicketLabel)(nil), (*ticket.TicketLabel)(nil), func(arg, out interface{}) error {
		Convert_ticketmodel_TicketLabel_ticket_TicketLabel(arg.(*ticketmodel.TicketLabel), out.(*ticket.TicketLabel))
		return nil
	})
	s.Register(([]*ticketmodel.TicketLabel)(nil), (*[]*ticket.TicketLabel)(nil), func(arg, out interface{}) error {
		out0 := Convert_ticketmodel_TicketLabels_ticket_TicketLabels(arg.([]*ticketmodel.TicketLabel))
		*out.(*[]*ticket.TicketLabel) = out0
		return nil
	})
	s.Register((*ticket.TicketLabel)(nil), (*ticketmodel.TicketLabel)(nil), func(arg, out interface{}) error {
		Convert_ticket_TicketLabel_ticketmodel_TicketLabel(arg.(*ticket.TicketLabel), out.(*ticketmodel.TicketLabel))
		return nil
	})
	s.Register(([]*ticket.TicketLabel)(nil), (*[]*ticketmodel.TicketLabel)(nil), func(arg, out interface{}) error {
		out0 := Convert_ticket_TicketLabels_ticketmodel_TicketLabels(arg.([]*ticket.TicketLabel))
		*out.(*[]*ticketmodel.TicketLabel) = out0
		return nil
	})
	s.Register((*ticket.CreateTicketLabelArgs)(nil), (*ticket.TicketLabel)(nil), func(arg, out interface{}) error {
		Apply_ticket_CreateTicketLabelArgs_ticket_TicketLabel(arg.(*ticket.CreateTicketLabelArgs), out.(*ticket.TicketLabel))
		return nil
	})
	s.Register((*ticket.UpdateTicketLabelArgs)(nil), (*ticket.TicketLabel)(nil), func(arg, out interface{}) error {
		Apply_ticket_UpdateTicketLabelArgs_ticket_TicketLabel(arg.(*ticket.UpdateTicketLabelArgs), out.(*ticket.TicketLabel))
		return nil
	})
	s.Register((*ticketmodel.TicketLabelExternal)(nil), (*ticket.TicketLabelExternal)(nil), func(arg, out interface{}) error {
		Convert_ticketmodel_TicketLabelExternal_ticket_TicketLabelExternal(arg.(*ticketmodel.TicketLabelExternal), out.(*ticket.TicketLabelExternal))
		return nil
	})
	s.Register(([]*ticketmodel.TicketLabelExternal)(nil), (*[]*ticket.TicketLabelExternal)(nil), func(arg, out interface{}) error {
		out0 := Convert_ticketmodel_TicketLabelExternals_ticket_TicketLabelExternals(arg.([]*ticketmodel.TicketLabelExternal))
		*out.(*[]*ticket.TicketLabelExternal) = out0
		return nil
	})
	s.Register((*ticket.TicketLabelExternal)(nil), (*ticketmodel.TicketLabelExternal)(nil), func(arg, out interface{}) error {
		Convert_ticket_TicketLabelExternal_ticketmodel_TicketLabelExternal(arg.(*ticket.TicketLabelExternal), out.(*ticketmodel.TicketLabelExternal))
		return nil
	})
	s.Register(([]*ticket.TicketLabelExternal)(nil), (*[]*ticketmodel.TicketLabelExternal)(nil), func(arg, out interface{}) error {
		out0 := Convert_ticket_TicketLabelExternals_ticketmodel_TicketLabelExternals(arg.([]*ticket.TicketLabelExternal))
		*out.(*[]*ticketmodel.TicketLabelExternal) = out0
		return nil
	})
	s.Register((*ticket.CreateTicketLabelExternalArgs)(nil), (*ticket.TicketLabelExternal)(nil), func(arg, out interface{}) error {
		Apply_ticket_CreateTicketLabelExternalArgs_ticket_TicketLabelExternal(arg.(*ticket.CreateTicketLabelExternalArgs), out.(*ticket.TicketLabelExternal))
		return nil
	})
	s.Register((*ticket.UpdateTicketLabelExternalArgs)(nil), (*ticket.TicketLabelExternal)(nil), func(arg, out interface{}) error {
		Apply_ticket_UpdateTicketLabelExternalArgs_ticket_TicketLabelExternal(arg.(*ticket.UpdateTicketLabelExternalArgs), out.(*ticket.TicketLabelExternal))
		return nil
	})
}

//-- convert o.o/api/supporting/ticket.Ticket --//

func Convert_ticketmodel_Ticket_ticket_Ticket(arg *ticketmodel.Ticket, out *ticket.Ticket) *ticket.Ticket {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticket.Ticket{}
	}
	convert_ticketmodel_Ticket_ticket_Ticket(arg, out)
	return out
}

func convert_ticketmodel_Ticket_ticket_Ticket(arg *ticketmodel.Ticket, out *ticket.Ticket) {
	out.ExternalShippingCode = ""             // zero value
	out.ExternalID = arg.ExternalID           // simple assign
	out.ID = arg.ID                           // simple assign
	out.Code = arg.Code                       // simple assign
	out.AssignedUserIDs = arg.AssignedUserIDs // simple assign
	out.AccountID = arg.AccountID             // simple assign
	out.LabelIDs = arg.LabelIDs               // simple assign
	out.RefTicketID = arg.RefTicketID         // simple assign
	out.Title = arg.Title                     // simple assign
	out.Description = arg.Description         // simple assign
	out.Note = arg.Note                       // simple assign
	out.AdminNote = ""                        // zero value
	out.RefID = arg.RefID                     // simple assign
	out.RefType = arg.RefType                 // simple assign
	out.RefCode = arg.RefCode                 // simple assign
	out.Source = arg.Source                   // simple assign
	out.State = arg.State                     // simple assign
	out.Status = arg.Status                   // simple assign
	out.CreatedBy = arg.CreatedBy             // simple assign
	out.CreatedSource = arg.CreatedSource     // simple assign
	out.CreatedName = arg.CreatedName         // simple assign
	out.UpdatedBy = arg.UpdatedBy             // simple assign
	out.ConfirmedBy = arg.ConfirmedBy         // simple assign
	out.ClosedBy = arg.ClosedBy               // simple assign
	out.CreatedAt = arg.CreatedAt             // simple assign
	out.UpdatedAt = arg.UpdatedAt             // simple assign
	out.ConfirmedAt = arg.ConfirmedAt         // simple assign
	out.ClosedAt = arg.ClosedAt               // simple assign
	out.ConnectionID = arg.ConnectionID       // simple assign
	out.Type = arg.Type                       // simple assign
}

func Convert_ticketmodel_Tickets_ticket_Tickets(args []*ticketmodel.Ticket) (outs []*ticket.Ticket) {
	if args == nil {
		return nil
	}
	tmps := make([]ticket.Ticket, len(args))
	outs = make([]*ticket.Ticket, len(args))
	for i := range tmps {
		outs[i] = Convert_ticketmodel_Ticket_ticket_Ticket(args[i], &tmps[i])
	}
	return outs
}

func Convert_ticket_Ticket_ticketmodel_Ticket(arg *ticket.Ticket, out *ticketmodel.Ticket) *ticketmodel.Ticket {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticketmodel.Ticket{}
	}
	convert_ticket_Ticket_ticketmodel_Ticket(arg, out)
	return out
}

func convert_ticket_Ticket_ticketmodel_Ticket(arg *ticket.Ticket, out *ticketmodel.Ticket) {
	out.ID = arg.ID                           // simple assign
	out.Code = arg.Code                       // simple assign
	out.AssignedUserIDs = arg.AssignedUserIDs // simple assign
	out.AccountID = arg.AccountID             // simple assign
	out.LabelIDs = arg.LabelIDs               // simple assign
	out.ExternalID = arg.ExternalID           // simple assign
	out.RefTicketID = arg.RefTicketID         // simple assign
	out.Title = arg.Title                     // simple assign
	out.Description = arg.Description         // simple assign
	out.Note = arg.Note                       // simple assign
	out.RefID = arg.RefID                     // simple assign
	out.RefType = arg.RefType                 // simple assign
	out.RefCode = arg.RefCode                 // simple assign
	out.Source = arg.Source                   // simple assign
	out.State = arg.State                     // simple assign
	out.Status = arg.Status                   // simple assign
	out.CreatedBy = arg.CreatedBy             // simple assign
	out.CreatedSource = arg.CreatedSource     // simple assign
	out.CreatedName = arg.CreatedName         // simple assign
	out.UpdatedBy = arg.UpdatedBy             // simple assign
	out.ConfirmedBy = arg.ConfirmedBy         // simple assign
	out.ClosedBy = arg.ClosedBy               // simple assign
	out.CreatedAt = arg.CreatedAt             // simple assign
	out.UpdatedAt = arg.UpdatedAt             // simple assign
	out.ConfirmedAt = arg.ConfirmedAt         // simple assign
	out.ClosedAt = arg.ClosedAt               // simple assign
	out.WLPartnerID = 0                       // zero value
	out.ConnectionID = arg.ConnectionID       // simple assign
	out.Type = arg.Type                       // simple assign
}

func Convert_ticket_Tickets_ticketmodel_Tickets(args []*ticket.Ticket) (outs []*ticketmodel.Ticket) {
	if args == nil {
		return nil
	}
	tmps := make([]ticketmodel.Ticket, len(args))
	outs = make([]*ticketmodel.Ticket, len(args))
	for i := range tmps {
		outs[i] = Convert_ticket_Ticket_ticketmodel_Ticket(args[i], &tmps[i])
	}
	return outs
}

func Apply_ticket_CreateTicketArgs_ticket_Ticket(arg *ticket.CreateTicketArgs, out *ticket.Ticket) *ticket.Ticket {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticket.Ticket{}
	}
	createTicket(arg, out)
	return out
}

func apply_ticket_CreateTicketArgs_ticket_Ticket(arg *ticket.CreateTicketArgs, out *ticket.Ticket) {
	out.ExternalShippingCode = ""             // zero value
	out.ExternalID = ""                       // zero value
	out.ID = 0                                // zero value
	out.Code = ""                             // zero value
	out.AssignedUserIDs = arg.AssignedUserIDs // simple assign
	out.AccountID = arg.AccountID             // simple assign
	out.LabelIDs = arg.LabelIDs               // simple assign
	out.RefTicketID = arg.RefTicketID         // simple assign
	out.Title = arg.Title                     // simple assign
	out.Description = arg.Description         // simple assign
	out.Note = arg.Note                       // simple assign
	out.AdminNote = arg.AdminNote             // simple assign
	out.RefID = arg.RefID                     // simple assign
	out.RefType = arg.RefType                 // simple assign
	out.RefCode = arg.RefCode                 // simple assign
	out.Source = arg.Source                   // simple assign
	out.State = 0                             // zero value
	out.Status = 0                            // zero value
	out.CreatedBy = arg.CreatedBy             // simple assign
	out.CreatedSource = arg.CreatedSource     // simple assign
	out.CreatedName = arg.CreatedName         // simple assign
	out.UpdatedBy = 0                         // zero value
	out.ConfirmedBy = 0                       // zero value
	out.ClosedBy = 0                          // zero value
	out.CreatedAt = time.Time{}               // zero value
	out.UpdatedAt = time.Time{}               // zero value
	out.ConfirmedAt = time.Time{}             // zero value
	out.ClosedAt = time.Time{}                // zero value
	out.ConnectionID = 0                      // zero value
	out.Type = arg.Type                       // simple assign
}

func Apply_ticket_UpdateTicketInfoArgs_ticket_Ticket(arg *ticket.UpdateTicketInfoArgs, out *ticket.Ticket) *ticket.Ticket {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticket.Ticket{}
	}
	apply_ticket_UpdateTicketInfoArgs_ticket_Ticket(arg, out)
	return out
}

func apply_ticket_UpdateTicketInfoArgs_ticket_Ticket(arg *ticket.UpdateTicketInfoArgs, out *ticket.Ticket) {
	out.ExternalShippingCode = out.ExternalShippingCode      // no change
	out.ExternalID = out.ExternalID                          // no change
	out.ID = arg.ID                                          // simple assign
	out.Code = arg.Code                                      // simple assign
	out.AssignedUserIDs = out.AssignedUserIDs                // no change
	out.AccountID = arg.AccountID                            // simple assign
	out.LabelIDs = out.LabelIDs                              // no change
	out.RefTicketID = out.RefTicketID                        // types do not match
	out.Title = arg.Title.Apply(out.Title)                   // apply change
	out.Description = arg.Description.Apply(out.Description) // apply change
	out.Note = arg.Note.Apply(out.Note)                      // apply change
	out.AdminNote = out.AdminNote                            // no change
	out.RefID = arg.RefID.Apply(out.RefID)                   // apply change
	out.RefType = arg.RefType.Apply(out.RefType)             // apply change
	out.RefCode = out.RefCode                                // no change
	out.Source = arg.Source.Apply(out.Source)                // apply change
	out.State = out.State                                    // types do not match
	out.Status = status5.Status(arg.Status)                  // simple conversion
	out.CreatedBy = out.CreatedBy                            // no change
	out.CreatedSource = out.CreatedSource                    // no change
	out.CreatedName = out.CreatedName                        // no change
	out.UpdatedBy = out.UpdatedBy                            // no change
	out.ConfirmedBy = out.ConfirmedBy                        // no change
	out.ClosedBy = out.ClosedBy                              // no change
	out.CreatedAt = out.CreatedAt                            // no change
	out.UpdatedAt = out.UpdatedAt                            // no change
	out.ConfirmedAt = out.ConfirmedAt                        // no change
	out.ClosedAt = out.ClosedAt                              // no change
	out.ConnectionID = out.ConnectionID                      // no change
	out.Type = out.Type                                      // no change
}

//-- convert o.o/api/supporting/ticket.TicketComment --//

func Convert_ticketmodel_TicketComment_ticket_TicketComment(arg *ticketmodel.TicketComment, out *ticket.TicketComment) *ticket.TicketComment {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticket.TicketComment{}
	}
	convert_ticketmodel_TicketComment_ticket_TicketComment(arg, out)
	return out
}

func convert_ticketmodel_TicketComment_ticket_TicketComment(arg *ticketmodel.TicketComment, out *ticket.TicketComment) {
	out.ID = arg.ID                       // simple assign
	out.TicketID = arg.TicketID           // simple assign
	out.CreatedBy = arg.CreatedBy         // simple assign
	out.CreatedName = arg.CreatedName     // simple assign
	out.CreatedSource = arg.CreatedSource // simple assign
	out.AccountID = arg.AccountID         // simple assign
	out.ParentID = arg.ParentID           // simple assign
	out.ExternalCreatedAt = ""            // types do not match
	out.Message = arg.Message             // simple assign
	out.ImageUrls = arg.ImageUrls         // simple assign
	out.DeletedAt = arg.DeletedAt         // simple assign
	out.DeletedBy = arg.DeletedBy         // simple assign
	out.CreatedAt = arg.CreatedAt         // simple assign
	out.UpdatedAt = arg.UpdatedAt         // simple assign
}

func Convert_ticketmodel_TicketComments_ticket_TicketComments(args []*ticketmodel.TicketComment) (outs []*ticket.TicketComment) {
	if args == nil {
		return nil
	}
	tmps := make([]ticket.TicketComment, len(args))
	outs = make([]*ticket.TicketComment, len(args))
	for i := range tmps {
		outs[i] = Convert_ticketmodel_TicketComment_ticket_TicketComment(args[i], &tmps[i])
	}
	return outs
}

func Convert_ticket_TicketComment_ticketmodel_TicketComment(arg *ticket.TicketComment, out *ticketmodel.TicketComment) *ticketmodel.TicketComment {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticketmodel.TicketComment{}
	}
	convert_ticket_TicketComment_ticketmodel_TicketComment(arg, out)
	return out
}

func convert_ticket_TicketComment_ticketmodel_TicketComment(arg *ticket.TicketComment, out *ticketmodel.TicketComment) {
	out.ID = arg.ID                       // simple assign
	out.TicketID = arg.TicketID           // simple assign
	out.CreatedBy = arg.CreatedBy         // simple assign
	out.CreatedSource = arg.CreatedSource // simple assign
	out.CreatedName = arg.CreatedName     // simple assign
	out.AccountID = arg.AccountID         // simple assign
	out.ParentID = arg.ParentID           // simple assign
	out.ExternalID = ""                   // zero value
	out.ExternalCreatedAt = time.Time{}   // types do not match
	out.Message = arg.Message             // simple assign
	out.ImageUrls = arg.ImageUrls         // simple assign
	out.DeletedAt = arg.DeletedAt         // simple assign
	out.DeletedBy = arg.DeletedBy         // simple assign
	out.CreatedAt = arg.CreatedAt         // simple assign
	out.UpdatedAt = arg.UpdatedAt         // simple assign
}

func Convert_ticket_TicketComments_ticketmodel_TicketComments(args []*ticket.TicketComment) (outs []*ticketmodel.TicketComment) {
	if args == nil {
		return nil
	}
	tmps := make([]ticketmodel.TicketComment, len(args))
	outs = make([]*ticketmodel.TicketComment, len(args))
	for i := range tmps {
		outs[i] = Convert_ticket_TicketComment_ticketmodel_TicketComment(args[i], &tmps[i])
	}
	return outs
}

func Apply_ticket_CreateTicketCommentArgs_ticket_TicketComment(arg *ticket.CreateTicketCommentArgs, out *ticket.TicketComment) *ticket.TicketComment {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticket.TicketComment{}
	}
	createComment(arg, out)
	return out
}

func apply_ticket_CreateTicketCommentArgs_ticket_TicketComment(arg *ticket.CreateTicketCommentArgs, out *ticket.TicketComment) {
	out.ID = 0                            // zero value
	out.TicketID = arg.TicketID           // simple assign
	out.CreatedBy = arg.CreatedBy         // simple assign
	out.CreatedName = arg.CreatedName     // simple assign
	out.CreatedSource = arg.CreatedSource // simple assign
	out.AccountID = arg.AccountID         // simple assign
	out.ParentID = arg.ParentID           // simple assign
	out.ExternalCreatedAt = ""            // zero value
	out.Message = arg.Message             // simple assign
	out.ImageUrls = arg.ImageUrls         // simple assign
	out.DeletedAt = time.Time{}           // zero value
	out.DeletedBy = 0                     // zero value
	out.CreatedAt = time.Time{}           // zero value
	out.UpdatedAt = time.Time{}           // zero value
}

func Apply_ticket_CreateTicketCommentWebhookArgs_ticket_TicketComment(arg *ticket.CreateTicketCommentWebhookArgs, out *ticket.TicketComment) *ticket.TicketComment {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticket.TicketComment{}
	}
	apply_ticket_CreateTicketCommentWebhookArgs_ticket_TicketComment(arg, out)
	return out
}

func apply_ticket_CreateTicketCommentWebhookArgs_ticket_TicketComment(arg *ticket.CreateTicketCommentWebhookArgs, out *ticket.TicketComment) {
	out.ID = 0                                    // zero value
	out.TicketID = arg.TicketID                   // simple assign
	out.CreatedBy = 0                             // zero value
	out.CreatedName = ""                          // zero value
	out.CreatedSource = 0                         // zero value
	out.AccountID = arg.AccountID                 // simple assign
	out.ParentID = arg.ParentID                   // simple assign
	out.ExternalCreatedAt = arg.ExternalCreatedAt // simple assign
	out.Message = arg.Message                     // simple assign
	out.ImageUrls = nil                           // zero value
	out.DeletedAt = time.Time{}                   // zero value
	out.DeletedBy = 0                             // zero value
	out.CreatedAt = arg.CreatedAt                 // simple assign
	out.UpdatedAt = arg.UpdatedAt                 // simple assign
}

func Apply_ticket_UpdateTicketCommentArgs_ticket_TicketComment(arg *ticket.UpdateTicketCommentArgs, out *ticket.TicketComment) *ticket.TicketComment {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticket.TicketComment{}
	}
	apply_ticket_UpdateTicketCommentArgs_ticket_TicketComment(arg, out)
	return out
}

func apply_ticket_UpdateTicketCommentArgs_ticket_TicketComment(arg *ticket.UpdateTicketCommentArgs, out *ticket.TicketComment) {
	out.ID = out.ID                               // identifier
	out.TicketID = out.TicketID                   // identifier
	out.CreatedBy = out.CreatedBy                 // no change
	out.CreatedName = out.CreatedName             // no change
	out.CreatedSource = out.CreatedSource         // no change
	out.AccountID = out.AccountID                 // identifier
	out.ParentID = out.ParentID                   // no change
	out.ExternalCreatedAt = out.ExternalCreatedAt // no change
	out.Message = arg.Message                     // simple assign
	out.ImageUrls = arg.ImageUrls                 // simple assign
	out.DeletedAt = out.DeletedAt                 // no change
	out.DeletedBy = out.DeletedBy                 // no change
	out.CreatedAt = out.CreatedAt                 // no change
	out.UpdatedAt = out.UpdatedAt                 // no change
}

//-- convert o.o/api/supporting/ticket.TicketLabel --//

func Convert_ticketmodel_TicketLabel_ticket_TicketLabel(arg *ticketmodel.TicketLabel, out *ticket.TicketLabel) *ticket.TicketLabel {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticket.TicketLabel{}
	}
	convert_ticketmodel_TicketLabel_ticket_TicketLabel(arg, out)
	return out
}

func convert_ticketmodel_TicketLabel_ticket_TicketLabel(arg *ticketmodel.TicketLabel, out *ticket.TicketLabel) {
	out.ID = arg.ID               // simple assign
	out.ShopID = arg.ShopID       // simple assign
	out.Type = arg.Type           // simple assign
	out.Name = arg.Name           // simple assign
	out.Code = arg.Code           // simple assign
	out.ParentID = arg.ParentID   // simple assign
	out.Color = arg.Color         // simple assign
	out.Children = nil            // zero value
	out.CreatedAt = arg.CreatedAt // simple assign
	out.UpdatedAt = arg.UpdatedAt // simple assign
}

func Convert_ticketmodel_TicketLabels_ticket_TicketLabels(args []*ticketmodel.TicketLabel) (outs []*ticket.TicketLabel) {
	if args == nil {
		return nil
	}
	tmps := make([]ticket.TicketLabel, len(args))
	outs = make([]*ticket.TicketLabel, len(args))
	for i := range tmps {
		outs[i] = Convert_ticketmodel_TicketLabel_ticket_TicketLabel(args[i], &tmps[i])
	}
	return outs
}

func Convert_ticket_TicketLabel_ticketmodel_TicketLabel(arg *ticket.TicketLabel, out *ticketmodel.TicketLabel) *ticketmodel.TicketLabel {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticketmodel.TicketLabel{}
	}
	convert_ticket_TicketLabel_ticketmodel_TicketLabel(arg, out)
	return out
}

func convert_ticket_TicketLabel_ticketmodel_TicketLabel(arg *ticket.TicketLabel, out *ticketmodel.TicketLabel) {
	out.ID = arg.ID               // simple assign
	out.ShopID = arg.ShopID       // simple assign
	out.Type = arg.Type           // simple assign
	out.Name = arg.Name           // simple assign
	out.Code = arg.Code           // simple assign
	out.Color = arg.Color         // simple assign
	out.ParentID = arg.ParentID   // simple assign
	out.CreatedAt = arg.CreatedAt // simple assign
	out.UpdatedAt = arg.UpdatedAt // simple assign
	out.DeletedAt = time.Time{}   // zero value
	out.WLPartnerID = 0           // zero value
}

func Convert_ticket_TicketLabels_ticketmodel_TicketLabels(args []*ticket.TicketLabel) (outs []*ticketmodel.TicketLabel) {
	if args == nil {
		return nil
	}
	tmps := make([]ticketmodel.TicketLabel, len(args))
	outs = make([]*ticketmodel.TicketLabel, len(args))
	for i := range tmps {
		outs[i] = Convert_ticket_TicketLabel_ticketmodel_TicketLabel(args[i], &tmps[i])
	}
	return outs
}

func Apply_ticket_CreateTicketLabelArgs_ticket_TicketLabel(arg *ticket.CreateTicketLabelArgs, out *ticket.TicketLabel) *ticket.TicketLabel {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticket.TicketLabel{}
	}
	createLabel(arg, out)
	return out
}

func apply_ticket_CreateTicketLabelArgs_ticket_TicketLabel(arg *ticket.CreateTicketLabelArgs, out *ticket.TicketLabel) {
	out.ID = 0                  // zero value
	out.ShopID = arg.ShopID     // simple assign
	out.Type = arg.Type         // simple assign
	out.Name = arg.Name         // simple assign
	out.Code = arg.Code         // simple assign
	out.ParentID = arg.ParentID // simple assign
	out.Color = arg.Color       // simple assign
	out.Children = nil          // zero value
	out.CreatedAt = time.Time{} // zero value
	out.UpdatedAt = time.Time{} // zero value
}

func Apply_ticket_UpdateTicketLabelArgs_ticket_TicketLabel(arg *ticket.UpdateTicketLabelArgs, out *ticket.TicketLabel) *ticket.TicketLabel {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticket.TicketLabel{}
	}
	apply_ticket_UpdateTicketLabelArgs_ticket_TicketLabel(arg, out)
	return out
}

func apply_ticket_UpdateTicketLabelArgs_ticket_TicketLabel(arg *ticket.UpdateTicketLabelArgs, out *ticket.TicketLabel) {
	out.ID = out.ID                                 // identifier
	out.ShopID = arg.ShopID                         // simple assign
	out.Type = arg.Type                             // simple assign
	out.Name = arg.Name.Apply(out.Name)             // apply change
	out.Code = arg.Code.Apply(out.Code)             // apply change
	out.ParentID = arg.ParentID.Apply(out.ParentID) // apply change
	out.Color = arg.Color                           // simple assign
	out.Children = out.Children                     // no change
	out.CreatedAt = out.CreatedAt                   // no change
	out.UpdatedAt = out.UpdatedAt                   // no change
}

//-- convert o.o/api/supporting/ticket.TicketLabelExternal --//

func Convert_ticketmodel_TicketLabelExternal_ticket_TicketLabelExternal(arg *ticketmodel.TicketLabelExternal, out *ticket.TicketLabelExternal) *ticket.TicketLabelExternal {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticket.TicketLabelExternal{}
	}
	convert_ticketmodel_TicketLabelExternal_ticket_TicketLabelExternal(arg, out)
	return out
}

func convert_ticketmodel_TicketLabelExternal_ticket_TicketLabelExternal(arg *ticketmodel.TicketLabelExternal, out *ticket.TicketLabelExternal) {
	out.ID = arg.ID                     // simple assign
	out.ConnectionID = arg.ConnectionID // simple assign
	out.ExternalID = arg.ExternalID     // simple assign
	out.ExternalName = arg.ExternalName // simple assign
	out.CreatedAt = arg.CreatedAt       // simple assign
	out.UpdatedAt = arg.UpdatedAt       // simple assign
}

func Convert_ticketmodel_TicketLabelExternals_ticket_TicketLabelExternals(args []*ticketmodel.TicketLabelExternal) (outs []*ticket.TicketLabelExternal) {
	if args == nil {
		return nil
	}
	tmps := make([]ticket.TicketLabelExternal, len(args))
	outs = make([]*ticket.TicketLabelExternal, len(args))
	for i := range tmps {
		outs[i] = Convert_ticketmodel_TicketLabelExternal_ticket_TicketLabelExternal(args[i], &tmps[i])
	}
	return outs
}

func Convert_ticket_TicketLabelExternal_ticketmodel_TicketLabelExternal(arg *ticket.TicketLabelExternal, out *ticketmodel.TicketLabelExternal) *ticketmodel.TicketLabelExternal {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticketmodel.TicketLabelExternal{}
	}
	convert_ticket_TicketLabelExternal_ticketmodel_TicketLabelExternal(arg, out)
	return out
}

func convert_ticket_TicketLabelExternal_ticketmodel_TicketLabelExternal(arg *ticket.TicketLabelExternal, out *ticketmodel.TicketLabelExternal) {
	out.ID = arg.ID                     // simple assign
	out.ConnectionID = arg.ConnectionID // simple assign
	out.ExternalID = arg.ExternalID     // simple assign
	out.ExternalName = arg.ExternalName // simple assign
	out.CreatedAt = arg.CreatedAt       // simple assign
	out.UpdatedAt = arg.UpdatedAt       // simple assign
	out.DeletedAt = time.Time{}         // zero value
}

func Convert_ticket_TicketLabelExternals_ticketmodel_TicketLabelExternals(args []*ticket.TicketLabelExternal) (outs []*ticketmodel.TicketLabelExternal) {
	if args == nil {
		return nil
	}
	tmps := make([]ticketmodel.TicketLabelExternal, len(args))
	outs = make([]*ticketmodel.TicketLabelExternal, len(args))
	for i := range tmps {
		outs[i] = Convert_ticket_TicketLabelExternal_ticketmodel_TicketLabelExternal(args[i], &tmps[i])
	}
	return outs
}

func Apply_ticket_CreateTicketLabelExternalArgs_ticket_TicketLabelExternal(arg *ticket.CreateTicketLabelExternalArgs, out *ticket.TicketLabelExternal) *ticket.TicketLabelExternal {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticket.TicketLabelExternal{}
	}
	apply_ticket_CreateTicketLabelExternalArgs_ticket_TicketLabelExternal(arg, out)
	return out
}

func apply_ticket_CreateTicketLabelExternalArgs_ticket_TicketLabelExternal(arg *ticket.CreateTicketLabelExternalArgs, out *ticket.TicketLabelExternal) {
	out.ID = 0                          // zero value
	out.ConnectionID = arg.ConnectionID // simple assign
	out.ExternalID = arg.ExternalID     // simple assign
	out.ExternalName = arg.ExternalName // simple assign
	out.CreatedAt = time.Time{}         // zero value
	out.UpdatedAt = time.Time{}         // zero value
}

func Apply_ticket_UpdateTicketLabelExternalArgs_ticket_TicketLabelExternal(arg *ticket.UpdateTicketLabelExternalArgs, out *ticket.TicketLabelExternal) *ticket.TicketLabelExternal {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &ticket.TicketLabelExternal{}
	}
	apply_ticket_UpdateTicketLabelExternalArgs_ticket_TicketLabelExternal(arg, out)
	return out
}

func apply_ticket_UpdateTicketLabelExternalArgs_ticket_TicketLabelExternal(arg *ticket.UpdateTicketLabelExternalArgs, out *ticket.TicketLabelExternal) {
	out.ID = arg.ID                     // simple assign
	out.ConnectionID = out.ConnectionID // no change
	out.ExternalID = out.ExternalID     // no change
	out.ExternalName = arg.ExternalName // simple assign
	out.CreatedAt = out.CreatedAt       // no change
	out.UpdatedAt = out.UpdatedAt       // no change
}
