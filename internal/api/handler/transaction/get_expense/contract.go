package get_expense

import (
	"context"
	"finance/internal/usecase/transaction"
	"time"
)

type usecaseGetExpense interface {
	GetExpense(ctx context.Context, categoryID int, dateFrom, dateTo time.Time) (transaction.Transactions, error)
}
