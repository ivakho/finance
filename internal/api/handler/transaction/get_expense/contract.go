package get_expense

import (
	"context"
	"finance/internal/usecase/transaction"
)

type usecaseGetExpense interface {
	GetExpense(ctx context.Context, categoryID int) (transaction.Transactions, error)
}
