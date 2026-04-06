package category

import (
	"context"
	"finance/internal/usecase/category"
	"fmt"
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

func (r *Repository) UpdateCategory(ctx context.Context, id int, name string) error {
	return r.categoryStorage.UpdateCategory(ctx, id, name)
}

func (r *Repository) DeleteCategory(ctx context.Context, id int) error {
	return r.categoryStorage.DeleteCategory(ctx, id)
}
