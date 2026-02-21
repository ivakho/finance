package get_all

import (
	"context"
	modelcategory "finance/internal/model"
)

type categoryRepo interface {
	GetAllCategory(ctx context.Context) ([]modelcategory.Category, error)
}
