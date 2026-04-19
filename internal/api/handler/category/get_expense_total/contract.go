package get_expense_total

import (
	"context"
	"finance/internal/usecase/category"
	"time"
)

type usecaseGetExpenseTotal interface {
	GetCategoryExpenseTotal(ctx context.Context, dateFrom, dateTo time.Time) ([]category.CategoryTotal, error)
}
