package transaction

import (
	"context"
	"finance/internal/storage/transaction"
	"time"
)

type storage interface {
	AddTransaction(ctx context.Context, categoryID int, amount int64, createdAt time.Time) error
	GetTransactionByID(ctx context.Context, id int) (transaction.TransactionByID, error)
	GetTransaction(ctx context.Context, categoryID int, dateFrom, dateTo time.Time) ([]transaction.Transaction, error)
	UpdateTransaction(ctx context.Context, id int, amount int64) error
	DeleteTransaction(ctx context.Context, id int) error
}
