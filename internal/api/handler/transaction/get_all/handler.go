package get_all

import (
	"context"
)

type Handler struct {
	ctx                   context.Context
	usecaseGetAllTransaction usecaseGetAllTransaction
}

func New(ctx context.Context, usecaseGetAllTransaction usecaseGetAllTransaction) *Handler {
	return &Handler{
		ctx:                   ctx,
		usecaseGetAllTransaction: usecaseGetAllTransaction,
	}
}
