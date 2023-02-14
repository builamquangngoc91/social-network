package handlers

import "context"

type CommentHandler struct {}

func NewCommentHandler() *CommentHandler {
	return &CommentHandler{}
}

func (h *CommentHandler) Handle(ctx context.Context, msg []byte) error {
	return nil
}