package get_income_total

import (
	"context"
	"finance/internal/usecase/category"
	"fmt"
	"time"
)

func (u *Usecase) GetCategoryIncomeTotal(ctx context.Context, dateFrom, dateTo time.Time) ([]category.CategoryTotal, error) {
	categories, err := u.categoryRepo.GetCategoryIncomeTotal(ctx, dateFrom, dateTo)
	if err != nil {
		return nil, fmt.Errorf("Failed to get all categories: %w", err)
	}

	return categories, nil
}
