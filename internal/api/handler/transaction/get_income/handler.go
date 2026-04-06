package get_income

import (
	"context"
)

type Handler struct {
	ctx                   context.Context
	usecaseGetIncome usecaseGetIncome
}

func New(ctx context.Context, usecaseGetIncome usecaseGetIncome) *Handler {
	return &Handler{
		ctx:                   ctx,
		usecaseGetIncome: usecaseGetIncome,
	}
}
