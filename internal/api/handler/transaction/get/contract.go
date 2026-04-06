package get

import (
	"context"
	"finance/internal/usecase/transaction"
	"time"
)

type usecaseGetTransaction interface {
	Get(ctx context.Context, categoryID int, dateFrom, dateTo time.Time) (transaction.Transactions, error)
}
