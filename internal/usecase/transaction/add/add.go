package add

import (
	"context"
	"finance/internal/usecase/transaction"
	"fmt"
	"strings"
)

const (
	TxTypeExpense = "expense"
	TxTypeIncome  = "income"
)

func (u *Usecase) Add(ctx context.Context, tx transaction.TransactionAdd) error {
	multiplier, err := txMultiplier(tx.TxType)
	if err != nil {
		return err
	}

	tx.Amount *= multiplier

	if err := u.transactionRepo.AddTransaction(ctx, tx.CategoryID, tx.Amount, tx.CreatedAt); err != nil {
		return fmt.Errorf("Failed to add transaction: %w", err)
	}

	return nil
}

func txMultiplier(txType string) (int64, error) {
	switch strings.ToLower(txType) {
	case TxTypeExpense:
		return -1, nil
	case TxTypeIncome:
		return 1, nil
	default:
		return 0, fmt.Errorf("invalid transaction type: %s", txType)
	}
}
