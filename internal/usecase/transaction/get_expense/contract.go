package get_expense

import (
	"context"
	"finance/internal/usecase/transaction"
	"time"
)

type transactionRepo interface {
	GetTransaction(ctx context.Context, categoryID int, dateFrom, dateTo time.Time) ([]transaction.Transaction, error)
}
