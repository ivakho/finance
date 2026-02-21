package get_all

import (
	"context"
	"finance/internal/model"
	"fmt"
)

func (u *Usecase) GetAll(ctx context.Context, categoryID int) ([]model.Transaction, error) {
	transactions, err := u.transactionRepo.GetAllTransaction(ctx, categoryID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get all transactions: %w", err)
	}

	return transactions, nil
}
