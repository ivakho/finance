package transaction

import (
	"context"
	"finance/internal/storage/transaction"
)

type storage interface {
	AddTransaction(ctx context.Context, categoryID int, txType string, amount int64) error
	GetAllTransaction(ctx context.Context, categoryID int) ([]transaction.Transaction, error)
	GetExpense(ctx context.Context, categoryID int) ([]transaction.Transaction, error)
	GetIncome(ctx context.Context, categoryID int) ([]transaction.Transaction, error)
	GetTransaction(ctx context.Context, filter transaction.TransactionFilter) ([]transaction.Transaction, error)
	UpdateTransaction(ctx context.Context, id int, amount int64) error
	DeleteTransaction(ctx context.Context, id int) error
}
