package add

import (
	"context"
	"fmt"
)

func (u *Usecase) Add(ctx context.Context, categoryID int, amount float64) error {
	if err := u.transactionRepo.AddTransaction(ctx, categoryID, amount); err != nil {
		return fmt.Errorf("Failed to add transaction: %w", err)
	}

	return nil
}
