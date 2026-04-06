package get_all

import (
	"context"
	"finance/internal/usecase/category"
)

type usecaseGetAllCategory interface {
	GetAll(ctx context.Context) ([]category.Category, error)
}
