package get_expense

import (
	"context"
)

type Handler struct {
	ctx               context.Context
	usecaseGetExpense usecaseGetExpense
}

func New(ctx context.Context, usecaseGetExpense usecaseGetExpense) *Handler {
	return &Handler{
		ctx:               ctx,
		usecaseGetExpense: usecaseGetExpense,
	}
}
