package get_income

import (
	"context"
	"finance/internal/usecase/transaction"
	"time"
)

type usecaseGetIncome interface {
	GetIncome(ctx context.Context, categoryID int, dateFrom, dateTo time.Time) (transaction.Transactions, error)
}
