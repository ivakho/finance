package add

import "context"

type transactionRepo interface {
	AddTransaction(ctx context.Context, categoryID int, amount int64) error
}
