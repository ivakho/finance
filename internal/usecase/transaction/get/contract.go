package get

import (
	"context"
	storagemodel "finance/internal/storage/transaction"
	"finance/internal/usecase/transaction"
)

type transactionRepo interface {
	GetTransaction(ctx context.Context, filter storagemodel.TransactionFilter) ([]transaction.Transaction, error)
}
