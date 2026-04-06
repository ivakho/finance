package delete

import "context"

type Handler struct {
	ctx                      context.Context
	usecaseDeleteTransaction usecaseDeleteTransaction
}

func New(ctx context.Context, usecaseDeleteTransaction usecaseDeleteTransaction) *Handler {
	return &Handler{ctx: ctx, usecaseDeleteTransaction: usecaseDeleteTransaction}
}
