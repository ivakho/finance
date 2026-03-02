package get_all

import (
	"context"
	"finance/internal/usecase/transaction"
)

type transactionRepo interface {
	GetAllTransaction(ctx context.Context, categoryID int) ([]transaction.Transaction, error)
}
