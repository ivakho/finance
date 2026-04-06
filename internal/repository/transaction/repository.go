package transaction

import (
	"context"
	storagemodel "finance/internal/storage/transaction"
	usecasemodel "finance/internal/usecase/transaction"

	"fmt"
)

type Repository struct {
	transactionStorage storage
}

func New(transactionStorage storage) *Repository {
	return &Repository{transactionStorage: transactionStorage}
}

func (r *Repository) AddTransaction(ctx context.Context, categoryID int, txType string, amount int64) error {
	return r.transactionStorage.AddTransaction(ctx, categoryID, txType, amount)
}

func (r *Repository) GetAllTransaction(ctx context.Context, categoryID int) ([]usecasemodel.Transaction, error) {
	result, err := r.transactionStorage.GetAllTransaction(ctx, categoryID)
	if err != nil {
		return nil, fmt.Errorf("GetAllTransaction: %w", err)
	}
	transactions := make([]usecasemodel.Transaction, 0, len(result))

	for _, v := range result {
		transactions = append(transactions, usecasemodel.Transaction{ID: v.ID, CategoryID: v.CategoryID, Amount: v.Amount, CreatedAt: v.CreatedAt, UpdatedAt: *v.UpdatedAt})
	}

	return transactions, nil
}

func (r *Repository) GetIncome(ctx context.Context, categoryID int) ([]usecasemodel.Transaction, error) {
	result, err := r.transactionStorage.GetIncome(ctx, categoryID)
	if err != nil {
		return nil, fmt.Errorf("GetIncome:%w", err)
	}
	transactions := make([]usecasemodel.Transaction, 0, len(result))

	for _, v := range result {
		transactions = append(transactions, usecasemodel.Transaction{ID: v.ID, CategoryID: v.CategoryID, Amount: v.Amount, CreatedAt: v.CreatedAt, UpdatedAt: *v.UpdatedAt})
	}

	return transactions, nil
}

func (r *Repository) GetExpense(ctx context.Context, categoryID int) ([]usecasemodel.Transaction, error) {
	result, err := r.transactionStorage.GetExpense(ctx, categoryID)
	if err != nil {
		return nil, fmt.Errorf("GetExpense:%w", err)
	}
	transactions := make([]usecasemodel.Transaction, 0, len(result))

	for _, v := range result {
		transactions = append(transactions, usecasemodel.Transaction{ID: v.ID, CategoryID: v.CategoryID, Amount: v.Amount, CreatedAt: v.CreatedAt, UpdatedAt: *v.UpdatedAt})
	}

	return transactions, nil
}

func (r *Repository) GetTransaction(ctx context.Context, filter storagemodel.TransactionFilter) ([]usecasemodel.Transaction, error) {
	result, err := r.transactionStorage.GetTransaction(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("GetTransaction:%w", err)
	}
	transactions := make([]usecasemodel.Transaction, 0, len(result))

	for _, v := range result {
		transactions = append(transactions, usecasemodel.Transaction{ID: v.ID, CategoryID: v.CategoryID, Amount: v.Amount, CreatedAt: v.CreatedAt, UpdatedAt: *v.UpdatedAt})
	}

	return transactions, nil
}

func (r *Repository) UpdateTransaction(ctx context.Context, id int, amount int64) error {
	return r.transactionStorage.UpdateTransaction(ctx, id, amount)
}

func (r *Repository) DeleteTransaction(ctx context.Context, id int) error {
	return r.transactionStorage.DeleteTransaction(ctx, id)
}
