package transaction

import (
	"context"
	"finance/internal/usecase/transaction"
	"log"
	"time"

	"fmt"
)

type Repository struct {
	transactionStorage storage
}

func New(transactionStorage storage) *Repository {
	return &Repository{transactionStorage: transactionStorage}
}

func (r *Repository) AddTransaction(ctx context.Context, categoryID int, amount int64, createdAt time.Time) error {
	return r.transactionStorage.AddTransaction(ctx, categoryID, amount, createdAt)
}

func (r *Repository) GetTransactionByID(ctx context.Context, id int) (transaction.TransactionByID, error) {
	result, err := r.transactionStorage.GetTransactionByID(ctx, id)
	if err != nil {
		return transaction.TransactionByID{}, fmt.Errorf("GetTransactionByID: %w", err)
	}

	return transaction.TransactionByID{
		ID:         result.ID,
		CategoryID: result.CategoryID,
		CategoryName: result.CategoryName,
		Amount:     result.Amount,
		CreatedAt:  result.CreatedAt,
		UpdatedAt:  result.UpdatedAt,
	}, nil
}

func (r *Repository) GetTransaction(ctx context.Context, categoryID int, dateFrom, dateTo time.Time) ([]transaction.Transaction, error) {
	result, err := r.transactionStorage.GetTransaction(ctx, categoryID, dateFrom, dateTo)

	if dateFrom.IsZero() {
		log.Printf("WARNING: dateFrom is zero")
	}
	if dateTo.IsZero() {
		log.Printf("WARNING: dateTo is zero")
	}

	if err != nil {
		return nil, fmt.Errorf("GetTransaction:%w", err)
	}
	transactions := make([]transaction.Transaction, 0, len(result))

	for _, v := range result {
		tx := transaction.Transaction{
			ID:         v.ID,
			CategoryID: v.CategoryID,
			Amount:     v.Amount,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}

		transactions = append(transactions, tx)

	}

	return transactions, nil
}

func (r *Repository) UpdateTransaction(ctx context.Context, id int, amount int64) error {
	return r.transactionStorage.UpdateTransaction(ctx, id, amount)
}

func (r *Repository) DeleteTransaction(ctx context.Context, id int) error {
	return r.transactionStorage.DeleteTransaction(ctx, id)
}
