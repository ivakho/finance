package add

import (
	"context"
)

type Handler struct {
	ctx                   context.Context
	usecaseTransactionAdd usecaseTransactionAdd
}

func New(ctx context.Context, usecaseTransactionAdd usecaseTransactionAdd) *Handler {
	return &Handler{
		ctx:                   ctx,
		usecaseTransactionAdd: usecaseTransactionAdd,
	}
}
