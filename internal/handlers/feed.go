package handlers

import "context"

type FeedHandler struct {}

func NewFeedHandler() *FeedHandler {
	return &FeedHandler{}
}

func (h *FeedHandler) Handle(ctx context.Context, msg []byte) error {
	return nil
}