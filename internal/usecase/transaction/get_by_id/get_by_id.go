package get_by_id

import (
	"context"
	"finance/internal/usecase/transaction"
	"fmt"
)

func (u *Usecase) GetTransactionByID(ctx context.Context, id int) (transaction.TransactionByID, error) {
	transactionByID, err := u.transactionRepo.GetTransactionByID(ctx, id)
	if err != nil {
		return transaction.TransactionByID{}, fmt.Errorf("Failed to get category by ID: %w", err)
	}

	return transactionByID, nil
}
