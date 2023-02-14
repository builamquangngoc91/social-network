package handlers

import "context"

type EventHandler struct{}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

func (h *EventHandler) Handle(ctx context.Context, msg []byte) error {
	return nil
}
