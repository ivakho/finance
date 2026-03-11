package add

import "context"

type usecaseTransactionAdd interface {
	Add(ctx context.Context, categoryID int, amount int64) error
}
