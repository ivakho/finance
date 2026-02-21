package transaction

import (
	"context"
	"finance/internal/model"
)

type storage interface {
	AddTransaction(ctx context.Context, categoryID int, amount float64) error
	GetAllTransaction(ctx context.Context, categoryID int) ([]model.Transaction, error)
	GetTransaction(ctx context.Context, id int) (model.Transaction, error)
	UpdateTransaction(ctx context.Context, id int, amount float64) error
	DeleteTransaction(ctx context.Context, id int) error
}
