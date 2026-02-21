package get_all

import (
	"context"
	"finance/internal/model"
)

type transactionRepo interface {
	GetAllTransaction(ctx context.Context, categoryID int) ([]model.Transaction, error)
}
