package category

import (
	"context"
	"finance/internal/storage/category"
	"time"
)

type storage interface {
	AddCategory(ctx context.Context, name string) error
	GetCategoryByID(ctx context.Context, id int) (category.Category, error)
	GetAllCategory(ctx context.Context) ([]category.Category, error)
	GetCategoryIncomeTotal(ctx context.Context, dateFrom, dateTo time.Time) ([]category.CategoryTotal, error)
	GetCategoryExpenseTotal(ctx context.Context, dateFrom, dateTo time.Time) ([]category.CategoryTotal, error)
	UpdateCategory(ctx context.Context, id int, name string) error
	DeleteCategory(ctx context.Context, id int) error
}
