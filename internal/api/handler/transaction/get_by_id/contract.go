package get_by_id

import (
	"context"
	"finance/internal/usecase/transaction"
)

type usecaseGetTransactionByID interface {
	GetTransactionByID(ctx context.Context, id int) (transaction.TransactionByID, error)
}
