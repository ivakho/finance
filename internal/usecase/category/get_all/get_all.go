package get_all

import (
	"context"
	"finance/internal/usecase/category"
	"fmt"
)

func (u *Usecase) GetAll(ctx context.Context) ([]category.Category, error) {
	categories, err := u.categoryRepo.GetAllCategory(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to get all categories: %w", err)
	}

	return categories, nil
}
