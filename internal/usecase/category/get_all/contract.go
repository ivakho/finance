package get_all

import (
	"context"
	"finance/internal/usecase/category"
)

type categoryRepo interface {
	GetAllCategory(ctx context.Context) ([]category.Category, error)
}
