package get

import (
	"context"
	"finance/internal/model"
)

type transactionRepo interface {
	GetTransaction(ctx context.Context, id int) (model.Transaction, error)
}

