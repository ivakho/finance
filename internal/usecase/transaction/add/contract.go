package add

import (
	"context"
	"time"
)

type transactionRepo interface {
	AddTransaction(ctx context.Context, categoryID int, amount int64, createdAt time.Time) error
}
