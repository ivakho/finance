package get_expense

import (
	"context"
	"finance/internal/usecase/transaction"
)

type transactionRepo interface {
	GetExpense(ctx context.Context, categoryID int) ([]transaction.Transaction, error)
}
