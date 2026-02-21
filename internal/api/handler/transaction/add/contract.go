package add

import "context"

type usecaseTransactionAdd interface {
	Add(ctx context.Context, categoryID int, amount float64) error
}
