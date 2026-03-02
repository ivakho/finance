package get

import (
	"context"
	storagemodel "finance/internal/storage/transaction"
	"finance/internal/usecase/transaction"
	"fmt"
)

func (u *Usecase) Get(ctx context.Context, filter storagemodel.TransactionFilter) (transaction.Transactions, error) {
	result, err := u.transactionRepo.GetTransaction(ctx, filter)
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
