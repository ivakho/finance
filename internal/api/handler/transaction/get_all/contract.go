package get_all

import (
	"context"
	"finance/internal/usecase/transaction"
)

type usecaseGetAllTransaction interface {
	GetAll(ctx context.Context, categoryID int) (transaction.Transactions, error)
}
