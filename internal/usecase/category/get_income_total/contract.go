package get_income_total

import (
	"context"
	"finance/internal/usecase/category"
	"time"
)

type categoryRepo interface {
	GetCategoryIncomeTotal(ctx context.Context, dateFrom, dateTo time.Time) ([]category.CategoryTotal, error)
}
