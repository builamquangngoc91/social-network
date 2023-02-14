package kafka

import "context"

type Handler interface {
	Handle(context.Context, []byte) error
}

