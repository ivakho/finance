package add

import (
	"context"
)

type Handler struct {
	ctx                context.Context
	usecaseCategoryAdd usecaseCategoryAdd
}

func New(ctx context.Context, usecaseCategoryAdd usecaseCategoryAdd) *Handler {
	return &Handler{
		ctx:                ctx,
		usecaseCategoryAdd: usecaseCategoryAdd,
	}
}
