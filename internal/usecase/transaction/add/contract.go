package add

import "context"

type transactionRepo interface {
	AddTransaction(ctx context.Context, categoryID int, amount float64) error
}
