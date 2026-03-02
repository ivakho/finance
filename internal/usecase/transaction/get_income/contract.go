package get_income

import (
	"context"
	"finance/internal/usecase/transaction"
)

type transactionRepo interface {
	GetIncome(ctx context.Context, categoryID int) ([]transaction.Transaction, error)
}
