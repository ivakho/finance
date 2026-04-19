package get_by_id

import "context"

type Handler struct {
	ctx                       context.Context
	usecaseGetTransactionByID usecaseGetTransactionByID
}

func New(ctx context.Context, usecaseGetTransactionByID usecaseGetTransactionByID) *Handler {
	return &Handler{
		ctx:                       ctx,
		usecaseGetTransactionByID: usecaseGetTransactionByID,
	}
}
