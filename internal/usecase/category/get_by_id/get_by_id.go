package get_by_id

import (
	"context"
	"finance/internal/usecase/category"
	"fmt"
)

func (u *Usecase) GetCategoryByID(ctx context.Context, id int) (category.Category, error) {
	categoryByID, err := u.categoryRepo.GetCategoryByID(ctx, id)
	if err != nil {
		return category.Category{}, fmt.Errorf("Failed to get category by ID: %w", err)
	}

	return categoryByID, nil
}
