package transaction

import (
	"context"
	"finance/internal/model"
)

type Repository struct {
	transactionStorage storage
}

func New(transactionStorage storage) *Repository {
	return &Repository{transactionStorage: transactionStorage}
}

func (r *Repository) AddTransaction(ctx context.Context, categoryID int, amount int64) error {
	return r.transactionStorage.AddTransaction(ctx, categoryID, amount)
}

func (r *Repository) GetAllTransaction(ctx context.Context, categoryID int) ([]model.Transaction, error) {
	return r.transactionStorage.GetAllTransaction(ctx, categoryID)
}

func (r *Repository) GetTransaction(ctx context.Context, id int) (model.Transaction, error) {
	return r.transactionStorage.GetTransaction(ctx, id)
}

func (r *Repository) UpdateTransaction(ctx context.Context, id int, amount int64) error {
	return r.transactionStorage.UpdateTransaction(ctx, id, amount)
}

func (r *Repository) DeleteTransaction(ctx context.Context, id int) error {
	return r.transactionStorage.DeleteTransaction(ctx, id)
}
