package get_income

import (
	"context"
	"finance/internal/usecase/transaction"
	"fmt"
	"time"
)

func (u *Usecase) GetIncome(ctx context.Context, categoryID int, dateFrom, dateTo time.Time) (transaction.Transactions, error) {
	result, err := u.transactionRepo.GetTransaction(ctx, categoryID, dateFrom, dateTo)
	if err != nil {
		return transaction.Transactions{}, fmt.Errorf("Failed to get income: %w", err)
	}

	var total int64

	var transactions []transaction.Transaction

	for _, v := range result {
		if v.Amount > 0 {
			transaction := transaction.Transaction{
				ID:         v.ID,
				CategoryID: v.CategoryID,
				Amount:     v.Amount,
				CreatedAt:  v.CreatedAt,
				UpdatedAt:  v.UpdatedAt,
			}

			transactions = append(transactions, transaction)

			total += v.Amount
		}
	}

	return transaction.Transactions{
		Value: transactions,
		Total: total,
	}, nil
}
