package category

import (
	"context"
	"finance/internal/model"
)

type storage interface {
	AddCategory(ctx context.Context, name string) error
	GetAllCategory(ctx context.Context) ([]model.Category, error)
	UpdateCategory(ctx context.Context, id int, name string) error
	DeleteCategory(ctx context.Context, id int) error
}
