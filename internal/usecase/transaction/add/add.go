package add

import (
	"context"
	"fmt"
)

func (u *Usecase) Add(ctx context.Context, categoryID int, txType string, amount int64) error {
	if txType == "expense" {
		amount = -amount
	}

	if err := u.transactionRepo.AddTransaction(ctx, categoryID, txType, amount); err != nil {
		return fmt.Errorf("Failed to add transaction: %w", err)
	}

	return nil
}
