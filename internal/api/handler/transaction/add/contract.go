package add

import "context"

type usecaseTransactionAdd interface {
	Add(ctx context.Context, categoryID int, txType string, amount int64) error
}
