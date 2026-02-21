package update

import "context"

type Handler struct {
	ctx                   context.Context
	usecaseUpdateCategory usecaseUpdateCategory
}

func New(ctx context.Context, usecaseUpdateCategory usecaseUpdateCategory) *Handler {
	return &Handler{
		ctx:                   ctx,
		usecaseUpdateCategory: usecaseUpdateCategory,
	}
}
