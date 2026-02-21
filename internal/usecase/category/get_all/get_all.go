package get_all

import (
	"context"
	"finance/internal/model"
	"fmt"
)

func (u *Usecase) GetAll(ctx context.Context) ([]model.Category, error) {
	categories, err := u.categoryRepo.GetAllCategory(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to get all categories: %w", err)
	}

	return categories, nil
}
