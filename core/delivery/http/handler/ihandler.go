package handler

import "context"

type RequestHandler interface {
	Handle(ctx context.Context) (interface{}, error)
}
