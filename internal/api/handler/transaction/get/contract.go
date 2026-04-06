package get

import (
	"context"
	storagemodel "finance/internal/storage/transaction"
	"finance/internal/usecase/transaction"
)

type usecaseGetTransaction interface {
	Get(ctx context.Context, filter storagemodel.TransactionFilter) (transaction.Transactions, error)
}
