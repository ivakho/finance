package add

import (
	"context"
	"finance/internal/usecase/transaction"
	"fmt"
	"strings"
)

func (u *Usecase) Add(ctx context.Context, tx transaction.TransactionAdd) error {
	if strings.ToLower(tx.TxType) == "expense" {
		tx.Amount *= -1
	}

	if err := u.transactionRepo.AddTransaction(ctx, tx.CategoryID, tx.Amount, tx.CreatedAt); err != nil {
		return fmt.Errorf("Failed to add transaction: %w", err)
	}

	return nil
}
