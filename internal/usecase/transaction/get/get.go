package get

import (
	"context"
	"finance/internal/model"
	"fmt"
)

func (u *Usecase) Get(ctx context.Context, id int) (model.Transaction, error) {
	transaction, err := u.transactionRepo.GetTransaction(ctx, id)
	if err != nil {
		return transaction, fmt.Errorf("Failed to get transaction: %w", err)
	}

	return transaction, nil
}
