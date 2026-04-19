package get_expense_total

import (
	"context"
	"finance/internal/usecase/category"
	"fmt"
	"time"
)

func (u *Usecase) GetCategoryExpenseTotal(ctx context.Context, dateFrom, dateTo time.Time) ([]category.CategoryTotal, error) {
	categories, err := u.categoryRepo.GetCategoryExpenseTotal(ctx, dateFrom, dateTo)
	if err != nil {
		return nil, fmt.Errorf("Failed to get all categories: %w", err)
	}

	return categories, nil
}
