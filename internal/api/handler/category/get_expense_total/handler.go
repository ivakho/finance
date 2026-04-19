package get_expense_total

import "context"

type Handler struct {
	ctx                   context.Context
	usecaseGetExpenseTotal usecaseGetExpenseTotal
}

func New(ctx context.Context, usecaseGetExpenseTotal usecaseGetExpenseTotal) *Handler {
	return &Handler{
		ctx:                   ctx,
		usecaseGetExpenseTotal: usecaseGetExpenseTotal,
	}
}
