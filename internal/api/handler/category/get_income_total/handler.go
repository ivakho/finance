package get_income_total

import "context"

type Handler struct {
	ctx                   context.Context
	usecaseGetIncomeTotal usecaseGetIncomeTotal
}

func New(ctx context.Context, usecaseGetIncomeTotal usecaseGetIncomeTotal) *Handler {
	return &Handler{
		ctx:                   ctx,
		usecaseGetIncomeTotal: usecaseGetIncomeTotal,
	}
}
