package get_income_total

import (
	"context"
	"finance/internal/usecase/category"
	"time"
)

type usecaseGetIncomeTotal interface {
	GetCategoryIncomeTotal(ctx context.Context, dateFrom, dateTo time.Time) ([]category.CategoryTotal, error)
}
