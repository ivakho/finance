package get_by_id

import (
	"context"
	"finance/internal/usecase/category"
)

type categoryRepo interface {
	GetCategoryByID(ctx context.Context, id int) (category.Category, error)
}
