package get

import (
	"context"
	"finance/internal/usecase/transaction"
	"fmt"
	"time"
)

func (u *Usecase) Get(ctx context.Context, categoryID int, dateFrom, dateTo time.Time) (transaction.Transactions, error) {
	result, err := u.transactionRepo.GetTransaction(ctx, categoryID, dateFrom, dateTo)
	if err != nil {
		return transaction.Transactions{}, fmt.Errorf("Failed to get transaction: %w", err)
	}

	var total int64

	for _, v := range result {
		total += v.Amount
	}

	return transaction.Transactions{
		Value: result,
		Total: total,
	}, nil
}
