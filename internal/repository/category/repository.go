package category

import (
	"context"
	"finance/internal/usecase/category"
	"fmt"
	"time"
)

type Repository struct {
	categoryStorage storage
}

func New(categoryStorage storage) *Repository {
	return &Repository{categoryStorage: categoryStorage}
}

func (r *Repository) AddCategory(ctx context.Context, name string) error {
	return r.categoryStorage.AddCategory(ctx, name)
}

func (r *Repository) GetCategoryByID(ctx context.Context, id int) (category.Category, error) {
	result, err := r.categoryStorage.GetCategoryByID(ctx, id)
	if err != nil {
		return category.Category{}, fmt.Errorf("GetCategoryByID: %w", err)
	}

	return category.Category{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (r *Repository) GetAllCategory(ctx context.Context) ([]category.Category, error) {
	result, err := r.categoryStorage.GetAllCategory(ctx)
	if err != nil {
		return nil, fmt.Errorf("GetAlCategory: %w", err)
	}

	categories := make([]category.Category, 0, len(result))

	for _, v := range result {
		categories = append(categories, category.Category{
			ID:        v.ID,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return categories, nil
}

func (r *Repository) GetCategoryIncomeTotal(ctx context.Context, dateFrom, dateTo time.Time) ([]category.CategoryTotal, error) {
	result, err := r.categoryStorage.GetCategoryIncomeTotal(ctx, dateFrom, dateTo)
	if err != nil {
		return nil, fmt.Errorf("GetCategoryIncomeTotal: %w", err)
	}

	categories := make([]category.CategoryTotal, 0, len(result))

	for _, v := range result {
		categories = append(categories, category.CategoryTotal{
			ID:    v.ID,
			Name:  v.Name,
			Total: v.Total,
		})
	}

	return categories, nil
}

func (r *Repository) GetCategoryExpenseTotal(ctx context.Context, dateFrom, dateTo time.Time) ([]category.CategoryTotal, error) {
	result, err := r.categoryStorage.GetCategoryExpenseTotal(ctx, dateFrom, dateTo)
	if err != nil {
		return nil, fmt.Errorf("GetCategoryExpenseTotal: %w", err)
	}

	categories := make([]category.CategoryTotal, 0, len(result))

	for _, v := range result {
		categories = append(categories, category.CategoryTotal{
			ID:    v.ID,
			Name:  v.Name,
			Total: v.Total,
		})
	}

	return categories, nil
}

func (r *Repository) UpdateCategory(ctx context.Context, id int, name string) error {
	return r.categoryStorage.UpdateCategory(ctx, id, name)
}

func (r *Repository) DeleteCategory(ctx context.Context, id int) error {
	return r.categoryStorage.DeleteCategory(ctx, id)
}
