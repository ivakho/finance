package get_all

import (
	"context"
)

type Handler struct {
	ctx                   context.Context
	usecaseGetAllCategory usecaseGetAllCategory
}

func New(ctx context.Context, usecaseGetAllCategory usecaseGetAllCategory) *Handler {
	return &Handler{
		ctx:                   ctx,
		usecaseGetAllCategory: usecaseGetAllCategory,
	}
}
