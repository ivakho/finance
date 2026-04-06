package transaction

import (
	"context"
	"log"
	// storagemodel "finance/internal/storage/transaction"
	usecasemodel "finance/internal/usecase/transaction"
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

func (r *Repository) GetAllTransaction(ctx context.Context, categoryID int, createdAt time.Time) ([]usecasemodel.Transaction, error) {
	result, err := r.transactionStorage.GetAllTransaction(ctx, categoryID, createdAt)
	if err != nil {
		return nil, fmt.Errorf("GetAllTransaction: %w", err)
	}
	transactions := make([]usecasemodel.Transaction, 0, len(result))

	for _, v := range result {
		transactions = append(transactions, usecasemodel.Transaction{ID: v.ID,
			CategoryID: v.CategoryID,
			Amount:     v.Amount,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  *v.UpdatedAt,
		})
	}

	return transactions, nil
}

func (r *Repository) GetIncome(ctx context.Context, categoryID int, createdAt time.Time) ([]usecasemodel.Transaction, error) {
	result, err := r.transactionStorage.GetIncome(ctx, categoryID, createdAt)
	if err != nil {
		return nil, fmt.Errorf("GetIncome:%w", err)
	}
	transactions := make([]usecasemodel.Transaction, 0, len(result))

	for _, v := range result {
		transactions = append(transactions, usecasemodel.Transaction{
			ID:         v.ID,
			CategoryID: v.CategoryID,
			Amount:     v.Amount,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  *v.UpdatedAt,
		})
	}

	return transactions, nil
}

func (r *Repository) GetExpense(ctx context.Context, categoryID int, createdAt time.Time) ([]usecasemodel.Transaction, error) {
	result, err := r.transactionStorage.GetExpense(ctx, categoryID, createdAt)
	if err != nil {
		return nil, fmt.Errorf("GetExpense:%w", err)
	}
	transactions := make([]usecasemodel.Transaction, 0, len(result))

	for _, v := range result {
		transactions = append(transactions, usecasemodel.Transaction{
			ID:         v.ID,
			CategoryID: v.CategoryID,
			Amount:     v.Amount,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  *v.UpdatedAt,
		})
	}

	return transactions, nil
}

func (r *Repository) GetTransaction(ctx context.Context, categoryID int, dateFrom, dateTo time.Time) ([]usecasemodel.Transaction, error) {
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
	transactions := make([]usecasemodel.Transaction, 0, len(result))

	for _, v := range result {
		// transactions = append(transactions, usecasemodel.Transaction{ID: v.ID,
		// 	CategoryID: v.CategoryID,
		// 	Amount:     v.Amount,
		// 	CreatedAt:  v.CreatedAt,
		// 	UpdatedAt:  *v.UpdatedAt,
		// })
		tx := usecasemodel.Transaction{
			ID:         v.ID,
			CategoryID: v.CategoryID,
			Amount:     v.Amount,
			CreatedAt:  v.CreatedAt,
		}

		if v.UpdatedAt != nil {
			tx.UpdatedAt = *v.UpdatedAt
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
