package update

import "context"

type transactionRepo interface {
	UpdateTransaction(ctx context.Context, id int, amount float64) error
}
