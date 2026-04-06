package update

import "context"

type Handler struct {
	ctx                      context.Context
	usecaseUpdateTransaction usecaseUpdateTransaction
}

func New(ctx context.Context, usecaseUpdateTransaction usecaseUpdateTransaction) *Handler {
	return &Handler{
		ctx:                      ctx,
		usecaseUpdateTransaction: usecaseUpdateTransaction,
	}
}
