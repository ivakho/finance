package delete

import (
	"context"
	"fmt"
)

func (u *Usecase) Delete(ctx context.Context, id int) error {
	if err := u.transactionRepo.DeleteTransaction(ctx, id); err != nil {
		return fmt.Errorf("Failed to delete transaction: %w", err)
	}

	return nil
}
