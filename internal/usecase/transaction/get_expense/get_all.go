package get_expense

import (
	"context"
	"finance/internal/usecase/transaction"
	"fmt"
)

func (u *Usecase) GetExpense(ctx context.Context, categoryID int) (transaction.Transactions, error) {
	result, err := u.transactionRepo.GetExpense(ctx, categoryID)
	if err != nil {
		return transaction.Transactions{}, fmt.Errorf("Failed to get expense: %w", err)
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
