package get_income

import (
	"context"
	"finance/internal/usecase/transaction"
)

type usecaseGetIncome interface {
	GetIncome(ctx context.Context, categoryID int) (transaction.Transactions, error)
}
