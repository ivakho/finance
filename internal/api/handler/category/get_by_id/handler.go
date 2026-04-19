package get_by_id

import "context"

type Handler struct {
	ctx                    context.Context
	usecaseGetCategoryByID usecaseGetCategoryByID
}

func New(ctx context.Context, usecaseGetCategoryByID usecaseGetCategoryByID) *Handler {
	return &Handler{
		ctx:                    ctx,
		usecaseGetCategoryByID: usecaseGetCategoryByID,
	}
}
