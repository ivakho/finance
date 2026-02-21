package get

import (
	"context"
)

type Handler struct {
	ctx                   context.Context
	usecaseGetTransaction usecaseGetTransaction
}

func New(ctx context.Context, usecaseGetTransaction usecaseGetTransaction) *Handler {
	return &Handler{
		ctx:                   ctx,
		usecaseGetTransaction: usecaseGetTransaction,
	}
}
