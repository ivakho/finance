package update

import (
	"context"
	"fmt"
)

func (u *Usecase) Update(ctx context.Context, id int, amount float64) error {
	if err := u.transactionRepo.UpdateTransaction(ctx, id, amount); err != nil {
		return fmt.Errorf("Failed to edit transaction: %w", err)
	}

	return nil
}
