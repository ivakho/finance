package delete

import "context"

type transactionRepo interface {
	DeleteTransaction(ctx context.Context, id int) error
}
