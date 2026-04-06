package delete

import "context"

type Handler struct {
	ctx                   context.Context
	usecaseDeleteCategory usecaseDeleteCategory
}

func New(ctx context.Context, useusecaseDeleteCategory usecaseDeleteCategory) *Handler {
	return &Handler{ctx: ctx, usecaseDeleteCategory: useusecaseDeleteCategory}
}
