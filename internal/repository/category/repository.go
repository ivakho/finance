package category

import (
	"context"
	modelcategory "finance/internal/model"
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

func (r *Repository) GetAllCategory(ctx context.Context) ([]modelcategory.Category, error) {
	return r.categoryStorage.GetAllCategory(ctx)
}

func (r *Repository) UpdateCategory(ctx context.Context, id int, name string) error {
	return r.categoryStorage.UpdateCategory(ctx, id, name)
}

func (r *Repository) DeleteCategory(ctx context.Context, id int) error {
	return r.categoryStorage.DeleteCategory(ctx, id)
}
