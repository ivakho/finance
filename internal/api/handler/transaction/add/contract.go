package add

import (
	"context"
	"finance/internal/usecase/transaction"
)

type usecaseTransactionAdd interface {
	Add(ctx context.Context, tx transaction.TransactionAdd) error
}
