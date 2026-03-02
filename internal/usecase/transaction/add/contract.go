package add

import "context"

type transactionRepo interface {
	AddTransaction(ctx context.Context, categoryID int, txType string, amount int64) error
}
